package http

import (
	"encoding/json"
	"testing"
)

func TestGet(t *testing.T) {
	v := make(map[string]interface{})
	_, err := GetJson("http://api.kanzhihu.com/getposts", &v)
	if err != nil {
		t.Error(err)
	}
	d, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(d))
}

func TestPost(t *testing.T) {
	req := make(map[string]interface{})
	resp := make(map[string]interface{})
	req["title"] = "foo"
	req["body"] = "bar"
	req["userId"] = 1
	_, err := PostJson("http://jsonplaceholder.typicode.com/posts", req, &resp)
	if err != nil {
		t.Error(err)
	}
	d, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(d))
}
