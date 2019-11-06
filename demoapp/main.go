package main

import (
	"fmt"
	"math"
	"runtime"
)

type testInt func(int) bool

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human
	speciality string
}

func addDataWithPoint(a *int) int {
	*a = *a + 1
	return  *a
}

func addData(a int) int {
	a = a + 1
	return  a
}

func isOdd(integer int) bool  {
	if integer % 2 == 0 {
		return false
	}else{
		return true
	}
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int  {
	var result []int
	for _,val := range slice{
		if f(val){
			result = append(result,val)
		}
	}
	return  result
}

type Rectangle struct {
	width,height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area()float64
}

type ShapeR interface {
	area()float64
	length()float64
}

func (r Rectangle) area() float64 {
	return r.width*r.height
}

func (r Rectangle)length() float64  {
	return (r.width+r.height)*2.0
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {

	x1 := 2
	y1 := addData(x1)
	fmt.Println("x1=",x1)
	fmt.Println("y1=",y1)

	x := 2
	y := addDataWithPoint(&x)
	fmt.Println("x=",x)
	fmt.Println("y=",y)

	for i:=0;i<5;i++{
		defer fmt.Printf("%d",i)
	}

	slice := []int{1,2,3,4,5,6,7}
	fmt.Println("slice=",slice)

	odd := filter(slice,isOdd)
	fmt.Println("odd elements of slice are ",odd)

	even := filter(slice,isEven)
	fmt.Println("event elements of slice are ",even)

	mark := Student{Human{"Mark",25,125},"Computer science"}
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	var s Shape
	var s1 ShapeR

	r := Rectangle{30,4}
	s = r
	fmt.Println("shape of Rectangle is: ",s.area())
	s1 = r
	fmt.Println("shape of Rectangle lenth is: ",s1.length())


	c :=Circle{4.2}
	s = c
	fmt.Println("shape of Circle is: ",s.area())

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumCgoCall())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))

}
