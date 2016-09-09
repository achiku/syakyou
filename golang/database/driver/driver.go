package driver

// Value is a value that drivers must be able to handle.
// It is either nil or an instance of one of these types:
//   int64
//   float64
//   bool
//   []byte
//   string   [*] everywhere except from Rows.Next.
//   time.Time
type Value interface{}

// Driver is the interface that must be implemented by a database
// driver.
type Driver interface {
	Open(name string) (Conn, error)
}

// Conn is a connection to a database. It is not used concurrently
// by multiple goroutines.
type Conn interface {
	Prepare(query string) (Stmt, error)
	Close() error
	Begin() (Tx, error)
}

// Stmt is prepared statement. It is bound to a Conn and not
// used by multiple goroutines concurrently.
type Stmt interface {
	Close() error
	NumInput() int
	Exec(args []Value) (Result, error)
	Query(args []Value) (Rows, error)
}

// Result is the result of a query execution.
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Rows is an interator over an executed query's results.
type Rows interface {
	Columns() []string
	Close() error
	Next(dest []Value) error
}

// Tx database transaction
type Tx interface {
	Commit() error
	Rollback() error
}

// RowsAffected implements Result for an INSERT or UPDATE operation
// which mutates a number of rows.
type RowsAffected int64
