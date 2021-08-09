package main

import (
	"fmt"
	"sync"
	"time"
	"unicode"
)

func main() {
	var (
		size  int
		mutex sync.Mutex
		wg    sync.WaitGroup
	)

	fmt.Print("Введите размерность массива: ")
	fmt.Scanln(&size)

	var sem = make(chan int, size)
	defer close(sem)
	var array []string = make([]string, size)

	for i := 0; i < size; i++ {
		fmt.Printf("%d-й элемент: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("Получен массив: ", array)
	fmt.Println("Его размерность: ", size)

	wg.Add(2)
	go work(array, &mutex, sem, &wg)
	<-sem
	go SumElement(array, &mutex, &wg)

	<-sem
	fmt.Print("Итоговый массив: ")
	for i := 0; i < size; i++ {
		<-sem
		fmt.Print(array[i], " ")
	}
	fmt.Println()
	wg.Wait()
	fmt.Scanln()
	fmt.Scanln()
}

func work(array []string, mutex *sync.Mutex, sem chan int, wg *sync.WaitGroup) {
	mutex.Lock()
	sem <- 1
	defer wg.Done()
	fmt.Println("Начало работы потока work...")

	var duration time.Duration
	fmt.Print("Введите временной интервал паузы (мс): ")
	fmt.Scan(&duration)

	sem <- 1
	count := 0
	for index, word := range array {
		if unicode.IsDigit(rune(word[0])) {
			array[count] = array[index]
			count++
			sem <- 1
		}
	}

	for ; count < len(array); count++ {
		array[count] = "_"
		sem <- 1
		time.Sleep(duration * time.Millisecond)
	}

	fmt.Println("Конец работы потока work...")
	mutex.Unlock()
}

func SumElement(array []string, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	defer wg.Done()

	fmt.Println("Начало работы потока SumElement")
	sum := 0

	for _, word := range array {
		for _, char := range word {
			sum += int(char)
		}
	}

	fmt.Println("Сумма кодов массива: ", sum)
	mutex.Unlock()
}
