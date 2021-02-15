package main

import	{
	"testing"
	"fmt"
}

func	TestInput(t* testing.T) {

	emptyResult := main("")

	if emptyResult != "Error: wrong number of arguments" {
		t.Error("main(\"\") failed, expected %v, got %v", "Error: wrong number of arguments", emptyResult)
	}
}