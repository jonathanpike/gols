package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printDir(dir string) error {
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, dir := range info {
		fmt.Println(dir.Name())
	}
	return nil
}

func main() {
	if len(os.Args) > 1 {
		err := printDir(os.Args[1])
		if err != nil {
			log.Println(err)
		}
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		err = printDir(dir)
		if err != nil {
			log.Println(err)
		}
	}
}
