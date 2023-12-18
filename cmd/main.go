package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"test-writer/internal/directory"
	"test-writer/internal/fileviewer"
	"test-writer/internal/gpt4"
)

const cyanBold = "\033[1;36m" // Bold Cyan
const reset = "\033[0m"

func main() {
	folderPath := flag.String("path", "", "folder path to list directories and TypeScript files")
	flag.Parse()

	if *folderPath == "" {
		log.Fatal("--path param required")
	}

	navigate(*folderPath)
}

func navigate(currentPath string) {
	for {
		dirs, err := directory.List(currentPath)
		if err != nil {
			log.Fatalf("Error listing contents: %v", err)
		}

		tsFiles, err := directory.FindTypeScriptFiles(currentPath)
		if err != nil {
			log.Fatalf("Error finding TypeScript files: %v", err)
		}
		fmt.Printf("\nViewing %s:\n", currentPath)
		for i, dir := range dirs {
			fmt.Printf("%d: /%s\n", i+1, dir) // Prefixed with '/'
		}

		for i, file := range tsFiles {
			fmt.Printf("%d: %s%s%s (TypeScript file)\n", i+len(dirs)+1, cyanBold, file, reset) // Bold Cyan for TypeScript files
		}

		fmt.Print("\nEnter number to navigate or 'q' to quit: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "q" {
			break
		}

		choice, err := parseInput(input, len(dirs)+len(tsFiles))
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		if choice <= len(dirs) {
			currentPath = filepath.Join(currentPath, dirs[choice-1])
		} else {
			selectedFile := tsFiles[choice-len(dirs)-1]
			contents, err := fileviewer.DisplayFileContents(filepath.Join(currentPath, selectedFile))
			if err != nil {
				log.Fatalf("Error displaying file contents: %v", err)
			}
			gpt4.Generate(selectedFile, contents, "mock")
			//os.Exit(0) // Exit the program after displaying the file
		}
	}
}

func parseInput(input string, max int) (int, error) {
	if choice, err := strconv.Atoi(input); err == nil {
		if choice >= 1 && choice <= max {
			return choice, nil
		}
	}
	return 0, fmt.Errorf("invalid input")
}
