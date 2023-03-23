// Copyright 2021 Ross Light
// SPDX-License-Identifier: ISC

package main

import (
	"github.com/kwilteam/go-sqlite/sqlitefile"
)

func main() {
	var f *sqlitefile.File
	f, _ = sqlitefile.NewFile(nil)
	var buf *sqlitefile.Buffer
	buf, _ = sqlitefile.NewBuffer(nil)

	_, _ = f, buf
}
