package main

import (
	"testing"
	"reflect"
)


func TestFindFuncIndexes(t *testing.T){
	str := "package hello func hello() string {}"
	if funcIndexes := FindFuncIndexes(str); funcIndexes[0] != 19 {
		t.Errorf("Expected 20 got %d ", funcIndexes)
	}
}

func TestExtractFuncName(t *testing.T){
	str := "package hello func hello() string {}"
	if funcName := ExtractFuncName(str, 19); funcName != "hello" {
		t.Errorf("Expected 'hello', got '%s'", funcName)
	}
}

func TestFuncsReturnType(t *testing.T){
	str := "package hello func hello() string {}"
	if returnType := FindFuncsReturnType(str, 19); returnType != "string" {
		t.Errorf("Expected 'string', got '%s'", returnType)
	}

}

func TestGetFuncs(t *testing.T){
	str := "package hello func hello() string {} other stuff here func hi() int {}  other stuff"
	if funcNames := GetFuncs(str); !reflect.DeepEqual(funcNames, []string{"hello -> string", "hi -> int"}) {
		t.Errorf("Expected '[hello -> string, hi -> int]', got '%v'", funcNames)
	}
}


