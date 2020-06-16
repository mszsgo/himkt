package genid

import (
	"log"
	"testing"
)

func TestGenId(t *testing.T) {
	//g, _ := snowflake.NewNode(int64(1))
	for i := 0; i < 1000; i++ {
		log.Println(GenId())
		//log.Println(rand.Int63n(31))
	}
}
