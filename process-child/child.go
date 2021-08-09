package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type emp struct {
	Num   int
	Name  string
	Grade float64
}

func find(list []emp, name string) (int, error) {
	for index, value := range list {
		if value.Name == name {
			return index, nil
		}
	}
	return -1, errors.New("Зачетки с такой фамилией не существует")
}

func main() {
	var (
		list  []emp
		new   emp
		scaner = bufio.NewScanner(os.Stdin)
	)

	file, err := os.OpenFile("D:\\Golang\\Projects\\OS_LB3\\document.txt", os.O_RDWR, 0755)
	if err != nil { // если возникла ошибка
		fmt.Println("Не получается открыть файл с зачётками:", err)
		os.Exit(1) // выходим из программы
	}
	defer file.Close()

	for {
		_, err = fmt.Fscanf(file, "%d - %s - %f \n", &new.Num, &new.Name, &new.Grade)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			list = append(list, new)
		}
	}

	var surname string
	fmt.Print("Введите фамилию студента: ")
	scaner.Scan()
	surname = scaner.Text()

	index, err := find(list, surname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}else{
		fmt.Println("Найдена следующая зачётка: ")
		fmt.Printf("%d - %s - %.1f \n", list[index].Num, list[index].Name, list[index].Grade)
		
	}

	fmt.Print("Введите новый бал студента: ")
	fmt.Scanln(&list[index].Grade)

	file.Truncate(0)
	file.Seek(0,0)
	for _,value:=range list{
		fmt.Fprintf(file, "%d - %s - %f \n", value.Num, value.Name, value.Grade)
	}
	
	fmt.Println("Нажмите кнопку ENTER для продолжения...")
	fmt.Scanln()
}
