package main

import (
	"github.com/baddin/codister"
	"fmt"
)


func main() {
	str := "hrllo kdjfkkk func main(){} ksjkd func hello() string {} skkskks func hi() []int  sdsdjjvyfgfdgdfgfdgu fdgdfgfd func test() int {} ri"
	for _, v := range codister.GetFuncs(str) {
		fmt.Println(v)
	}

}
