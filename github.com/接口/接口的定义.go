package main

import "fmt"

/*
Go语言提倡面向接口编程。

每个接口由数个方法组成，接口的定义格式如下：

type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
其中：

接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
*/

//定义一个接口
type Writer interface {
	Write([]string) error
}

func main() {
	fmt.Println("接口的定义")
}
