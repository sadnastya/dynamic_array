package main

import (
	"fmt"
	"strings"
)

type MyMas struct {
	myArray *[]int
}

func (array *MyMas) addToEnd(n int) {
	var object MyMas = *array
	var mainArray []int = *object.myArray
	if len(mainArray) == cap(mainArray) {
		var newCapacity int = len(mainArray) * 2
		slice := make([]int, 0, newCapacity)
		for i := 0; i < len(mainArray); i++ {
			slice = append(slice, mainArray[i])
		}
		slice = append(slice, n)
		array.myArray = &slice
	} else {
		mainArray[len(mainArray)+1] = n
	}
}

func (array *MyMas) addToTop(n int) {
	var object MyMas = *array
	var mainArray []int = *object.myArray
	if len(mainArray) == cap(mainArray) {
		var newCapacity int = len(mainArray) * 2
		slice := make([]int, 0, newCapacity)
		slice = append(slice, n)
		for i := 1; i < len(mainArray); i++ {
			slice = append(slice, mainArray[i])
		}
		array.myArray = &slice
	} else {
		var k int = mainArray[0]
		for i := 1; i < len(mainArray); i++ {
			k, mainArray[i] = mainArray[i], k
		}
		mainArray = append(mainArray, k)
		mainArray[0] = n
		array.myArray = &mainArray
	}
}

func (array *MyMas) deleteTheFirst() {
	var object MyMas = *array
	var mainArray []int = *object.myArray
	var newArray []int = make([]int, len(mainArray)-1, cap(mainArray))
	for i := 1; i < len(mainArray); i++ {
		newArray[i-1] = mainArray[i]
	}
	array.myArray = &newArray
}

func (array *MyMas) deleteLast() {
	var object MyMas = *array
	var mainArray []int = *object.myArray
	var newArray []int = make([]int, len(mainArray)-1, cap(mainArray))
	for i := 0; i < len(mainArray)-1; i++ {
		newArray[i] = mainArray[i]
	}
	array.myArray = &newArray
}

func (array *MyMas) clear() {
	var structure MyMas = *array
	slice := make([]int, 0, cap(*structure.myArray))
	array.myArray = &slice
}

func (array *MyMas) length() int {
	var structure MyMas = *array
	var mainArray []int = *structure.myArray
	return len(mainArray)
}

func (array *MyMas) capacity() int {
	var structure MyMas = *array
	var mainArray []int = *structure.myArray
	return cap(mainArray)
}

func (array *MyMas) toString() string {
	var structure MyMas = *array
	var mainArray []int = *structure.myArray
	string := fmt.Sprint(mainArray)
	res := strings.Split(string, " ")
	return strings.Join(res, ", ")
}

func main() {
	mas := []int{1, 2, 3}
	array := MyMas{myArray: &mas}
	fmt.Println(*array.myArray)
	array.addToEnd(10)
	fmt.Println(*array.myArray)
	array.addToTop(0)
	fmt.Println(*array.myArray)
	var stringArray string = array.toString()
	fmt.Println(stringArray)
	fmt.Println(array.length())
	fmt.Println(array.capacity())
	array.deleteTheFirst()
	fmt.Println(*array.myArray)
	array.deleteLast()
	fmt.Println(*array.myArray)
	array.clear()
	fmt.Println(*array.myArray)
}
