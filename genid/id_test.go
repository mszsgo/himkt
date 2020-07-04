package genid

import (
	"log"
	"testing"
	"time"
)

func TestGenId(t *testing.T) {
	//g, _ := snowflake.NewNode(int64(1))
	for i := 0; i < 100; i++ {
		log.Println(GenId())
		//log.Println(rand.Int63n(31))
	}
}

func TestA(t *testing.T) {

	et := time.Duration(1592304944216)
	t.Log(time.Now().Unix())

	tt := time.Now().Add(et * time.Millisecond)
	log.Println(tt.Format(time.RFC3339))
}

func TestE(t *testing.T) {
	nm := make(map[string]int)
	for i := 0; i < 1000; i++ {
		k := GenId()[16:]
		if nm[k] != 0 {
			t.Log("重复了", i)
			break
		}
		nm[k] = i
	}
}
