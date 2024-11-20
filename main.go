package test

import "fmt"

func Printy(data []string) {
	for _, val := range data {
		go func() {
			fmt.Println(val)
		}()
	}
}
