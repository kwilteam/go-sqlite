// Copyright 2023 Ross Light
// SPDX-License-Identifier: ISC

package generateseries_test

import (
	"fmt"
	"log"

	"github.com/kwilteam/go-sqlite"
	"github.com/kwilteam/go-sqlite/ext/generateseries"
	"github.com/kwilteam/go-sqlite/sqlitex"
)

func Example() {
	conn, err := sqlite.OpenConn(":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := generateseries.Register(conn); err != nil {
		log.Fatal(err)
	}
	err = sqlitex.ExecuteTransient(
		conn,
		`SELECT * FROM generate_series(0, 20, 5);`,
		&sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				fmt.Printf("%2d\n", stmt.ColumnInt(0))
				return nil
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	//  0
	//  5
	// 10
	// 15
	// 20
}
