package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	The idea behind a key-value store is modest:
	answer queries fast and work as fast as possible.

	ایده ی \شت یک  کی ولیو استور ساده این هست که شما به کوعری ها سزیع پاسخ بدهید و تا حد امکان سریع کار کنید

	 This translates into using simple algorithms and simple data structures.

	 این باعث میشه که شما از الگوریتم ها و دیتا استراکچر های ساده استفاده کنید

*/

/*
	این برنامه  ۴ وظیفه ی اساسی یک  کی ولیو استور رو برای ما انجام میدهد  : به صورت زیر

	1 . adding a new element
	2 . deleting an existing element from key-value store based on a key
	3 . Looking up the value of specific key
	4 . changing the value of an existing key

*/

// ADD DELETE CHANGE LOOKUP STOP PRINT

//type of elements
type Item struct {
	Name    string
	SurName string
	Id      string
}

var DATA = make(map[string]Item)

func ADD(key string, value Item) bool {

	if key == "" {
		return false
	}
	if LOOKUP(key) == nil {

		DATA[key] = value

		return true

	}

	return false

}

func DELETE(key string) bool {

	if LOOKUP(key) != nil {

		delete(DATA, key)

		return true

	}

	return false

}

func LOOKUP(key string) *Item {

	_, ok := DATA[key]

	if ok {
		element := DATA[key]

		return &element
	} else {

		return nil
	}

}

func CHANGE(key string, newValue Item) bool {

	DATA[key] = newValue

	return true

}

func PRINT() {

	for key, value := range DATA {

		fmt.Printf("key: %v  value: %v\n", key, value)

	}

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		text := scanner.Text()

		text = strings.TrimSpace(text)

		tokens := strings.Fields(text)

		switch len(tokens) {
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "ADD":
			element := Item{
				Name:    tokens[2],
				SurName: tokens[3],
				Id:      tokens[4],
			}

			if !ADD(tokens[1], element) {

				fmt.Println("Add operation failed!")
			}

		case "LOOKUP":
			element := LOOKUP(tokens[1])

			if element != nil {
				fmt.Printf("%v\n", element)
			}
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Println("Delete operation failed!")
			}

		case "CHANGE":
			element := Item{
				Name:    tokens[1],
				SurName: tokens[2],
				Id:      tokens[3],
			}

			if !CHANGE(tokens[1], element) {
				fmt.Println("Update operation failed!")

			}

		case "PRINT":
			PRINT()
		case "STOP":
			return
		default:
			fmt.Println("Unknown command = please try again!")

		}

	}

}
