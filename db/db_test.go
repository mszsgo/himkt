package db

import (
	"encoding/json"
	"github.com/mszsgo/himkt/env"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestConnect(t *testing.T) {
	c := Connect(env.HM_MONGO_CONNECTION_STRING)
	t.Log(c)
}

func TestJsonQuery(t *testing.T) {
	query := `[{"$match": {"actId":"1234567890","prodId":"10"}},{"$project": {"_id":0,"mobile":1}}]`

	var a bson.A
	err := bson.UnmarshalExtJSON([]byte(query), true, &a)
	if err != nil {
		t.Error(err)
		return
	}

	c := Connect(env.HM_MONGO_CONNECTION_STRING)
	cursor, err := c.Collection("hm_rights_order").Aggregate(nil, a)
	if err != nil {
		t.Error(err)
		return
	}
	var rmap []map[string]interface{}
	err = cursor.All(nil, &rmap)
	if err != nil {
		t.Error(err)
		return
	}
	bytes, err := json.Marshal(rmap)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
}
