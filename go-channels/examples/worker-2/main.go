package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for el := range in {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Эмуляция работы
		out <- el * 2                                                // Возвращаем результат
	}
}

func main() {
	var limWorkers = 5
	var limJobs = 10

	jobs := make(chan int, limJobs)    // Канал для задач
	results := make(chan int, limJobs) // Канал для результатов

	wg := &sync.WaitGroup{}

	for i := 0; i < limWorkers; i++ {
		wg.Add(1)
		go worker(jobs, results, wg)
	}

	// Fan-Out: Отправляем задачи в канал
	for j := 0; j < limJobs; j++ {
		jobs <- j
	}
	close(jobs) // Закрываем канал, чтобы воркеры знали, что задач больше не будет

	// Ожидаем завершения всех воркеров
	go func() {
		wg.Wait()
		close(results) // Закрываем канал результатов
	}()

	// Fan-In: Собираем результаты
	for result := range results {
		fmt.Println("results:", result)
	}
}
