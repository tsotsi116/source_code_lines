package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var dir string

	flag.StringVar(&dir, "directory", "", "Directory to walk")

	flag.Parse()

	if len(dir) == 0 {
		fmt.Println("No directory set. Exiting")
		os.Exit(0)
	}

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		skipFolders := []string{".git", "vendor", "node_modules", ".idea", ".vscode", "logs", "venv", "tmp", "target"}

		for _, i := range skipFolders {
			if i == info.Name() {
				return filepath.SkipDir
			}
		}

		if !info.IsDir() {
			lines, err := readSourceCode(path)

			if err != nil {
				fmt.Println("skipped ", path)
			}
			fmt.Println(path + "," + strconv.Itoa(lines))
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func readSourceCode(path string) (int, error) {
	var (
		lineNumbers int
	)
	lineNumbers = 0

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error reading file ", path)
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	text := ""

	for scanner.Scan() {
		text = strings.ReplaceAll(scanner.Text(), " ", "")
		if len(text) > 0 {
			lineNumbers += 1
		}
		text = ""
	}

	return lineNumbers, nil
}
