package tg

import (
	"bytes"
	"fmt"
	"net/url"
	"reflect"
	"sort"
)

// query represents query parameters because I find url.Values not handy enough.
type Query map[string]interface{}

func (q Query) Encode() string {
	var buf bytes.Buffer
	first := true
	write := func(k string, v interface{}) {
		if first {
			first = false
		} else {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(fmt.Sprint(v)))
	}
	keys := make([]string, len(q))
	i := 0
	for k, _ := range q {
		keys[i] = k
		i += 1
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := q[k]
		vv := reflect.ValueOf(v)
		switch vv.Kind() {
		case reflect.Array, reflect.Slice:
			for i := 0; i < vv.Len(); i++ {
				write(k, vv.Index(i).Interface())
			}
		default:
			write(k, v)
		}
	}
	return buf.String()
}
