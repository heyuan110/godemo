package main

import (
	"fmt"
	"strconv"
)

const (
	i = 5 << iota
	j = 5 << iota
	k
	l
)

func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(a, b int) int {
	return a + b
}

func multiplicationTable() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			var ret string
			if i*j < 10 && j != 1 {
				ret = " " + strconv.Itoa(i*j)
			} else {
				ret = strconv.Itoa(i * j)
			}
			fmt.Print(j, " * ", i, " = ", ret, "   ")
		}
		fmt.Print("\n")
	}
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

type User struct {
	name        string
	sex         string
	salary      string
	position_id int
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book aRuthor : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printUser(user *User) {
	fmt.Printf("user name: %s \n", user.name)
}

func printSliceStr(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func compressionString(str string, max_length int) string {
	string_len := len(str)
	if string_len <= max_length {
		return str
	}
	half_len := max_length/2
	return str[0:half_len] +"..."+ " [^...^] " + "..." + str[string_len-half_len:string_len]
}

func main() {
	teststr := "start... patpatArabicgn=265&utm_term=FB-@patpatArabic-Amy&nt=12-15-Baby-Toddlers-05adlink-54905   -  https://www.patpat.com?utm_source=FBFree&utm_medium=PPC-HZFB-CTT&utm_campaign=265&utm_term=FB-@patpatArabic-Amy&utm_content=12-15-Stylish-Fish-Scales-Pattern-Long-Sleeves-Romper-for-Baby-Girl-01 adlink-54904   -  https://www.patpat.com?utm_source=FB&utm_medium=PPC-Facebook-BR&utm_campaign=105&utm_term=Lanhan-PatPat-Daisy-2&utm_content=12-15-Baby-Toddlers-04adlink-54903   -  https://www.patpat.com?utm_source=FB&utm_medium=PPC-Facebook-BR&utm_campaign=105&utm_term=Lanhan-PatPat-Daisy-2&utm_content=12-15-Baby-Toddlers-03 adlink-54902   -  https://ar.patpat.com/product/Cute-Fox-Print-Sleeveless-Dress-for-Baby-Girl.html?flag=detail_recommend_3&utm_source=FBFree&utm_medium=PPC-HZFB-CTT&utm_campaign=265&utm_term=FB-@patpatArabic-Amy&utm_content=12-15-Cute-Fox-Print-Sleeveless-Dress-for-Baby-Girl-01 adlink-54901   -  https://ar.patpat.com/category/Baby-Toddlers/Baby-Toddler-Girl/One-Pieces/Rompers-Bodysuits.html?utm_source=FBFree&utm_medium=PPC-HZFB-CTT&utm_campaign=265&utm_term=FB-@patpatArabic-Amy&utm_content=12-15-Rompers-Bodysuits-01 adlink-54900   -  https://www.patpat.com?utm_source=FB&utm_medium=PPC-Facebook-BR&utm_campaign=105&utm_term=Lanhan-PatPat-Daisy-2&utm_content=12-15-Baby-Toddlers-02 adlink-54899   -  https://www.patpat.com?utm_source=FB&utm_medium=PPC-Facebook-MZL&utm_campaign=321&utm_term=PatPat-Conversion-MZL-1&utm_content=12-15-Simple-Splice-Large-Capacity-Shoulder-Bag-01 adlink-54898   -  https://www.patpat.com?utm_source=FB&utm_medium=PPC-Facebook-BR&utm_campaign=105&utm_term=Lanhan-PatPat-Daisy-2&utm_content=12-15-Baby-Toddlers-01 End!"
	fmt.Println(compressionString(teststr,2000))
	return

	var numbers = make([]int, 3, 5)
	printSliceStr(numbers)

	var b1 Books

	b1.title = "go title"
	b1.author = "go author"
	b1.book_id = 23223266666663
	fmt.Println(b1)
	printBook(&b1)

	fmt.Println(Books{"go", "bruce", "go subject", 13233})
	fmt.Println(Books{title: "go", author: "bruce", subject: "go subject", book_id: 13233})
	fmt.Println(Books{subject: "go subject", book_id: 13233})

	fmt.Println("hello")
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第1行c=%d\n", c)
	c = a - b
	fmt.Printf("第2行c=%d\n", c)

	if a == b {
		fmt.Printf("第一行 - a 等于 b\n")
	} else {
		fmt.Printf("第一行 - a 不等于 b\n")
	}
	if a < b {
		fmt.Printf("第二行 - a 小于 b\n")
	} else {
		fmt.Printf("第二行 - a 不小于 b\n")
	}

	if a > b {
		fmt.Printf("第三行 - a 大于 b\n")
	} else {
		fmt.Printf("第三行 - a 不大于 b\n")
	}

	/* Lets change value of a and b */
	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第四行 - a 小于等于 b\n")
	}
	if b >= a {
		fmt.Printf("第五行 - b 大于等于 a\n")
	}
	c = a
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c)

	c += a
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c)

	c -= a
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c)

	c *= a
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c)

	c /= a
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c)

	c = 200

	c <<= 2
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c)

	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c)

	/*  & 和 * 运算符实例 */
	var ptr *int = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	/* 调用函数并返回最大值 */
	var ret int = max(a, b)
	fmt.Printf("最大值是 : %d\n", ret)

	x, y := swap("Mahesh", "Kumar")
	fmt.Println(x, y)

	println("\n")

	multiplicationTable()

	// var balance [10]float32

	var ba = [5]float32{2.3, 2.34, 4.3, 23.3}
	ba[1] = 23423.3
	// var bc = [...]float32{4, 45, 6, 7, 7, 342, 45, 876}
	fmt.Println("ba[1]", ba[1])

	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var ar = [3][4]int{
		{0, 1, 29, 3},  /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	fmt.Println("ar[1][2]=", ar[1][2])

	var x1 int = 33
	fmt.Printf("变量的地址: %x\n", &x1)

	var h int = 20
	var ip *int

	ip = &h

	fmt.Printf("a address %x\n", &a)
	fmt.Printf("ip address %x\n", ip)
	fmt.Printf("*ip value %d\n", *ip)

	var ptrx *int
	fmt.Printf("ptr 的值为 : %x\n", ptrx)

	var numbers1 = make([]int, 3, 5)
	printSliceStr(numbers1)

	var numbers2 []int
	printSliceStr(numbers2)

	if numbers2 == nil {
		fmt.Printf("切片是空的!")
	}

	s2 := []int{1, 3, 34, 5, 56, 7, 8, 43}
	printSliceStr(s2)

	fmt.Println("s2[1:4]==", s2[3:4])

	s2 = append(s2, 3)

	fmt.Println(s2)

	s2 = append(s2, 1, 2, 12)

	fmt.Println(s2)

	nums := []int{1, 2, 43, 3, 5, 22}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index=", i)
		}
	}

	fmt.Println("sum:", sum)

	var countryCapMap map[string]string
	countryCapMap = make(map[string]string)

	countryCapMap["country"] = "China"
	countryCapMap["city"] = "ShenZhen"

	for k := range countryCapMap {
		fmt.Println(k, "键值的值是", countryCapMap[k])
	}

	ccm := map[string]string{"name": "bruce", "sex": "men", "age": "23"}
	fmt.Println(ccm)

	ccm["color"] = "red"
	fmt.Println(ccm)

	delete(ccm, "name")
	fmt.Println(ccm)

	var x11 int = 4
	fmt.Printf("%d的阶乘是%d\n", x11, Factorial(uint64(x11)))

	for x12 := 0; x12 < 30; x12++ {
		fmt.Printf("%d \t", fib(x12))
	}
}

/*
递归阶乘
*/
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
