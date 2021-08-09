package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"io/ioutil"
)

type emp struct {
	Num   int
	Name  string
	Grade float64
}



func main() {

	var (
		choice int
		list   []emp
		new    emp
		scaner = bufio.NewScanner(os.Stdin)
	)

	file, err := os.OpenFile("D:\\Golang\\Projects\\OS_LB3\\document.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
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

	for {

		fmt.Println("\n~~~ МЕНЮ ПРОГРАММЫ ~~~")
		fmt.Println("1 - Вывести все зачётки на экран")
		fmt.Println("2 - Добавить зачётку в файл")
		fmt.Println("3 - Отредактировать нужную зачётку")
		fmt.Println("4 - Выйти из программы")

		fmt.Print("\nВаше действие: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if len(list) == 0 {
				fmt.Println("Зачёток в файле нет!")
			} else {
				for _, value := range list {
					fmt.Printf("%d - %s - %.1f \n", value.Num, value.Name, value.Grade)
				}
			}

		case 2:
			fmt.Print("Ввведите имя студента: ")
			scaner.Scan()
			new.Name = scaner.Text()

			fmt.Print("Ввведите оценку студента: ")
			fmt.Scanln(&new.Grade)

			new.Num = len(list) + 1
			fmt.Println("Зачётка получила номер: ", new.Num)

			list = append(list, new)
			fmt.Fprintf(file, "%d - %s - %f \n", new.Num, new.Name, new.Grade)

		case 3:
			
			cmd := exec.Command("cmd.exe","/C", "start", "D:\\Golang\\Projects\\OS_LB3\\process-child\\process-child.exe")

			stdoutPipe, err := cmd.StdoutPipe()
			if err != nil {
			  fmt.Println(err)
			}

			cmd.Start()

			_, err = ioutil.ReadAll(stdoutPipe)
			if err != nil {
			  fmt.Println("Stdout Pipe error: ",err)
			}

			cmd.Wait()

			list = nil
			file.Seek(0,0)
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
			fmt.Println("Данные сохранены!")

		case 4:
			os.Exit(1)
		default:
			fmt.Println("Такой кнопки нет!")
		}
	}
}
