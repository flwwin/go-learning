package main

import "fmt"

func main() {
	// go语义也有指针
	// c语言 ptr->name go: ptr,name

	// go语言在使用指针的时候，会使用内部的垃圾回收机制（gc：garbage collector），开发人员不需要手动释放内存
	// c语言不允许返回栈上指针，编译的时候就确定了分配的位置
	// 编译的时候，如果发现有必要的话，将内部的变量分配到堆上

	name := "lily"

	// 定义要给指针 指向name
	prt := &name

	fmt.Println("name->",*prt)
	fmt.Println("name ptr",prt)

	// 使用new 关键字定义指针
	name2prt :=new(string)

	*name2prt = "lily"

	fmt.Println("name2->",*name2prt)
	fmt.Println("name2 ptr:",name2prt)

	// 可以返回栈上的指针，编译器在编译程序的时候，会自动判断这段代码，将city变量分配在堆上

	res := testptr()
	fmt.Println("res city->",*res,"address->",res)

	// 空指针 nil
	if res == nil {
		fmt.Println("res is nil")
	}else {
		fmt.Println("res is not nil")
	}
}

// 定义要给函数，返回一个string类型的指针
func testptr() *string {
	city :="shenzhen"
	ptr := &city
	return ptr
}