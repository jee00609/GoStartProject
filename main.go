package main

import "fmt"

// 1. 변수
var a int
var f float32 = 11

var i,j,k int // 동일한 타입의 변수들일 경우, 나열하여 선언 가능하다.
var i2,j2,k2 int = 1,2,3

// 2. 상수
const c int = 10
const s string = "Hi"

const (
	Visa = "Visa"
	Master = "Master Card"
	Amex = "American Express"
)


// 한가지 유용한 팁으로 상수값을 0으로 순차적으로 부여하기 위해 iota 라는 identifier 를 사용가능하다
// iota 가 지정된 Apple 에는 0이 할당 되고 나머지 상수들은 순서대로 1씩 증가된 값을 부여받는다.
const (
	Apple = iota // 0
	Grape // 1
	Orange // 2
)

// 3. fmt 패키지

// 구조체 인스턴스 지원
type User struct {
	name string
	age  int
}


func main(){

	println("Hello World")
	a = 10
	i,j,k = 3,2,1

	u := User{"AWS",10}

	// 줄여서 변수 선언 가능
	shortS := "shortString"

	println(a)
	println(f)
	println(i,j,k)
	println(i2,j2,k2)

	println(c)
	println(s)

	println(Visa)
	println(Master)
	println(Amex)

	println(Apple,Grape,Orange)

	// 사용 불가!
	// illegal types for operand: print
	//println(u)
	fmt.Printf("%v\n", u) // {Gopher 10}
	fmt.Printf("%+v\n", u) // {name:Gopher, age:10}
	fmt.Printf("%#v\n", u) // main.user{name:"Gopher", age:10}

	fmt.Printf("Name: %s", shortS) // "Name: Gopher"
}
