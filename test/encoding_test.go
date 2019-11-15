package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
	"log"
	"regexp"
	"testing"
)

func TestDemo(t *testing.T) {
	username := "asfaeWs-d23423f"
	if ok,_ := regexp.MatchString("^[0-9a-zA-Z]+$",username); ok {
		fmt.Println(username," is matched")
	}else{
		fmt.Println(username, " is not matched")
	}

	msg := "His money is twice tainted: 'taint yours and 'taint mine."

	h := sha256.New()
	io.WriteString(h, msg)
	fmt.Printf("%x",h.Sum(nil))

	fmt.Println()

	h1 := sha256.New()
	io.WriteString(h1, msg)
	fmt.Printf("%x", h1.Sum(nil))

	fmt.Println()
	h2 := md5.New()
	io.WriteString(h2, msg)
	fmt.Printf("%x", h2.Sum(nil))

	salt := []byte("saltstring")

	dk, err := scrypt.Key([]byte(msg), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x", dk)
	fmt.Println()
	base64_str := base64.StdEncoding.EncodeToString(dk)
	fmt.Println(base64_str)
	decode_base64_bytes,_ := base64.StdEncoding.DecodeString(base64_str)
	fmt.Printf("%x",decode_base64_bytes)
	fmt.Println()

	//加密字符串
	//commonTv长度16位置
	var commonIV = []byte("zmknm.a23jkljl;k")
	plaintext := []byte(msg)
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"
	println(len(key_text))
	c, err := aes.NewCipher([]byte(key_text))
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s=>%x\n", plaintext, ciphertext)

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy)

}

