package main
import (
"sync"
"fmt"
)

var total struct{
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup){
	defer wg.Done()

	for i=0;i<1000000;i++{
		total.Lock()
		total.value+=1
		//fmt.Println(total.value)
		total.Unlock()
	}
}

func main(){
	var wg sync.WaitGroup
	wg.Add(3)
	go worker(&wg)
	go worker(&wg)
	go worker(&wg)

	wg.Wait()

	fmt.Println(total.value)
}
