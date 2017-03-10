package main

import (
	"errors"
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

func returnFileNames(root string, all bool) ([]string, error) {
	var files []string
	info, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, dir := range info {
		if all {
			files = append(files, dir.Name())
		} else {
			if []rune(dir.Name())[0] != 46 {
				files = append(files, dir.Name())
			}
		}
	}
	return files, nil
}

func printResults(files []string) error {
	if len(files) == 0 {
		return errors.New("The list of files is empty")
	}
	for i, file := range files {
		if i == len(files)-1 {
			fmt.Fprintf(Config.output, "%v\n", file)
		} else {
			fmt.Fprintf(Config.output, "%v ", file)
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
	var files []string
	var err error
	if len(flag.Args()) > 0 {
		files, err = returnFileNames(flag.Args()[0], Config.allBool)
		if err != nil {
			log.Println(err)
		}
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		files, err = returnFileNames(dir, Config.allBool)
		if err != nil {
			log.Println(err)
		}
	}
	err = printResults(files)
	if err != nil {
		log.Println(err)
	}
}
