package codelister


func GetFuncs(str string) []string {
	/*This function uses all the other functions, it takes a string then it finds all
	of the functions indexes (FindFuncIndexes) then it extracts their names(ExtractFunName())
	and their return types (ExtractFuncsReturnType())then it adds them into a slice of strings and return it.
	TODO: make every Item in the sclice into a map[string]string as map[function_name]return_type. */
	
	var funcNames []string
	var name string
	var returnType string
	for _, v := range FindFuncIndexes(str) {
	name = ExtractFuncName(str, v)
	returnType = FindFuncsReturnType(str, v)
	funcNames = append(funcNames, name + " -> " + returnType)
	//log.Println(funcNames)
	
	}
	
	return funcNames
}

func ExtractFuncName(str string, startIndex int) string {
	/*This function takes a string and a starting index of the function (from FindFuncIndexes) then it
	will extract the function name untill the first "(" symbol then returns it. */
	var name string
	index := startIndex
	for string(str[index]) != "(" {
		name += string(str[index])
		index++
	}
	return name
}

func FindFuncIndexes(str string) []int { //make it accepts a word whatever it is and search for its index!
	/*
	This function gets a string and return a list of every index that refers to the "func" keyword,
	the index is for the first word after the keyword "func".
	*/
		
	var indexs []int
	for i := 0; i < len(str)-1; i++ {
		if string(str[i]) == "f" && string(str[i+1]) == "u" && string(str[i+2]) == "n" && string(str[i+3]) == "c" && string(str[i+4]) == " " { //change that later!
			index := i + len("func")
			indexs = append(indexs, index + 1) // +1 is for the space before the func name.
		}
	}
	return indexs
}

func FindFuncsReturnType(str string, funcIndex int) string { //make it less longer!
	/*
	This func takes a string and the starting index of the keyword "func" (which we will get from 
	the other function "FindFuncIndexes") and it will extract the type of the function 
	in backward, starting from the symbol "{" until the first ")" then it will reverse it and return it
	TODO: fix the "no return function" bug.
	*/
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
	//extracting the return type in backward
	returnTypeR := ""
	for string(str[extractingIndex]) != ")" {
		
		returnTypeR += string(str[extractingIndex])
		extractingIndex--
	//	log.Println(returnTypeR)
		
	}
	//log.Println(returnTypeR)

	
	//reversing it
	for i := len(returnTypeR) - 2 ; i > 1 ; i-- {
		returnType += string(returnTypeR[i])
	//	log.Println(returnType)
	}
	//log.Println(returnType)
	return returnType

}
