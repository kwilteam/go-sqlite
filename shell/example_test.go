// Copyright 2021 Ross Light
// SPDX-License-Identifier: ISC

package shell_test

import (
	"fmt"
	"os"

	"github.com/kwilteam/go-sqlite"
	"github.com/kwilteam/go-sqlite/shell"
)

// This is a small program that emulates the behavior of the sqlite3 CLI.
// A path to a database can be passed on the command-line.
func Example() {
	dbName := ":memory:"
	if len(os.Args) > 1 {
		dbName = os.Args[1]
	}
	conn, err := sqlite.OpenConn(dbName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	shell.Run(conn)
	conn.Close()
}
