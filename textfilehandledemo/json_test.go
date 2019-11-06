package textfilehandledemo

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Server struct {
	ServerName string
	ServerIP string
}

type Serverslice struct {
	Servers []Server
}


func TestParseJsonToStruct(t *testing.T)  {
	var s Serverslice
	json_string := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(json_string),&s)
	fmt.Println(s)
	for _,v :=range s.Servers {
		fmt.Println(v)
		fmt.Println(v.ServerIP,"-",v.ServerName)
	}
}

func TestParseJsonToInterface(t *testing.T)  {
	var i interface{}
	json_bytes := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	json.Unmarshal(json_bytes,&i)
	fmt.Println(i)
	//通过断言方式取值
	m := i.(map[string]interface{})
	fmt.Println(m)
	for k,v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k,"is float64",vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}