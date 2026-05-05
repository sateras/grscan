package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	root := flag.String("path", ".", "root directory to scan")
	query := flag.String("q", "person-companies", "remote name or substring to search")
	flag.Parse()

	err := filepath.WalkDir(*root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		// ищем .git папки
		if d.IsDir() && d.Name() == ".git" {
			repoPath := filepath.Dir(path)

			out, err := exec.Command("git", "-C", repoPath, "remote", "-v").Output()
			if err != nil {
				return nil
			}

			matched := false
			scanner := bufio.NewScanner(strings.NewReader(string(out)))

			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, *query) {
					matched = true
				}
			}

			if matched {
				fmt.Println("FOUND:", repoPath)
				fmt.Println(string(out))
				fmt.Println("-----------")
			}

			// не заходим глубже внутрь .git
			return filepath.SkipDir
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}
