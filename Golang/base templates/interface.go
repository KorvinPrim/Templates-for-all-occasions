package main

import (
	"fmt"
	"reflect"
)

/*
Интерфейсы
Интерфейсы описывают некоторое поведение объектов которое мы от него ожидаем.
Для того что бы некий объект или тип данных соответствовал интерфейсу
он должен реализовывать все методы которые объявляет интерфейс
*/
//type animal interface {
//	makeSound()
//}
//
//type cat struct {
//}
//type dog struct {
//}
//
//func (self *cat) makeSound() {
//	fmt.Println("Meow!")
//}
//
//func (self *dog) makeSound() {
//	fmt.Println("Woof!")
//}
//
//func main() {
//	var c, d animal = &cat{}, &dog{}
//	c.makeSound()
//	d.makeSound()
//}

/*
Функции в го умеют принимать в качестве входящих аргументов интерфесы
*/
//type greeter interface {
//	greet(string) string
//}
//
//type russian struct {
//}
//
//type american struct {
//}
//
//type indonezia struct {
//}
//
//func (r russian) greet(name string) string {
//	return fmt.Sprintf("Привет, %s", name)
//}
//
//func (a american) greet(name string) string {
//	return fmt.Sprintf("Hello, %s", name)
//}
//
//func sayHello(g greeter, name string) {
//	fmt.Println(g.greet(name))
//}
//
//func main() {
//	sayHello(&russian{}, "Петя")
//	sayHello(&american{}, "Bob")
//	sayHello(&indonezia{}, "Ila")
//}

/*
Композитные интерфейсы
Помогают придерживать принципа разделения интерфесов (interface segregation)
1 из 5 элементов концепции solid (набор правил и принципов проектирования кода)
На практике позволяет более гранулярно определять то поведение которое
нам требуется внутри какой то функции, в каком-то конкретном месте.
При этом мы не тянем за собой какую-то не нужную информацию.
*/
//
//type animal interface {
//	walker
//	runner
//}
//
//type bird interface {
//	walker
//	flier
//}
//
//type walker interface {
//	walk()
//}
//
//type runner interface {
//	run()
//}
//
//type flier interface {
//	fly()
//}
//
//type cat struct {
//}
//
//type eagle struct {
//}
//
//func (c *cat) walk() {
//	fmt.Println("cat walk")
//}
//
//func (c *cat) run() {
//	fmt.Println("cat is running")
//}
//
//func (e *eagle) walk() {
//	fmt.Println("Eagle walk")
//}
//
//func (e *eagle) fly() {
//	fmt.Println("Eagle is flying")
//}
//
//func walk(w walker) {
//	w.walk()
//}
//
//func main() {
//	var c animal = &cat{}
//	var e bird = &eagle{}
//	walk(c)
//	walk(e)
//}

/*
Пустые интерфейсы
Не описывают не каких методов в плане требований, любой тип данных им соотвецтвует.
Используется когда есть неоднородность или не определённость в тех типах данных
которые используются.
*/

func main() {
	m := map[string]interface{}{}
	m["one"] = 1
	m["two"] = 2.0
	m["three"] = true

	for k, v := range m {
		switch v.(type) {
		case int:
			fmt.Printf("%s is an integer\n", k)
		case float64:
			fmt.Printf("%s is an float\n", k)
		default:
			fmt.Printf("%s is %v\n", k, reflect.TypeOf(v))
		}
	}
}
