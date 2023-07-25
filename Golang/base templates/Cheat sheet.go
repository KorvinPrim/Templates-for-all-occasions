package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
	}
}

func main() {
	nums := [3]float64{5.32, 34.5, 4.21}
	for i, value := range nums {
		fmt.Println(i, value)
	}

	webSites := make(map[string]float64)
	webSites["itProger"] = 0.8
	webSites["yandex"] = 0.9
	fmt.Println(webSites["itProger"])

	var a = 3
	var b = 4
	a, b = summ(a, b)
	fmt.Println(a, b)
}

func summ(num_1 int, num_2 int) (int, int) {
	var res int
	res = num_1 + num_2

	return res, num_1 * num_2
}

Замыкание
func main() {
	var num = 3
	multiple := func() int {
		num *= 2
		return num
	}
	fmt.Println(multiple())
}

Откладывание
func main() {
	defer two()
	one()
}
func one() {
	fmt.Println("1")
}
func two() {
	fmt.Println("2")
}

Указатели
func main() {
	var x = 0
	pointer(&x)
	fmt.Println(x)
}
func pointer(x *int) {
	*x = 2
}

Структуры

func main() {
	bob := Cats{"Bob", 7, 0.87}
	fmt.Println(bob.happines)

	fmt.Println(bob.test())

}

type Cats struct {
	name     string
	age      int
	happines float64
}

func (cat *Cats) test() float64 {
	return float64(cat.age) * cat.happines
}
