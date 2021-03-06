package main

import "fmt"

/*
区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。

要搞明白Go语言中的指针需要先知道3个概念：指针地址、指针类型和指针取值。

Go语言中的指针
任何程序数据载入内存后，在内存都有他们的地址，这就是指针。而为了保存一个数据在内存中的地址，我们就需要指针变量。

比如，“永远不要高估自己”这句话是我的座右铭，我想把它写入程序中，程序一启动这句话是要加载到内存（假设内存地址0x123456），我在程序中把这段话赋值给变量A，把内存地址赋值给变量B。这时候变量B就是一个指针变量。通过变量A和变量B都能找到我的座右铭。

Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：&（取地址）和*（根据地址取值）。

指针地址和指针类型
每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用&字符放在变量前面对变量进行“取地址”操作。 Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等。

取变量指针的语法如下：

ptr := &v    // v的类型为T
其中：

v:代表被取地址的变量，类型为T
ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。
*/

func main() {
	a := 520
	b := &a
	fmt.Printf("a:%d %p\n", a, &a)     //a:520 0xc0000140a8
	fmt.Printf("b:%p Type:%T\n", b, b) //b:0xc0000140a8 Type:*int
	fmt.Println(&b)                    //0xc0000d8020
}
