package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printDir(dir string, hidden bool) error {
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, dir := range info {
		if []rune(dir.Name())[0] != 46 {
			fmt.Println(dir.Name())
		}
	}
	return nil
}

func main() {
	if len(os.Args) > 1 {
		err := printDir(os.Args[1], false)
		if err != nil {
			log.Println(err)
		}
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		err = printDir(dir, false)
		if err != nil {
			log.Println(err)
		}
	}
}
