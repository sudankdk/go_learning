package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	now := time.Now()
	respCh := make(chan string,2)

	wg := &sync.WaitGroup{}




	id:= 20
	go UserData(id,respCh,wg)
	go FriendSuggestion(id,respCh,wg)
	wg.Add(2)
	wg.Wait()

	close(respCh)

	for res := range respCh{
		fmt.Println(res)
	}


	fmt.Println("time taken ", time.Since(now))
}

func UserData(id int,respCh chan string,wg *sync.WaitGroup){
	time.Sleep(50*time.Millisecond)
	respCh <- "user datat"
	wg.Done()
	
}

func FriendSuggestion(id int,respCh chan string, wg *sync.WaitGroup){
	time.Sleep(100*time.Millisecond)
	respCh <- "frnds suggg"
	wg.Done()
}