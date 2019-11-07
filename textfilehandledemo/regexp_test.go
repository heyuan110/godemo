package textfilehandledemo

import (
	"fmt"
	"regexp"
	"testing"
)

func IsIP(ip string)(b bool)  {
	if m,_ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$",ip); !m {
		return false
	}
	return true
}

func TestIP(t *testing.T)  {
	fmt.Println(IsIP("12.23.2w3.23"))
}