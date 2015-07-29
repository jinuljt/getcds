package getcds

import "testing"

//etcd directory
/*
{"action":"get","node":{"key":"/test","dir":true,"nodes":[{"key":"/test/i32","value":"32","modifiedIndex":72,"createdIndex":72},{"key":"/test/i64","value":"64","modifiedIndex":73,"createdIndex":73},{"key":"/test/str","value":"string","modifiedIndex":74,"createdIndex":74},{"key":"/test/s","dir":true,"nodes":[{"key":"/test/s/str","value":"string","modifiedIndex":77,"createdIndex":77},{"key":"/test/s/i32","value":"32","modifiedIndex":75,"createdIndex":75},{"key":"/test/s/i64","value":"64","modifiedIndex":76,"createdIndex":76}],"modifiedIndex":75,"createdIndex":75}],"modifiedIndex":72,"createdIndex":72}}
*/

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

var SFailure struct {
	I32       int `etcd:"i32"`
	NotExists int `etcd:"not_exists"`
}

func TestGoEtcdSchema(t *testing.T) {
	machines := []string{"http://192.168.1.58:2379"}
	client := NewClient(machines)
	defer client.Close()

	if err := client.Unmarshal("/test", &SSuccess); err != nil {
		t.Errorf("getcd unmarshal error due to", err)
	}

	if err := client.Unmarshal("/test", &SFailure); err == nil {
		t.Errorf("getcd unmarshal should error")
	}
}
