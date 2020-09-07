package main

import (
	"bufio"
	"fmt"
	"os"		
)

type word []string

func (w *word) getInput() string {
	fmt.Println("Game started!")
	fmt.Println("Type in the word to guess!")
	text := bufio.NewScanner(os.Stdin)
	text.Scan()
	cnt := len(text.Text())
	clearTerminal()
	w.initDashes(cnt)	
	return text.Text()		
}

func (w *word) initDashes(cnt int) {			
	for i := 0; i < cnt; i++ {
		*w = append (*w, "_")
	}
}
