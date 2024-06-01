package main

import "fmt"

func main() {
	// Membuat slice tasks berisi 10 Task dengan ID dari 0 sampai 9.
	tasks := make([]Task, 10)

	for i := 0; i < 10; i++ {
		tasks[i] = Task{Id: i}
	}

	// Membuat instance dari WorkerPool dengan daftar task yang telah dibuat dan tingkat konkuren sebesar 5.
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5,
	}

	wp.Run()
	fmt.Println("All tasks have been processed")
}
