package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var Config struct {
	allBool bool
	output  io.Writer
}

func printDir(root string, all bool) error {
	info, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}
	for _, dir := range info {
		if all {
			fmt.Fprintln(Config.output, dir.Name())
		} else {
			if []rune(dir.Name())[0] != 46 {
				fmt.Fprintln(Config.output, dir.Name())
			}
		}
	}
	return nil
}

func init() {
	// Command Line Options
	flag.BoolVar(&Config.allBool, "a", false, "do not ignore entries starting with .")
	flag.Parse()
	Config.output = os.Stdout
	log.SetOutput(os.Stderr)
}

func main() {
	if len(flag.Args()) > 0 {
		err := printDir(flag.Args()[0], Config.allBool)
		if err != nil {
			log.Println(err)
		}
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		err = printDir(dir, Config.allBool)
		if err != nil {
			log.Println(err)
		}
	}
}
