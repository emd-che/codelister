package main

import (
	"flag"
	"fmt"
	"github.com/baddin/codelister"
	"log"
	"io/ioutil"
)

func readFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bytes)
	return str

}

func main() {

	fileName := flag.String("file", "", "file:open the 'Go' source file.")
	flag.Parse()
	if *fileName != "" {
		str := readFile(*fileName)
		for _, v := range codelister.GetFuncs(str) {
			fmt.Println(v)
		}

	} else {
		fmt.Println("usage: codelister -file <filename>")
	}
	

}
