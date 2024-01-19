package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("Please enter a valid number to increase")
	// reader := bufio.NewReader(os.Stdin)
	// readerValue, _ := reader.ReadString('\n')
	// fmt.Println("Here you entered: ", readerValue)
	// numVal, err := strconv.ParseFloat(strings.TrimSpace(readerValue), 64)
	// if err != nil {
	// 	fmt.Println("Error occurred:", err)
	// } else {
	// 	fmt.Println("Number converted into number from string:", numVal+1)
	// }

	fmt.Println("Enter a valid number to get a feedback result")
	var getFeedback = bufio.NewReader(os.Stdin)
	feed, err := getFeedback.ReadString('\n')

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("loop is running", i)
	// }

	if err != nil {
		panic("Error occurred")
	} else {
		parsedValue, _ := strconv.ParseFloat(strings.TrimSpace(feed), 64)
		fmt.Println("feedback value: ", parsedValue)
	}

	// slices excersize
	var fruitList = []string{"Mango", "Banana", "pineapple", "guava", "papaya"}
	var veggieList [2]string
	fruitList = append(fruitList[:3], fruitList[4:]...)
	fmt.Printf("Variable type for fruitList is %v\n", fruitList)
	fmt.Printf("Variable type for veggieList is %v\n", veggieList)

	// trying to dig deeper into array
	classArray := make([]string, 4)
	addr := &classArray
	addr23 := &classArray
	classArray = append(classArray, "class 1", "class 2", "class 3", "class 4")
	addr1 := &classArray
	fmt.Println("Address for classArray", &addr, &addr23)
	fmt.Println("Address for classArray", &addr1)

}
