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
	allBool    bool
	longOutput bool
	output     io.Writer
}

func returnFiles(root string, all bool) ([]os.FileInfo, error) {
	var files []os.FileInfo
	info, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, file := range info {
		if all {
			files = append(files, file)
		} else {
			if []rune(file.Name())[0] != 46 {
				files = append(files, file)
			}
		}
	}
	return files, nil
}

func printResults(files []os.FileInfo, long bool) error {
	if len(files) == 0 {
		return errors.New("The list of files is empty")
	}
	for i, file := range files {
		if long {
			fmt.Fprintf(Config.output, "%v %v %v %v\n", file.Mode(), file.Size(), file.ModTime().Format("Jan _2 15:04 2006"), file.Name())
		} else {
			if i == len(files)-1 {
				fmt.Fprintf(Config.output, "%v\n", file.Name())
			} else {
				fmt.Fprintf(Config.output, "%v ", file.Name())
			}
		}
	}
	return nil
}

func init() {
	// Command Line Options
	flag.BoolVar(&Config.allBool, "a", false, "do not ignore entries starting with .")
	flag.BoolVar(&Config.longOutput, "l", false, "use a long listing format")
	flag.Parse()
	Config.output = os.Stdout
	log.SetOutput(os.Stderr)
}

func main() {
	// Grab file names from either specified directory or
	// the current directory (if no directory is specified)
	var files []os.FileInfo
	var err error
	if len(flag.Args()) > 0 {
		files, err = returnFiles(flag.Args()[0], Config.allBool)
		if err != nil {
			log.Println(err)
		}
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		files, err = returnFiles(dir, Config.allBool)
		if err != nil {
			log.Println(err)
		}
	}
	// Print results to Config.output
	err = printResults(files, Config.longOutput)
	if err != nil {
		log.Println(err)
	}
}
