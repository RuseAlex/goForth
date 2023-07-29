package forth

func Forth(input string) {
	stack := Stack{}
	evl := NewEval(input, &stack)
	evl.Exec()
}
