package getcds

import (
	"errors"
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strconv"

	goetcd "github.com/coreos/go-etcd/etcd"
)

var defaultTag = "etcd"

type Client struct {
	client *goetcd.Client
	tag    string
}

func NewClient(machines []string) *Client {
	c := new(Client)
	c.tag = defaultTag
	c.client = goetcd.NewClient(machines)
	return c
}

func (c *Client) setStruct(directory string, value *reflect.Value) (err error) {
	tv := value.Type()
	for i := 0; i < value.NumField(); i++ {
		if tagv := tv.Field(i).Tag.Get(c.tag); tagv != "" {
			field := value.Field(i)
			if field.Kind() == reflect.Struct {
				if err := c.setStruct(path.Join(directory, tagv), &field); err != nil {
					return err
				}
			} else {
				key := path.Join(directory, tagv)
				response, err := c.client.Get(key, false, false)
				if err != nil {
					return err
				}

				value := response.Node.Value
				switch field.Kind() {
				case reflect.String:
					field.SetString(value)
				case reflect.Int64:
					i, _ := strconv.Atoi(value)
					field.SetInt(int64(i))
				case reflect.Int:
					i, _ := strconv.Atoi(value)
					field.SetInt(int64(i))
				}
			}
		}
	}
	return nil
}

func (c *Client) Unmarshal(directory string, v interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			err = r.(error)
		}
	}()

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New(fmt.Sprintf("nil or not ptr"))
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("only accept struct; got %T", v))
	}

	return c.setStruct(directory, &rv)
}

func (c *Client) SetTag(tag string) {
	c.tag = tag
}

func (c *Client) Close() {
	c.client.Close()
}
