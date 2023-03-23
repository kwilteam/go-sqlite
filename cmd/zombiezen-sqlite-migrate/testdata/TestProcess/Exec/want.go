// Copyright 2021 Ross Light
// SPDX-License-Identifier: ISC

package main

import (
	"github.com/kwilteam/go-sqlite"
	sqlitefile2 "github.com/kwilteam/go-sqlite/sqlitefile"
	"github.com/kwilteam/go-sqlite/sqlitex"
)

func main() {
	var conn *sqlite.Conn
	var file *sqlitefile2.File
	sqlitex.ExecuteScriptFS(conn, nil, "foo.sql", &sqlitex.ExecOptions{
		Args: []interface{}{1, "foo"},
	})
	sqlitex.Exec(conn, `SELECT 1;`, nil)
	_ = file
}
