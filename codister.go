package codister

import (
	"fmt"
	//"log"
)


func main() {
	str := "hrllo kdjfkkk func main(){} ksjkd func hello() string {} skkskks func hi() []int  sdsdjjvyfgfdgdfgfdgu fdgdfgfd func test() int {} ri"
	for _, v := range GetFuncs(str) {
		fmt.Println(v)
	}

}


func GetFuncs(str string) []string {
	var funcNames []string
	var name string
	var returnType string
	for _, v := range FindFuncIndexes(str) {
		/*name = ""
		for string(str[v]) != "(" {
			name += string(str[v])
			//log.Println(name)
			v += 1
		}*/

		name = ExtractFuncName(str, v)
		returnType = FindFuncsReturnType(str, v)
		funcNames = append(funcNames, name + " -> " + returnType)
		//log.Println(funcNames)
	
	}
	
	return funcNames
}

func ExtractFuncName(str string, startIndex int) string {
	var name string
	index := startIndex
	for string(str[index]) != "(" {
		name += string(str[index])
		index++
	}
	return name
}

func FindFuncIndexes(str string) []int { //make it accept a word whatever it is and search for it's index!
	var indexs []int
	/*
		for strings.Contains(str, "func") {
			funcIndex := strings.Index(str, "func")
			indexStr := strconv.Itoa(funcIndex)
			indexs = append(indexs, indexStr)
			str = str[funcIndex:]
		}*/

	for i := 0; i < len(str)-1; i++ {
		if string(str[i]) == "f" && string(str[i+1]) == "u" && string(str[i+2]) == "n" && string(str[i+3]) == "c" { //change that later!
			index := i + len("func")
			indexs = append(indexs, index + 1) // +1 is for the space before the func name.
		}
	}
	return indexs
}

func FindFuncsReturnType(str string, funcIndex int) string { //make it less longer!
	var returnType string
	var extractingIndex int
	index := funcIndex
	for {
		if string(str[index]) != "{" {
			if index != len(str) - 1{
				index++
			}else{
				return "no return" //it doesn't work like that, I should find another way to know if a function has no return
			}
		} else {
			extractingIndex = index
			break
		}

	}
	
	returnTypeR := ""
	for string(str[extractingIndex]) != ")" {
		
		returnTypeR += string(str[extractingIndex])
		extractingIndex--
	//	log.Println(returnTypeR)
		
	}
	//log.Println(returnTypeR)

	for i := len(returnTypeR) - 2 ; i > 1 ; i-- {
		returnType += string(returnTypeR[i])
	//	log.Println(returnType)
	}
	//log.Println(returnType)
	return returnType

}
