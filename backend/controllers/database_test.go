package controllers

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"sync"
	"testing"

	"proyecto-desarrollo-sw/backend/db"
)

type controllerDBResult struct {
	columns []string
	rows    [][]driver.Value
	err     error
}

type controllerDBStub struct {
	mu      sync.Mutex
	results []controllerDBResult
}

func (stub *controllerDBStub) next() controllerDBResult {
	stub.mu.Lock()
	defer stub.mu.Unlock()
	if len(stub.results) == 0 {
		return controllerDBResult{err: errors.New("consulta inesperada")}
	}
	result := stub.results[0]
	stub.results = stub.results[1:]
	return result
}

type controllerDriver struct{ stub *controllerDBStub }

func (d controllerDriver) Open(string) (driver.Conn, error) {
	return &controllerConn{stub: d.stub}, nil
}

type controllerConn struct{ stub *controllerDBStub }

func (c *controllerConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("prepare no soportado")
}
func (c *controllerConn) Close() error               { return nil }
func (c *controllerConn) Begin() (driver.Tx, error)  { return nil, errors.New("tx no soportada") }
func (c *controllerConn) Ping(context.Context) error { return nil }
func (c *controllerConn) ExecContext(context.Context, string, []driver.NamedValue) (driver.Result, error) {
	result := c.stub.next()
	if result.err != nil {
		return nil, result.err
	}
	return driver.RowsAffected(1), nil
}
func (c *controllerConn) QueryContext(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
	result := c.stub.next()
	if result.err != nil {
		return nil, result.err
	}
	return &controllerRows{columns: result.columns, rows: result.rows}, nil
}

type controllerRows struct {
	columns []string
	rows    [][]driver.Value
	index   int
}

func (r *controllerRows) Columns() []string { return r.columns }
func (r *controllerRows) Close() error      { return nil }
func (r *controllerRows) Next(dest []driver.Value) error {
	if r.index >= len(r.rows) {
		return io.EOF
	}
	copy(dest, r.rows[r.index])
	r.index++
	return nil
}

func useControllerDB(t *testing.T, results ...controllerDBResult) {
	t.Helper()
	stub := &controllerDBStub{results: results}
	name := "controller_stub_" + t.Name()
	sql.Register(name, controllerDriver{stub: stub})
	connection, err := sql.Open(name, "")
	if err != nil {
		t.Fatal(err)
	}
	previous := db.DB
	db.DB = connection
	t.Cleanup(func() {
		_ = connection.Close()
		db.DB = previous
	})
}
