package pkg

import (
	"fmt"
	"sync"
)

//Количество запускаемых горутин (деление текста на количество горутин)
const numThread int = 2

type RabinKarp struct {
	Str, Txt        string
	pPower, hashs_t []int
	hash_s          int
}

//Расчет степеней
func (rb *RabinKarp) initializePPower() {
	const P = 31
	(*rb).pPower[0] = 1
	for i := 1; i < len((*rb).pPower); i++ {
		(*rb).pPower[i] = (*rb).pPower[i-1] * P
	}
}

//Инициализация переменных
func (rb *RabinKarp) RabinKarpInitialize() {
	rb.hashs_t = make([]int, len((*rb).Txt))
	rb.pPower = make([]int, len((*rb).Txt))
}

//Расчет хэшей текста
func (rb *RabinKarp) createHashT(wg *sync.WaitGroup) {
	for i := 0; i < len((*rb).Txt); i++ {
		(*rb).hashs_t[i] = (int((*rb).Txt[i]) - 97 + 1) * (*rb).pPower[i]
		if i > 0 {
			(*rb).hashs_t[i] += (*rb).hashs_t[i-1]
		}
	}
	wg.Done()
}

//Расчет хэша строки
func (rb *RabinKarp) createHashS(wg *sync.WaitGroup) {
	for i := 0; i < len((*rb).Str); i++ {
		(*rb).hash_s += (int((*rb).Str[i]) - 97 + 1) * (*rb).pPower[i]
	}
	wg.Done()
}

//Вывод результата
func (rb *RabinKarp) createResult(forWhere int, toWhere int, wg *sync.WaitGroup) {
	result := false
	fmt.Println("\nРезультат: ")
	for i := forWhere; i+len((*rb).Str)-1 < toWhere; i++ {
		cur_h := (*rb).hashs_t[i+len((*rb).Str)-1]
		if i > 0 {
			cur_h -= (*rb).hashs_t[i-1]
		}
		if cur_h == (*rb).hash_s*(*rb).pPower[i] {
			fmt.Printf("%v ", i)
			result = true
		}
	}
	if !result {
		fmt.Print("Вхождений не найдено")
	}
	wg.Done()
}

//Запуск алгоритма
func (rb *RabinKarp) Start() {
	(*rb).initializePPower()
	var wg sync.WaitGroup
	wg.Add(2)
	go (*rb).createHashT(&wg)
	go (*rb).createHashS(&wg)
	wg.Wait()

	forWhere := 0
	toWhere := len(rb.Txt) / numThread
	step := toWhere
	toWhere += len(rb.Txt) % numThread
	wg.Add(numThread)
	for i := 0; i < numThread; i++ {
		go (*rb).createResult(forWhere, toWhere, &wg)
		forWhere = toWhere - len(rb.Str) + 1
		toWhere += step
	}
	wg.Wait()
}
