package main
import (
"fmt"
"sync"
"runtime"
"time"
)

var counter int =0
func CountChan(ch chan int){
	ch <-1
	fmt.Println("CountChaning")
}
func OutChan(ch chan int){
	<-ch
}
func Count(lock *sync.Mutex){
	lock.Lock()
	counter++
	fmt.Println("counter=",counter)
	lock.Unlock()
}
func Add(x,y int){
	z:=x+y
	fmt.Println(z)
}
func main() {
	lock:=&sync.Mutex{}
	for i:=0;i<10;i++ {
		go Count(lock)
	}
	for{
		lock.Lock()
		c:=counter
		lock.Unlock()
		runtime.Gosched()
		if c>=10{
			break
		}
	}
	for i:=0;i<10;i++{
		go Add(i,i)
	}
	time.Sleep(3 * time.Second)

	chs:=make([] chan int,20)

	for i:=0;i<20;i++{
		chs[i]=make(chan int)
		go CountChan(chs[i])
	}
	
	for _,ch:=range(chs){
		//<-ch
		OutChan(ch)
	}
	time.Sleep(3 * time.Second)
}