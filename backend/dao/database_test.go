package dao

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"sync"
	"testing"

	"proyecto-desarrollo-sw/backend/db"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type daoDBResult struct {
	columns []string
	rows    [][]driver.Value
	err     error
}

type daoDBStub struct {
	mu      sync.Mutex
	results []daoDBResult
}

func (stub *daoDBStub) next() daoDBResult {
	stub.mu.Lock()
	defer stub.mu.Unlock()
	if len(stub.results) == 0 {
		return daoDBResult{err: errors.New("consulta inesperada")}
	}
	result := stub.results[0]
	stub.results = stub.results[1:]
	return result
}

type daoDriver struct{ stub *daoDBStub }

func (d daoDriver) Open(string) (driver.Conn, error) { return &daoConn{stub: d.stub}, nil }

type daoConn struct{ stub *daoDBStub }

type daoResult struct{ rows int64 }

func (r daoResult) LastInsertId() (int64, error) { return 1, nil }
func (r daoResult) RowsAffected() (int64, error) { return r.rows, nil }

func (c *daoConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("prepare no soportado")
}
func (c *daoConn) Close() error               { return nil }
func (c *daoConn) Begin() (driver.Tx, error)  { return nil, errors.New("tx no soportada") }
func (c *daoConn) Ping(context.Context) error { return nil }

func (c *daoConn) ExecContext(context.Context, string, []driver.NamedValue) (driver.Result, error) {
	result := c.stub.next()
	if result.err != nil {
		return nil, result.err
	}
	return daoResult{rows: 1}, nil
}

func (c *daoConn) QueryContext(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
	result := c.stub.next()
	if result.err != nil {
		return nil, result.err
	}
	return &daoRows{columns: result.columns, rows: result.rows}, nil
}

type daoRows struct {
	columns []string
	rows    [][]driver.Value
	index   int
}

func (r *daoRows) Columns() []string { return r.columns }
func (r *daoRows) Close() error      { return nil }
func (r *daoRows) Next(dest []driver.Value) error {
	if r.index >= len(r.rows) {
		return io.EOF
	}
	copy(dest, r.rows[r.index])
	r.index++
	return nil
}

func useDaoDB(t *testing.T, results ...daoDBResult) {
	t.Helper()
	stub := &daoDBStub{results: results}
	name := "dao_stub_" + t.Name()
	sql.Register(name, daoDriver{stub: stub})
	connection, err := sql.Open(name, "")
	if err != nil {
		t.Fatalf("no se pudo abrir la base de datos de prueba: %v", err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: connection, SkipInitializeWithVersion: true,
	}), &gorm.Config{SkipDefaultTransaction: true, DisableAutomaticPing: true})
	if err != nil {
		t.Fatalf("no se pudo configurar GORM para la prueba: %v", err)
	}
	previous := db.DB
	db.DB = gormDB
	t.Cleanup(func() {
		_ = connection.Close()
		db.DB = previous
	})
}
