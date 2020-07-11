package hm

import "testing"

func TestUrlAddParams(t *testing.T) {
	t.Log(UrlAddParams("http://sss.com/sss", "code", "vv"))
	t.Log(UrlAddParams("http://sss.com/sss?k1=v1", "code", "vv"))
	t.Log(UrlAddParams("http://sss.com/sss#/ss?k1=v1", "code", "vv"))
	t.Log(UrlAddParams("http://sss.com/sss?k1=v1#/ss?k1=v1", "code", "vv"))
}
