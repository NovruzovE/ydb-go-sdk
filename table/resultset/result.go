package resultset

import (
	"context"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/table/scanner"
)

// Result is a result of a query.
//
// Use NextResultSet(), NextRow() and Scan() to advance through the result sets,
// its rows and row's items.
//
//     res, err := s.Execute(ctx, txc, "SELECT ...")
//     defer res.Close()
//     for res.NextResultSet(ctx) {
//         for res.NextRow() {
//             var id int64
//             var name *string //optional value
//             res.Scan(&id,&name)
//         }
//     }
//     if err := res.err() { // get any error encountered during iteration
//         // handle error
//     }
//
// If current value under scan
// is not requested types, then res.err() become non-nil.
// After that, NextResultSet(), NextRow() will return false.
type Result interface {

	// NextResultSet selects next result set in the result.
	// columns - names of columns in the resultSet that will be scanned
	// It returns false if there are no more result sets.
	// Stream sets are supported.
	NextResultSet(ctx context.Context, columns ...string) bool

	// CurrentResultSet get current result set to use ColumnCount(), RowCount() and other methods
	CurrentResultSet() scanner.ResultSet

	// HasNextRow reports whether result row may be advanced.
	// It may be useful to call HasNextRow() instead of NextRow() to look ahead
	// without advancing the result rows.
	HasNextRow() bool

	// NextRow selects next row in the current result set.
	// It returns false if there are no more rows in the result set.
	NextRow() bool

	// ScanWithDefaults scan with default types values.
	// Nil values applied as default value types
	// Input params - pointers to types.
	ScanWithDefaults(values ...interface{}) error

	// Scan values.
	// Input params - pointers to types:
	//   bool
	//   int8
	//   uint8
	//   int16
	//   uint16
	//   int32
	//   uint32
	//   int64
	//   uint64
	//   float32
	//   float64
	//   []byte
	//   [16]byte
	//   string
	//   time.Time
	//   time.Duration
	//   ydb.Value
	// For custom types implement sql.Scanner interface.
	// For optional types use double pointer construction.
	// For unknown types use interface types.
	// Supported scanning byte arrays of various length.
	// For complex yql types: Dict, List, Tuple and own specific scanning logic implement ydb.Scanner with UnmarshalYDB method
	// See examples for more detailed information.
	// Output param - Scanner error
	Scan(values ...interface{}) error

	// Stats returns query execution QueryStats.
	Stats() (s scanner.QueryStats)

	// Err return scanner error
	// To handle errors, do not need to check after scanning each row
	// It is enough to check after reading all ResultSet
	Err() error

	// Close closes the Result, preventing further iteration.
	Close() error

	///<--------------non-stream-----------------

	// HasNextResultSet reports whether result set may be advanced.
	// It may be useful to call HasNextResultSet() instead of NextResultSet() to look ahead
	// without advancing the result set.
	// Note that it does not work with sets from stream.
	HasNextResultSet() bool

	// ResultSetCount returns number of result sets.
	// Note that it does not work if r is the result of streaming operation.
	ResultSetCount() int

	// TotalRowCount returns the number of rows among the all result sets.
	// Note that it does not work if r is the result of streaming operation.
	TotalRowCount() int

	///--------------non-stream-----------------/>
}

var _ Result = &scanner.Result{}