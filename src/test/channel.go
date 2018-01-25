package main
import "fmt"

var done=make(chan bool)

var msg string

var limit=make(chan int,3)
func aGoroutine(){
	msg="hello world"
	done <-true
}
func main() {
	go aGoroutine()
	<-done
	fmt.Println(msg)

	for _,w:=range work{
		go func(){
			limit<-1
			w()
			<-limit
		}()
	}
	select{}
}