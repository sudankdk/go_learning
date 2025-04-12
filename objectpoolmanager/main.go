package main

type Worker struct{
	Id int
	Job string

}

type Pool struct{
	worker []*Worker
	inUse map[*Worker]bool
}


func newPool (size int ) *Pool{
	p := &Pool{
		worker: make([]*Worker,size),
		inUse: make(map[*Worker]bool),
	}
	for i:=0; i<size; i++{
		p.worker[i]= &Worker{Id:i}
	}
	return p
}

func acquire(p *Pool) *Worker{
	for i:= range(len(p.worker)){
		if 
	}
}

func relese(){

}

func main(){
	
}