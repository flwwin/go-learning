package main

import "fmt"

// 基础数据类型：int int8 int16 int32 int 64
// unit8 ...unit64
// float32 float64  true/false

// 自增语法  i++ i--  自增的语法必须单独一行
func main() {
	//变量定义 var
	//常量 const

	//01- 先定义变量然后再赋值
	var name string
	name = "flw"

	var age int
	age = 20
	fmt.Println(name)
	fmt.Printf("名字，%s\n",name)
	fmt.Printf("年龄：%d\n",age)

	// 02 直接赋值
	var gendet = "man"

	fmt.Printf("gender: %s\n",gendet)

	// 03 直接定义赋值 ,使用自动推导（最常用的）
	address :="北京"
	fmt.Printf("address: %s\n",address)

	// 方法调用
	test(10,20)

	// 04 平行赋值
	i,j := 10,20
	fmt.Println("变换前==》",i,j)

	i,j =j,i
	fmt.Println("变换后==》",i,j)
}

func test(a int ,b int)  {
	fmt.Printf("a= %d,b=%d",a,b)

}