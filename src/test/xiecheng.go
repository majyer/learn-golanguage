package main

import (
    "fmt"
    "time"
    "sync"
)

type PipeData struct {
    value int
    handler func(int) int
    next chan int
}

func handle(queue chan *PipeData) {
    for data := range queue {
        data.next <- data.handler(data.value)
    }
}

var flag int32
var once sync.Once

func initialize() {
    flag = 3
    fmt.Println(flag)
}

func setup() {
    once.Do(initialize)
}

func main() {

    c1 := make(chan int)
    c2 := make(chan int)

    go func() {
        for i := 1; i <= 10; i++ {
            <-c1
            go func(i int) {
                fmt.Println(2*i - 1)
                c2<-1
            }(i)

        }
    }()
    go func() {
        for i := 1; i <= 10; i++ {
            <-c2
            go func(i int) {
                fmt.Println(2 * i)
                c1 <- 1
            }(i)
        }
    }()
    c1<-1
    //setup()
    time.Sleep(3 * time.Second)
    setup()
    setup()
}