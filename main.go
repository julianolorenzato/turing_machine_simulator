package main

import (
	"fmt"
	"tumasim/tm"
)

func main() {
	turingMachine := tm.New()

	prog := []map[tm.Symbol]tm.Transition{
		{},
		{
			'#': {WriteSymbol: '.', Direction: true, NewState: 0},
			'1': {WriteSymbol: '.', Direction: true, NewState: 2},
		},
		{
			'1': {WriteSymbol: '1', Direction: true, NewState: 2},
			'#': {WriteSymbol: '#', Direction: true, NewState: 2},
			'.': {WriteSymbol: '1', Direction: false, NewState: 3},
		},
		{
			'#': {WriteSymbol: '#', Direction: false, NewState: 3},
			'1': {WriteSymbol: '1', Direction: false, NewState: 3},
			'.': {WriteSymbol: '.', Direction: true, NewState: 1},
		},
	}

	input := tm.Tape{'.', '1', '1', '1', '#', '1', '1', '.'}

	output := turingMachine.Execute(input, prog)

	fmt.Println("output:", string(output))
}
