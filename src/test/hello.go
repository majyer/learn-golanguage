package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
  fmt.Printf("hello, world\n")
  var a = [...]int{3,1, 2,7, 4: 5, 6} // a 是一个数组
  var b = &a                // b 是指向数组的指针

  fmt.Println(a[0], a[1])   // 打印数组的前2个元素
  fmt.Println(b[0], b[1])   // 通过数组指针访问数组元素的方式和数组类似

  for i, v := range b {     // 通过数组指针迭代数组的元素
    fmt.Println(i, v)
   }

  //var data = [...]byte{'h', 'e', 'l', 'l', 'o', ',' ' ', 'w', 'o', 'r', 'l', 'd'}
  s := "hello, world"
  hello := s[:5]
  world := s[7:]

  fmt.Println("s is: ",s)
  fmt.Println("hello is: ",hello)
  fmt.Println("world is: ",world)

  s1 := "hello, world"[:5]
  s2 := "hello, world"[7:]
   
  fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
  fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
  fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5

  fmt.Printf("%#v\n", []byte("Hello, 世界"))

  fmt.Println("\xe4\xb8\x96") // 打印: 世
  fmt.Println("\xe7\x95\x8c") // 打印: 界
  
}

