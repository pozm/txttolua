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
	ext := path.Ext(pathTo)
	if ext != ".txt" {
		fmt.Println("I can only convert txt files into lua files.")
		return
	}

	abs, err := filepath.Abs(pathTo)
	if err != nil {
		fmt.Println("Path does not exist")
		return
	}
	err = os.Rename(abs, filepath.Dir(abs)+"\\"+(filepath.Base(abs)[0:len(ext)])+".lua") // this is so scuffed ik, sorry cy if you view this.
	if err != nil {
		fmt.Println("unable to change")
		return
	}
}
