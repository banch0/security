package main

import (
	"errors"
	"log"
)

func main() {
	abc := false
	// code, data := twoReturn()
	code, data, err := checkV(abc)
	log.Println(err)
	log.Println(code)
	log.Println(data)
	data = "new things"
	log.Println(data)
}

func checkV(abc bool) (code int, data interface{}, err error) {
	if abc {
		code = 234
		data = "abc equal true"
	} else {
		code = 332
		data = "abc not equal to true"
		err = errors.New("some error ")
	}
	log.Println("ehhil")
	twoReturn()
	return
}

func twoReturn() (code int, data string) {
	log.Println("dother")
	code = 404
	data = "Old things"
	return
}
