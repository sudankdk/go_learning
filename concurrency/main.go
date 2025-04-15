package main

import (
	"fmt"
	"time"
)

// func main(){
// 	now := time.Now()
// 	respCh := make(chan string,2)

// 	wg := &sync.WaitGroup{}

// 	id:= 20
// 	go UserData(id,respCh,wg)
// 	go FriendSuggestion(id,respCh,wg)
// 	wg.Add(2)
// 	wg.Wait()

// 	close(respCh)

// 	for res := range respCh{
// 		fmt.Println(res)
// 	}

// 	fmt.Println("time taken ", time.Since(now))
// }

// func UserData(id int,respCh chan string,wg *sync.WaitGroup){
// 	time.Sleep(50*time.Millisecond)
// 	respCh <- "user datat"
// 	wg.Done()

// }

// func FriendSuggestion(id int,respCh chan string, wg *sync.WaitGroup){
// 	time.Sleep(100*time.Millisecond)
// 	respCh <- "frnds suggg"
// 	wg.Done()
// }

type Message struct{
	From string
	Payload string
}

type Server struct{
	msgch chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen(){
	fmt.Println("Listening to the server")
	free: 
	for {
		select{
		case msg := <- s.msgch:
		fmt.Printf("the messsge is sent by %s and payload is %s",msg.From,msg.Payload)
		case <-s.quitch:
			fmt.Println("shutdown quite gracefully")
			break free
		default :
		
	}
	}
}

func SendMessageToServer(msg chan Message, payload string){
	message := Message{
		From: "sudan don",
		Payload: payload,
	}
	msg<- message


	fmt.Println("sending message to the server")
}

func ShutdownServer(quit chan struct{}){
	close(quit)
}

func main(){
	s := &Server{
		msgch: make(chan Message),
		quitch: make(chan struct{}),
	}

	go s.StartAndListen()
	SendMessageToServer(s.msgch,"Hohoho")
	go func() {
		ShutdownServer(s.quitch)
	}()

	time.Sleep(2 * time.Second)
	select{}
}




