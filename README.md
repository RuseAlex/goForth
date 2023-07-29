# goForth
A Go implementation of a Forth like language inspired by this small
ebook https://skilldrick.github.io/easyforth/#conditionals-and-loops

#TODO:
1.Replace ints with float64
2.Refactor code
3.Add If-Else expressions
4.Add Loops
5.Add variables

# How to use it
```go
import f "github.com/RuseAlex/goForth/forth"

f.Forth(`<insert source code here>`)
```

# Examples
Doing basic math operations (. is used to display the resulting number)
```forth
3 7 + .
2 2 * .
4 5 * 2 9 * - .
```
Calculating the polynomial (x^2 + 2x + 1 - x^3) of a number few numbers
```forth
    :poly dup 1 + dup * swap dup dup * * swap - .;
	1 poly
	2 poly
```
And another one 
```forth
	1 2 3 4 5 6 7 8 9 10 dup dup . . . cr
	rot . . . cr
	over 10 * 15 + emit cr
	drop dup . rot dup . rot dup . rot cr
	mod dup . dup dup = . crs
```forth