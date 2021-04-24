package main

import (
	"fmt"
	"os"

	"github.com/mserebryaakov/rabin-karp-goroutine/pkg"
)

//Функция чтения строки из файла
func readFunc(buff *string, fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Ошибка загрузки файла")
		os.Exit(1)
	}
	fmt.Fscanln(file, buff)
	file.Close()
}

func main() {
	var rb = pkg.RabinKarp{}

	readFunc(&rb.Txt, "text.txt")

	fmt.Println("Текст :", rb.Txt)

	readFunc(&rb.Str, "str.txt")

	fmt.Println("Cтрока :", rb.Str)

	rb.RabinKarpInitialize()
	rb.Start()
}
