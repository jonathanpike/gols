package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printDir(dir string, all bool) error {
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, dir := range info {
		if all {
			fmt.Println(dir.Name())
		} else {
			if []rune(dir.Name())[0] != 46 {
				fmt.Println(dir.Name())
			}
		}
	}
	return nil
}

func main() {
	// Command Line Options
	allBool := flag.Bool("a", false, "do not ignore entries starting with .")

	flag.Parse()

	if len(flag.Args()) > 0 {
		err := printDir(flag.Args()[0], *allBool)
		if err != nil {
			log.Println(err)
		}
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		err = printDir(dir, *allBool)
		if err != nil {
			log.Println(err)
		}
	}
}
