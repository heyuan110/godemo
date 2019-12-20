package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	formatjson "github.com/heyuan110/gorepertory/json"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("hello")
	js,err := simplejson.NewJson([]byte(`{"a":"12","b":12}`))
	if err != nil {
		panic("json format error")
	}
	fmt.Println(js)
	s1,_ := js.Get("a").String()
	fmt.Println("s1->",s1)
	s2,_ := js.Get("b").Int()
	fmt.Println("s2->",s2)

	m:= map[string]interface{}{
		"key":"hello world!",
	}
	json := formatjson.Encode(m)
	fmt.Println(json)

	<-done
}
