package main

import (
	"errors"
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am nokia, i can call you!")
}

type iPhone struct {
}

func (iphone iPhone) call() {
	fmt.Println("I am iPhone, can i call you ?")
}

type Person interface {
	name() string
	age() int
}

type Woman struct {
}

func (woman Woman) name() string {
	return "Lily"
}

func (woman Woman) age() int {
	return 22
}

type Men struct {
}

func (men Men) name() string {
	return "Bruce"
}

func (men Men) age() int {
	return 24
}

func sqrt(f float32) (float32, error) {
	if f < 0 {
		return 0, errors.New("值错误")
	}
	return f, nil
}

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(iPhone)
	phone.call()

	var person Person

	person = new(Men)
	fmt.Println(person.name())
	fmt.Println(person.age())

	person = new(Woman)
	fmt.Println(person.name())
	fmt.Println(person.age())

	result, erro := sqrt(-2)
	fmt.Println(result)
	if erro != nil {
		fmt.Println(erro)
	}

	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

}
