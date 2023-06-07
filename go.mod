module github.com/kwilteam/go-sqlite

go 1.19

require (
	crawshaw.io/iox v0.0.0-20181124134642-c51c3df30797
	github.com/chzyer/readline v1.5.0
	github.com/google/go-cmp v0.5.9
	modernc.org/libc v1.22.5
	modernc.org/sqlite v1.20.5-0.20230220170856-13895386cf24
)

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/sys v0.5.0 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
)

replace modernc.org/sqlite v1.20.5-0.20230220170856-13895386cf24 => github.com/kwilteam/sqlite-engine v0.0.0-20230607212633-2984ed83bdba

retract (
	v0.9.1 // Contains retractions only.
	v0.9.0 // Had libc memgrind issues.
)
