package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var (
	myName string
)

/*
int, strings, booleans, error, floats
arrays, slices, maps

	control flow
	for loop, if statement, switch statement
*/

func AddInt(x, y int) (int, string) {
	result := x + y
	FuncName := "Addition function"
	return result, FuncName
}

type Speaker interface {
	Sound() string
	Movement() string
}

type Animal struct {
	Breed  string
	Sounds string
}

func NewAnimal(name, sound string) *Animal {
	anAnimal := Animal{
		Breed:  name,
		Sounds: sound,
	}
	return &anAnimal
}

func (a *Animal) Sound() string {
	madeSound := "A " + a.Breed + " " + a.Sounds + "s"
	return madeSound
}

func printNUmbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	for i := 'a'; i < 'f'; i++ {
		fmt.Printf("%c\n", i)
	}
}

func main() {

	// var mySLice = []string{"0", "1", "2", "3", "4"}
	// for _, item := range mySLice {
	// 	fmt.Println(item)
	// }

	// var wg sync.WaitGroup

	// wg.Add(2)
	// go printNUmbers(&wg)
	// go printLetters(&wg)

	// wg.Wait()

	// NewResutl, function := AddInt(5, 4)
	// fmt.Println(NewResutl, function)
	// myFloat := float64(NewResutl)
	// fmt.Printf("my float: %f, is of type %T\n", myFloat, myFloat)

	// dog := NewAnimal("German Sheperd", "Bark")
	// animalSound := dog.Sound()
	// fmt.Println(animalSound)

	// cat := NewAnimal("Brownie", "Meow")
	// fmt.Println(cat.Sound())

	// myNum := 12
	// myColor := "Red"
	// isTherelight := true
	// myError := errors.New("Type my error")
	// passMark := 67.7
	// myRune := 'e'
	// fmt.Printf("%T", myRune)
	// fmt.Printf("%T", myNum)
	// fmt.Printf("%T", myColor)
	// fmt.Printf("%T", isTherelight)
	// fmt.Printf("%T", myError)
	// fmt.Printf("%T", passMark)

	// var myArr [2]int

	// myArr[0] = 10
	// myArr[1] = 20
	// fmt.Println(myArr)

	// // var ColorArray = [3]string {"red", "blue", "purple"}

	// var mySlice = []int{1, 2, 3, 4}
	// fmt.Print(len(mySlice))

	// anotherSlice := make([]int, 0)
	// anotherSlice = append(anotherSlice, 1, 2, 3)
	// fmt.Println(anotherSlice)

	// anotherMap := make(map[int]string)
	// anotherMap[0] = "Cradle"
	// anotherMap[1] = "birth"
	// fmt.Println(anotherMap)

	// for key, value := range anotherMap {
	// 	fmt.Println(key, ":", value)
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// // var myAge int
	// // fmt.Print("Enter your age: ")
	// // fmt.Scanf("%d", &myAge)

	// // if myAge < 18 {
	// // 	fmt.Println("You are not of legal age to drive")
	// // } else {
	// // 	fmt.Println("Go and get your driver's license!")
	// // }

	// var myOption string
	// fmt.Println("enter your option: ")
	// fmt.Scanf("%s", &myOption)

	// switch myOption {
	// case "r":
	// 	fmt.Println("your option r")

	// case "b":
	// 	fmt.Println("your option is b")

	// default:
	// 	fmt.Println("Your option is not valid")
	// }

	_, err := os.Create("file.txt")
	if err != nil {
		log.Println(err)
	}
}
