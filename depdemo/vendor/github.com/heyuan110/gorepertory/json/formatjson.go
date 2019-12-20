package formatjson

import (
	"encoding/json"
)

//Decode json string to interface
func Decode(json_str string)  interface{}{
	if json_str == "" {
		return  nil
	}
	var s interface{}
	err := json.Unmarshal([]byte(json_str),&s)
	if err != nil {
		return nil
	}
	return s
}

//Encode interface to json string
func Encode(i interface{}) string  {
	if i == nil {
		return ""
	}
	s,err := json.Marshal(i)
	if err != nil{
		return ""
	}
	return string(s)
}