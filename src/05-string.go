package main

import "fmt"

func main()  {
	name := "flw"

	// 需要换行，原生输出字符串时候，使用反引号 ``
	usage :=`hello
            -h 
            -a ....`

	fmt.Println("name->",name)
	fmt.Println("usage->",usage)

	// 2:长度 没有.length
	l := len(name)
	fmt.Println("name的长度->",l)
	for i :=0; i<len(name);i++ {
		fmt.Println(i)
	}

	// 拼接
	i,j :="hello","world"
	fmt.Println("i+j->",i+j)

	// 使用const 修饰的常量，不能修改

}
