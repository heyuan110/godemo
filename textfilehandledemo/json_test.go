package textfilehandledemo

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"testing"
)

type Server struct {
	ServerName string `json:"server_name"`
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

func TestSimpleJson(t *testing.T)  {
	js,err := simplejson.NewJson([]byte(`{
    "test": {
        "array": [1, "2", 3],
        "int": 10,
        "float": 5.150,
        "bignum": 9223372036854775807,
        "string": "simplejson",
        "bool": true
    }
}`))

	if err != nil {
		panic(err)
	}
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
	fmt.Println(arr,i,ms)
}

func TestExportJson(t *testing.T)  {
	var s Serverslice
	s.Servers = append(s.Servers,Server{ServerName:"ubuntu1",ServerIP:"192.168.11.119"})
	s.Servers = append(s.Servers,Server{ServerName:"ubuntu2",ServerIP:"192.168.11.112"})
	b,err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err",err)
	}
	fmt.Println(string(b))

	var s_arr []Server
	s_arr = append(s_arr,Server{ServerName:"ubuntu1",ServerIP:"192.168.11.119"})
	s_arr = append(s_arr,Server{ServerName:"ubuntu2",ServerIP:"192.168.11.129"})
	b1,_ := json.Marshal(s_arr)
	fmt.Println(string(b1))

	}

