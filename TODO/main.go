// todo
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const filename = "task.json"

type TASK struct{
	Title string
	Completed bool
}

var tasks []TASK

func main() {
	if len(os.Args) <2 {
		fmt.Println("Usage: go run main.go [add|list|done] [task]")
        return
	}

	command := os.Args[1]

	tasks = loadTask()

	switch command{
	case "add":
		if len(os.Args)<3{
			fmt.Println("please enter your tasks")
			return
		}

		tasks = append(tasks,TASK{Title: os.Args[2]})
		saveTasks(tasks)
		fmt.Println("Task added!")
	
	case "list":
		if len(os.Args) == 0 {
			fmt.Println("No task found")
		}
		for i,v := range tasks{
			status :="[]"
			if v.Completed {
				status ="[X]"
			}
			fmt.Printf("%d. %s %s\n", i+1, status, v.Title)
		}

	case "done":
		if len(os.Args)<3{
			fmt.Println("please select the job id that is done")
			return
		}
		
		index,err :=strconv.Atoi(os.Args[2])
		if err != nil || index <1 || index > len(tasks){
			fmt.Println("Invalid index provided.")
			return
		}
		tasks[index-1].Completed =true
		saveTasks(tasks)
		fmt.Printf("Marked task %d as done!\n", index)
	default :
		fmt.Println("Unknown command:", command)
}
	
}

func loadTask () []TASK {
	data,err := os.ReadFile(filename)
	if err != nil{
		return []TASK{}
	}
	var loadedTask []TASK

	json.Unmarshal(data,&loadedTask)
	return loadedTask
}

func saveTasks(tasks []TASK){
	data,_ := json.MarshalIndent(tasks,""," ")
	os.WriteFile(filename,data,0644)

}