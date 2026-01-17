package main

import (
	"fmt"
	"time"
)


func printMessage(id int) {
	fmt.Printf("Привіт з потоку №%d\n", id)
}

func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("Відправлено %d у канал\n", i)
	}
	close(ch)
}

func bufferedChannelDemo() {
	fmt.Println("\n=== Буферизований канал ===")
	ch := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("Відправлено %d\n", i)
	}


	for i := 0; i < 5; i++ {
		val := <-ch
		fmt.Printf("Отримано %d\n", val)
	}
}

func selectDemo() {
	fmt.Println("\n=== Select і таймер ===")
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Повідомлення після 2 секунд"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ch <- "Повідомлення після 4 секунд"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Отримано:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Таймаут: минуло 3 секунди")
	}
}

func worker(id int, in <-chan int, out chan<- int) {
	for n := range in {
		square := n * n
		fmt.Printf("Worker %d: %d² = %d\n", id, n, square)
		out <- square
	}
}

func fanOutDemo() {
	fmt.Println("\n=== Goroutine-пул (Fan-out) ===")

	// Дані для варіанта 12
	nums := []int{41, 20, 26, 11, 46, 48, 5, 6, 29, 12}

	in := make(chan int)
	out := make(chan int)

	for i := 1; i <= 3; i++ {
		go worker(i, in, out)
	}

	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()

	for i := 0; i < len(nums); i++ {
		res := <-out
		fmt.Printf("Результат: %d\n", res)
	}
}

func main() {
	fmt.Println("=== 1. Горутинa ===")
	go printMessage(1)
	go printMessage(2)
	go printMessage(3)

	time.Sleep(time.Second)

	fmt.Println("\n=== 2. Канал ===")
	ch := make(chan int)
	go sendNumbers(ch)
	for val := range ch {
		fmt.Printf("Отримано %d з каналу\n", val)
	}

	bufferedChannelDemo()
	selectDemo()
	fanOutDemo()
}
