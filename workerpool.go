package main

import (
	"fmt"
	"sync"
	"time"
)

// Mendefinisikan Task yang memiliki id bertipe int
type Task struct {
	Id int
}

// Membuat method Process untuk mengeksekusi task yang ada
func (t *Task) Process() {
	fmt.Printf("Processing task %d \n", t.Id)
	// Mensimulasikan waktu process dalam 2 detik
	time.Sleep(2 * time.Second)
}

// Mendefinisikan WorkerPool untuk mengatur sekumpulan worker
// yang akan menjalankan sebuah task secara concurrent
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	taskChan    chan Task
	wg          sync.WaitGroup
}

// Method worker fungsinya dijalankan oleh setiap worker untuk memproses
// task yang diterima melalui channel taskChan
func (wp *WorkerPool) worker() {
	for task := range wp.taskChan {
		task.Process()
		wp.wg.Done()
	}
}

// Method Run
func (wp *WorkerPool) Run() {
	// Menginisialisasi channel taskChan dengan kapasitas sebesar jumlah task yang ada.
	wp.taskChan = make(chan Task, len(wp.taskChan))

	// Menjalankan sejumlah goroutine yang menjalankan fungsi worker sebanyak jumlah concurrency
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	// Menambah counter sync.WaitGroup sebanyak jumlah task yang akan diproses.
	wp.wg.Add(len(wp.Tasks))

	// Mengirim setiap proses dari task channel ke taskChan
	for _, task := range wp.Tasks {
		wp.taskChan <- task
	}

	// Meng-close channel taskChan setelah semua task dikirim, sehingga menghindari dari deadlock
	// dan berfungsi untuk memberi sinyal bahwa tidak ada goroutine yang akan di proses
	defer close(wp.taskChan)

	//Menunggu semua proses sampai selesai
	wp.wg.Wait()
}
