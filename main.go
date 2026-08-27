package main

import "fmt"

type Task struct {
	ID 		int
	Title       string
	Completed   bool
}

func main() {
	task1 := Task{ID: 1, Title: "Learn Go", Completed: false}
	task2 := Task{ID: 2, Title: "Learn Kubernetes", Completed: false}
	var task3 Task
	fmt.Println(task1.Title, task1.Completed)
	fmt.Println(task2.Title, task2.Completed)
	fmt.Println(task3) 
}