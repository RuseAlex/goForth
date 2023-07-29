package forth

import (
	"fmt"
)

// Remove the number two at the top of the stack,
// and push their sum to the stack
func (e *Eval) add() error { //+
	firstNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	secondNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	e.stack.Push(firstNum + secondNum)
	return nil
}

// Remove the number two at the top of the stack,
// and push their difference to the stack
func (e *Eval) subtract() error { // -
	firstNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	secondNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	e.stack.Push(firstNum - secondNum)
	return nil
}

// Remove the number two at the top of the stack,
// and push their multiple to the stack
func (e *Eval) multiply() error { // *
	firstNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	secondNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	e.stack.Push(firstNum * secondNum)
	return nil
}

// Remove the number two at the top of the stack (s1 & s2)
// find the result of the expression [s2 mod s1],
// pushes the result to the stack.
func (e *Eval) mod() error { // mod
	firstNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	secondNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	e.stack.Push(secondNum % firstNum)
	return nil
}

// Prints a string to the screen
func (e *Eval) dotPrint(tok Token) {
	toEdit := tok.Literal //the string literal is going to be of the form ."<something>"
	if len(toEdit) == 3 {
		fmt.Printf("  ") //the string is empty
	}
	toEdit = toEdit[2 : len(toEdit)-1]
	fmt.Printf(" %s ", toEdit)
}

// Remove the number at the top of the stack,
// and prints it.
func (e *Eval) dot() error { // .
	printNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	fmt.Printf(" %d ", printNum)
	return nil
}

// Remove the number at the top of the stack,
// finds its remainder modulo 128,
// and prints its corresponding ASCII character.
func (e *Eval) emit() error { // emit
	topNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	toPrint := int(topNum) % 128
	fmt.Printf(" %c ", toPrint)
	return nil
}

// Prints a new line
func (e *Eval) cr() error { // cr
	fmt.Println()
	return nil
}

// Rotates the top 3 elements of the stack. The third element
// from the top gets moved to the top and the rest move down
func (e *Eval) rot() error {
	return e.stack.Rot()
}

// Duplicates the second element on the stack and adds it to the stack
func (e *Eval) over() error {
	return e.stack.Over()
}

// Swaps the top 2 elements in the stack
func (e *Eval) swap() error {
	return e.stack.Swap()
}

// Removes the top element of the stack
func (e *Eval) drop() error {
	return e.stack.Drop()
}

// Duplicates the top element and adds it to the stack
func (e *Eval) dup() error {
	return e.stack.Dup()
}

func (e *Eval) lt() error { //+
	firstNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	secondNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	if firstNum > secondNum {
		e.stack.Push(-1)
	} else {
		e.stack.Push(0)
	}
	return nil
}

func (e *Eval) gt() error { //+
	firstNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	secondNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	if firstNum < secondNum {
		e.stack.Push(-1)
	} else {
		e.stack.Push(0)
	}
	return nil
}

func (e *Eval) eq() error {
	firstNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	secondNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	if firstNum == secondNum {
		e.stack.Push(-1)
	} else {
		e.stack.Push(0)
	}
	return nil
}

func (e *Eval) and() error {
	firstNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	secondNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	if (firstNum == -1) && (secondNum == -1) {
		e.stack.Push(-1)
	} else {
		e.stack.Push(0)
	}
	return nil
}

func (e *Eval) or() error {
	firstNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	secondNum, err := e.stack.Pop()
	if err != nil {
		return err
	}
	if (firstNum == -1) || (secondNum == -1) {
		e.stack.Push(-1)
	} else {
		e.stack.Push(0)
	}
	return nil
}

func (e *Eval) invert() error {
	num, err := e.stack.Pop()
	if err != nil {
		return err
	}
	if num == 0 {
		e.stack.Push(-1)
	}
	if num == -1 {
		e.stack.Push(0)
	}
	return nil
}
