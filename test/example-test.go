package main

import	{
	"testing"
	"fmt"
	"github.com/rakyll/gotest"
}

func	TestInput(t* testing.T) {

	emptyResult := main("")

	if emptyResult != "Error: wrong number of arguments" {
		t.Errorf("main(\"\") failed, expected %v, got %v", "Error: wrong number of arguments", emptyResult)
	}
	else {
		t.Logf("the program started properly")
	}
}