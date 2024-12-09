package tm

import "fmt"

type Symbol byte

type Tape []Symbol

type TuringMachine struct {
	tape     Tape
	state    State
	position uint32
}

type Transition struct {
	WriteSymbol Symbol
	Direction   bool
	NewState    uint32
}

type State uint32

type Program []map[Symbol]Transition

func New() *TuringMachine {
	return &TuringMachine{
		tape:     make(Tape, 32), // limiting tape to 32 positions (theoretically it can be unlimited)
		state:    1,              // stars at the state 1, state 0 is the final state by convention
		position: 1,
	}
}

func (tm *TuringMachine) Execute(input Tape, prog Program) Tape {
	// fill the tm tape
	for i := range tm.tape {
		if i >= len(input) {
			tm.tape[i] = '.'
		} else {
			tm.tape[i] = input[i]
		}
	}

	for {
		fmt.Println("(state):", tm.state, "(head position):", tm.position, "(tape):", string(tm.tape))
		if tm.state == 0 {
			fmt.Println("Final state reached, stopping turing machine")
		}

		readSymbol := tm.read()                      // read from tape
		transition, ok := prog[tm.state][readSymbol] // look for valid transition
		if !ok {
			fmt.Println("No valid transitions, stopping turing machine")
			break
		} else {
			tm.write(transition.WriteSymbol) // write symbol on the tape

			if transition.Direction { // go to left or right, depending of transition
				tm.right()
			} else {
				tm.left()
			}

			tm.state = State(transition.NewState) // change state
		}
	}

	return tm.tape
}

func (tm *TuringMachine) read() Symbol {
	return tm.tape[tm.position]
}

func (tm *TuringMachine) write(sym Symbol) {
	tm.tape[tm.position] = sym
}

func (tm *TuringMachine) left() {
	tm.position--
}

func (tm *TuringMachine) right() {
	tm.position++
}
