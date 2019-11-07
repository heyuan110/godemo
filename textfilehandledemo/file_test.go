package textfilehandledemo

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestHandleFile(t *testing.T)  {
	os.Mkdir("testdir",0777)
	os.MkdirAll("testdir/t1/t2/t3",0777)
	err := os.Remove("testdir")
	if err != nil{
		fmt.Println(err)
	}
	//os.RemoveAll("testdir")

	userFile := "test_str.txt"
	fout,err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i:=0;i<5;i++{
		fmt.Println(strconv.Itoa(i))
		 fout.WriteString("string -> just a test! " + fmt.Sprintf("%d",i)+"\n")
		 fout.Write([]byte("bype -> Just a test! "+fmt.Sprintf("%d",i)+"\n"))
	}
	fmt.Println("start read ........")
	fi,err := os.Open(userFile)
	defer fi.Close()
	buf := make([]byte,3<<10)
	n,_ := fi.Read(buf)
	fmt.Println(string(buf))
	fmt.Println("n->",n)
	fmt.Println(len(buf),"-",cap(buf),"-",binary.Size(buf))

	os.Stdout.Write(buf)
}