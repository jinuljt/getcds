package getcds

import (
	"fmt"
	"testing"
)

var SSuccess struct {
	I32 int    `etcd:"i32"`
	I64 int    `etcd:"i64"`
	Str string `etcd:"str"`

	S struct {
		I32 int    `etcd:"i32"`
		I64 int    `etcd:"i64"`
		Str string `etcd:"str"`
	} `etcd:"s"`
}

func TestGoEtcdSchema(t *testing.T) {
	machines := []string{"http://192.168.1.58:2357"}
	etcdClient := NewClient(machines)
	if err := etcdClient.Unmarshal("/test", &SSuccess); err != nil {
		t.Errorf("getcd unmarshal error due to", err)
	}
	fmt.Println(SSuccess)
}
