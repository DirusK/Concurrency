package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("Введите размерность массива: ")
	var size int
	fmt.Scan(&size)

	array := make([]int, size)
	fmt.Print("Введите элементы массива: ")
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println("Полученный массив: ", array)

	var channel chan float64 = make(chan float64)
	go worker(array, channel)

	min, max := MinAndMax(array...)
	fmt.Printf("Минимальный элемент в массиве: %d\nМаксимальный элемент в массиве: %d\n", min, max)

	average := <-channel
	fmt.Println("Среднее значение в массиве: ", average)

	fmt.Print("Элементы массива, значения которых больше среднего: ")
	count := 0
	for _, value := range array {
		if float64(value) > average {
			fmt.Print(value, " ")
			count++
		}
	}
	fmt.Println("\nВсего: ", count, " элементов")

	channel <- float64(min)

	fmt.Println("Сумма  нечетных  элементов  массива  и минимального элемента: ", <-channel)

	fmt.Scanln()
	fmt.Scanln()
}

func worker(array []int, channel chan float64) {

	var (
		sumAll float64 = 0
		sumNeg float64 = 0
	)

	for _, value := range array {
		sumAll += float64(value)

		if value % 2 !=0 {
			sumNeg += float64(value)
		}

		time.Sleep(12 * time.Millisecond)
	}

	average := sumAll / float64(len(array))
	channel <- average

	min := <-channel
	sumNeg += min
	channel <- sumNeg

}

func MinAndMax(number ...int) (min int, max int) {
	min = number[0]
	max = number[0]
	for _, value := range number {

		if min > value {
			min = value
		}
		if max < value {
			max = value
		}

		time.Sleep(14 * time.Millisecond)
	}
	return
}
