package services

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

type databaseResult struct {
	columns []string
	rows    [][]driver.Value
	err     error
}

type databaseStub struct {
	mu      sync.Mutex
	results []databaseResult
}

func (stub *databaseStub) next() databaseResult {
	stub.mu.Lock()
	defer stub.mu.Unlock()
	if len(stub.results) == 0 {
		return databaseResult{err: errors.New("consulta inesperada")}
	}
	result := stub.results[0]
	stub.results = stub.results[1:]
	return result
}

type stubDriver struct{ stub *databaseStub }

func (d stubDriver) Open(string) (driver.Conn, error) { return &stubConn{stub: d.stub}, nil }

type stubConn struct{ stub *databaseStub }

func (c *stubConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("prepare no soportado")
}
func (c *stubConn) Close() error               { return nil }
func (c *stubConn) Begin() (driver.Tx, error)  { return nil, errors.New("tx no soportada") }
func (c *stubConn) Ping(context.Context) error { return nil }

func (c *stubConn) ExecContext(context.Context, string, []driver.NamedValue) (driver.Result, error) {
	result := c.stub.next()
	if result.err != nil {
		return nil, result.err
	}
	return driver.RowsAffected(1), nil
}

func (c *stubConn) QueryContext(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
	result := c.stub.next()
	if result.err != nil {
		return nil, result.err
	}
	return &stubRows{columns: result.columns, rows: result.rows}, nil
}

type stubRows struct {
	columns []string
	rows    [][]driver.Value
	index   int
}

func (r *stubRows) Columns() []string { return r.columns }
func (r *stubRows) Close() error      { return nil }
func (r *stubRows) Next(dest []driver.Value) error {
	if r.index >= len(r.rows) {
		return io.EOF
	}
	copy(dest, r.rows[r.index])
	r.index++
	return nil
}

func useDatabaseStub(t *testing.T, results ...databaseResult) {
	t.Helper()
	stub := &databaseStub{results: results}
	name := "service_stub_" + t.Name()
	sql.Register(name, stubDriver{stub: stub})
	connection, err := sql.Open(name, "")
	if err != nil {
		t.Fatalf("no se pudo abrir la base de datos de prueba: %v", err)
	}
	previous := db.DB
	db.DB = connection
	t.Cleanup(func() {
		_ = connection.Close()
		db.DB = previous
	})
}
