package main

import "fmt"

func hash(f, s []string) {
	Str := make([]string, 0, 0)
	Result := make(map[string]int)
	for _, x := range f {
		for _, y := range s {
			if x == y {
				Result[y] += 1
			}
		}
	}
	for key, _ := range Result {
		Str = append(Str, key)
	}
	if len(Str) == 0 {
		fmt.Println("Нет одинаковых элементов")
	} else {
		fmt.Println(Str)
	}
}

func main() {
	var s string
	var firstCap, secondCap int
	fmt.Println("Введите размер первого массива:")
	fmt.Scan(&firstCap)
	var first []string
	fmt.Println("Введите размер второго массива:")
	fmt.Scan(&secondCap)
	var second []string
	fmt.Println("теперь введите элементы первого массива через Enter")
	for i := 0; i < firstCap; i++ {
		fmt.Scan(&s)
		first = append(first, s)
	}
	fmt.Println("теперь введите элементы второго массива через Enter")
	for i := 0; i < secondCap; i++ {
		fmt.Scan(&s)
		second = append(second, s)
	}
	hash(first, second)
}
