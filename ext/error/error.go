package error

import (
	"fmt"

	"github.com/kwilteam/go-sqlite"
)

// impl is the implementation of the ERROR function.
var impl = &sqlite.FunctionImpl{
	NArgs:         1,
	Deterministic: true,
	AllowIndirect: true,
	Scalar:        errorFunc,
}

// Register registers the "ERROR" function on the given connection.
func Register(c *sqlite.Conn) error {
	return c.CreateFunction("error", impl)
}

func errorFunc(ctx sqlite.Context, args []sqlite.Value) (sqlite.Value, error) {
	return sqlite.Value{}, fmt.Errorf("error: %s", args[0].Text())
}
