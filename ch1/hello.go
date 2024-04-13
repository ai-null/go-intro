package main

import "fmt"

// why can't go add curly braces in another line?
// because the lexer (lexical analysis or tokenization)
// programmed to add ; semicolon after ) if there's a line break

// which turned into :
// func
// main
// (
// )
// \n << line break
// ; << the error
// {

// func main ();
// {}

func main() {
	fmt.Printf("Hello, %s!\n", "World")
}
