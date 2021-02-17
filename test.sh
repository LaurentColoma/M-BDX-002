#!/bin/bash

echo "Testing parsing"
go test parsing/parsing_test.go parsing/parsing.go

echo "Testing pathFinding"
go test pathFinding/pathFinding_test.go pathFinding/pathFinding.go