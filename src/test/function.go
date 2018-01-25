package main

import(
    //"fmt"
    // "io"
    // "os"
    // "bytes"
)

// type io.Writer interface {
//     Write(p []byte) (n int, err error)
// }

// type error interface {
//     Error() string
// }

// type UpperWriter struct {
//     io.Writer
// }

// func (p *UpperWriter) Write(data []byte) (n int, err error) {
//     return p.Writer(bytes.ToUpper(data))
// }

func main() {
    for i := 0; i < 3; i++ {
        // 通过函数函数传入i
        // defer 语句会马上对调用参数求值
        defer func(i int){ println(i) } (i)
    }

    for i := 0; i < 3; i++ {
        defer func(){ println(i) } ()
    }

    for i := 0; i < 3; i++ {
        i := i // 定义一个循环体内局部变量i
        defer func(){ println(i) } ()
    }
    
    //fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")

}