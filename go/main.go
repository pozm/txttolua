package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	if 1 == len(os.Args) { // lol lazy
		fmt.Println("I Require an argument to be inputted.")
		return
	}

	pathTo := os.Args[1]

	abs, err := filepath.Abs(pathTo)
	ext := path.Ext(abs)
	if ext != ".txt" {
		fmt.Println("I can only convert txt files into lua files.")
		return
	}

	if err != nil {
		fmt.Println("Path does not exist")
		return
	}
	err = os.Rename(abs, abs[0:len(abs)-len(ext)]+".lua") // this is so scuffed ik, sorry cy if you view this.
	if err != nil {
		fmt.Println("unable to change")
		return
	}
}
