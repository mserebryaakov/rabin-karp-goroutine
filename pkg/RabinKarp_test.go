package pkg_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mserebryaakov/rabin-karp-goroutine/pkg"
)

func readFunc(buff *string, fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Ошибка загрузки файла")
		os.Exit(1)
	}
	fmt.Fscanln(file, buff)
	file.Close()
}

func BenchmarkRabinKarp(b *testing.B) {
	var rb = pkg.RabinKarp{}

	readFunc(&rb.Txt, "text.txt")

	fmt.Println("Текст :", rb.Txt)

	readFunc(&rb.Str, "str.txt")

	fmt.Println("Cтрока :", rb.Str)

	rb.RabinKarpInitialize()
	rb.Start()
}
