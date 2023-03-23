// Copyright 2021 Ross Light
// SPDX-License-Identifier: ISC

package main

import "github.com/kwilteam/go-sqlite"

func main() {
	var db *sqlite.Conn
	db, _ = sqlite.OpenConn(":memory:", 0)
	db.AutocommitEnabled()
}
