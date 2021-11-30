package main

import (
	"fmt"
	"sync"
	"time"
)
func reloadTags(wg *sync.WaitGroup,t chan int, cnt int){
	
	cnt +=1
	fmt.Println(cnt)
	
	fmt.Println("Refreshing/reloading tickers after every (say) 30 seconds")
	if cnt >= 5{
		fmt.Println("DONE")
		wg.Done()
		close(t)
	}
	if cnt <5 {
		t <- cnt
	}
	
}

func main(){
	c := make(chan int) // initializing channel
	wg := &sync.WaitGroup{}
	cnt := 0
	wg.Add(1)
	go reloadTags(wg,c,cnt)
	for {
		cnt,open := <- c
		if !open {
			break
		}
		fmt.Println("in for")
		go func(cnt int){
			time.Sleep(30 * time.Second)
			// if get is called => fire reloadTags
			reloadTags(wg,c,cnt)
	}(cnt)
}	
	wg.Wait()
}	