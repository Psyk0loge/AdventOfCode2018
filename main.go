package main

import (
	"fmt"
	"os"
	"strings"
)

// taken by go by example
func check(e error) {
	if e != nil {
		panic(e)
	}
}

var stack []string
var input []string
var pointerStack = 0
var pointerInput = 0

func main() {
	//taken from go by example
	dat, err := os.ReadFile("input.txt")
	check(err)
	inputString := string(dat[:])
	input = strings.Split(inputString, "")
	stack = make([]string, len(input))
	fmt.Println(input[pointerInput])
	stack[pointerStack] = input[pointerInput]
	pointerInput++
	for pointerInput := pointerInput; pointerInput < len(input); pointerInput++ {
		morph(input[pointerInput])
	}
	//length = index + 1
	fmt.Println(pointerStack)
	//reset both Pointer and Stack
	pointerInput = 0
	pointerStack = 0
	stack = make([]string, 5000)
	input = make([]string, 5000)
}

func morph(a string) string {
	//here check if it is empty or null ....
	var antiMorph = GetAntiMorph(stack[pointerStack])
	if a == antiMorph {
		stack[pointerStack] = ""
		//Todo: test how you would programm this with side effects
		pointerStack = decreaseByOne(pointerStack)
		pointerInput++
	} else {
		pointerStack++
		stack[pointerStack] = a
		pointerInput++
	}
	return "default"
}

func decreaseByOne(pointer int) int {
	//pointer must not be negative
	if pointer != 0 {
		pointer--
	}
	return pointer
}

func GetAntiMorph(a string) string {
	if strings.ToLower(a) == a {
		return strings.ToUpper(a)
	} else {
		return strings.ToLower(a)
	}
}

//input ist ein seeehr langer String...
//Idee 1: zu Beginn zwei Character einlesen und Morphen
//dann resultat als neunen string für den Morpehn
//dann neuenMorphString mit erstem neuem String von restStrings
//morphen
