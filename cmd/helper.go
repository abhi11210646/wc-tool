package main

import "fmt"

func outputHelp() {
	fmt.Println(`
	USAGE:  wc [OPTIONS] FILE(S)
		-m      print the character counts
		-l      print the newline counts
		-w      print the word counts  
	EXAMPLE: 
		wc -l logs.txt
		wc -w logs.txt
		wc -l logs.txt logs2.txt
	  `)
}
