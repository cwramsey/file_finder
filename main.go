package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filename := flag.String("file", "", "Name (or part of) the file you want to find. ex: main.go | main")
	directory := flag.String("dir", "", "Directory you'd like to search in. ex: /var/www/")
	recursive := flag.Bool("recurse", false, "Searches directories recursively")
	flag.Parse()

	if *filename == "" {
		fmt.Println("Option[file] must be set.")
		os.Exit(1)
	}

	if *directory == "" {
		*directory = "./"
	}

	if *recursive {
		dir := []string{*directory}
		searchRecursively(*filename, dir)
	} else {
		search(*filename, *directory)
	}

}

func search(filename string, directory string) []string {
	os.Chdir(directory)

	directoryItems, err := ioutil.ReadDir(directory)

	if err == nil {
		var directories []string
		for i := range directoryItems {
			thisItem := directoryItems[i]

			if thisItem.IsDir() == false && strings.Contains(thisItem.Name(), filename) {
				fmt.Printf("Match: %v\n", directory+"/"+thisItem.Name())
			} else if thisItem.IsDir() {
				directories = append(directories, directory+"/"+thisItem.Name())
			}
		}

		return directories
	}

	return nil
}

func searchRecursively(filename string, nextDirectories []string) []string {

	if nextDirectories != nil {
		var nextLevelDirectories []string
		if nextDirectories != nil {
			for i := range nextDirectories {
				nextLevelDirectories = append(nextLevelDirectories, search(filename, nextDirectories[i])...)
			}
		}

		if len(nextDirectories) > 0 {
			searchRecursively(filename, nextLevelDirectories)
		}
	}

	return nil
}
