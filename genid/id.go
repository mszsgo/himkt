package genid

import (
	"github.com/mszsgo/snowflake"
	"math/rand"
)

var g, _ = snowflake.NewNode(rand.Int63n(31))

func GenId() string {
	return g.Generate().String()
}
