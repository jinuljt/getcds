## getcds ##

getcds 是一个把etcd中的directory直接映射到struct的工具


## getcds 的限制 ##
1、getcds 只实现了 unmarshal，未实现marshal(struct->etcd directory)
2、struct中只允许使用int、int64、string、struct




## 如何使用 getcds ##

安装getcds
```
go get github.com/coreos/go-etcd/etcd
go get github.com/jinuljt/getcds
```


使用getcds
```
// 定义struct
var S struct {
	I32 int    `etcd:"i32"`
	I64 int    `etcd:"i64"`
	Str string `etcd:"str"`

	S struct {
		I32 int    `etcd:"i32"`
		I64 int    `etcd:"i64"`
		Str string `etcd:"str"`
	} `etcd:"s"`
}

machines := []string{"http://192.168.1.58:2379"}
client := NewClient(machines)
defer client.Close()

if err := client.Unmarshal("/test", &S); err != nil {
	fmt.Println("getcd unmarshal error due to", err)
}
```
