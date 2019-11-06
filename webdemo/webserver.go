package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	Pid int
	Product_name string
	Image_url string
	Product_Code string
	Supplier_Id int
}

type header struct {
	Encryption  string `json:"encryption"`
	Timestamp   int64  `json:"timestamp"`
	Key         string `json:"key"`
	Partnercode int    `json:"partnercode"`
}

type ProductList []Product

func sayHelloName(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key",k)
		fmt.Println(v)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello bruce!")
}

func hotProducts(w http.ResponseWriter, r *http.Request)  {
	plist := ProductList{
		Product{Pid:123,Product_name:"XieZi",Image_url:"http://image.demo.com"},
		Product{Pid:124,Product_name:"WaZi",Image_url:"http://image.demo.com"},
		Product{Pid:125,Product_name:"YiFu",Image_url:"http://image.demo.com"},
	}
	jsonBytes, err := json.Marshal(plist)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w,string(jsonBytes))
	fmt.Println(string(jsonBytes))
}

func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:",r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w,nil))
	}else{
		r.ParseForm()
		r.Form.Set("uid","1234")
		fmt.Println(r.Form)
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
		fmt.Println("uid:",r.Form.Get("uid"))
	}
}

func upload(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("method:",r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))
		t,_ := template.ParseFiles("upload.gtpl")
		t.Execute(w,token)
	}else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",handler.Header)
		f,err := os.OpenFile("./test/"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f,file)
	}
}

func main() {
	var bodyBuf *bytes.Buffer = &bytes.Buffer{}
	fmt.Println(reflect.TypeOf(bodyBuf))

	var teststring string = "teststring"
	fmt.Println(teststring)
	fmt.Println(reflect.TypeOf(teststring))

	var bytestring []byte = []byte(teststring)
	fmt.Println(bytestring)
	fmt.Println(reflect.TypeOf(bytestring))

	var md5bytes [16]byte = md5.Sum(bytestring)
	fmt.Println(md5bytes)
	fmt.Println(reflect.TypeOf(md5bytes))

	var md5string string = fmt.Sprintf("%x",md5bytes)
	fmt.Println(md5string)
	fmt.Println(64 << 20)
	fmt.Println(32 << 20)
	fmt.Println("Start webserver...")
	http.HandleFunc("/",sayHelloName)
	http.HandleFunc("/hot/products",hotProducts)
	http.HandleFunc("/login",login)
	http.HandleFunc("/upload",upload)
	port := "9090"
	fmt.Println("Webserver running")
	fmt.Println("Please open http://localhost:"+port)

	ptest:=Product{Pid:123,Product_name:"XieZi",Image_url:"http://image.demo.com"}
	jsonBytes, erro := json.Marshal(ptest)
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(string(jsonBytes))


	headerTest := header{
		Encryption:  "sha",
		Timestamp:   1482463793,
		Key:         "2342874840784a81d4d9e335aaf76260",
		Partnercode: 10025,
	}
	jsons, errs := json.Marshal(headerTest) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println(string(jsons)) //byte[]转换成string 输出

	err := http.ListenAndServe(":"+port,nil)
	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}



}