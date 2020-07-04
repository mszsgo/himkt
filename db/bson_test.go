package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestJsonToBson(t *testing.T) {
	b1 := []byte(`[{ "a":"a1" }]`)
	var d bson.A
	err := bson.UnmarshalExtJSON(b1, true, &d)
	if err != nil {
		t.Error(err)
	}
	t.Log(d)

}
