package genid

import (
	"log"
	"math/rand"
	"time"

	"github.com/mszsgo/snowflake"
)

var (
	_genNode *snowflake.Node
)

func IdNode() *snowflake.Node {
	if _genNode != nil {
		return _genNode
	}
	rand.Seed(time.Now().UnixNano())
	_genNode, err := snowflake.NewNode(rand.Int63n(31))
	if err != nil {
		log.Panic(err)
	}
	return _genNode
}

func GenId() string {
	return IdNode().Generate().String()
}
