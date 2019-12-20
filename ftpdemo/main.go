package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
)

func main() {
	//done := make(chan os.Signal, 1)
	//signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//fmt.Println("hello")
	//js,err := simplejson.NewJson([]byte(`{"a":"12","b":12}`))
	//if err != nil {
	//	panic("json format error")
	//}
	//fmt.Println(js)
	//s1,_ := js.Get("a").String()
	//fmt.Println("s1->",s1)
	//s2,_ := js.Get("b").Int()
	//fmt.Println("s2->",s2)
	//
	//m:= map[string]interface{}{
	//	"key":"hello world!",
	//}
	//json := formatjson.Encode(m)
	//fmt.Println(json)

	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}

	//<-done
}
