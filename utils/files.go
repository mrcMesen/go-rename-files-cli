package utils

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

func isCamelCase(str string) bool {
	for _, letter := range str {
		if unicode.IsUpper(letter) {
			return true
		}
	}
	return false
}

// check if the path include part of any item blacklist
func isInBlackList(str string) bool {
	blackList := []string{"/node_modules", "/lib/", "/bin/", "/dist/", "/build/", "DS_Store", "README"}
	for _, item := range blackList {
		if strings.Contains(str, item) {
			return true
		}
	}

	return false
}

func cleanPath(path string) (dir string, file string) {
	return filepath.Split(path)
}

func GetSnakeCase(path string) string {
	dir, cleanPath := cleanPath(path)
	ext := filepath.Ext(cleanPath)
	cleanPath = cleanPath[:len(cleanPath)-len(ext)]
	snakeCase := ""
	for i, letter := range cleanPath {
		// validate if it is a letter
		if !unicode.IsLetter(letter) {
			continue
		}
		// if is upper
		if unicode.IsUpper(letter) {
			// if it is not the first or last letter
			if len(snakeCase) > 0 && i != len(cleanPath)-1 {
				snakeCase += "-" + string(unicode.ToLower(letter))
			} else {
				snakeCase += string(unicode.ToLower(letter))
			}
			continue
		}
		snakeCase += string(letter)
	}
	fullNewPath := dir + snakeCase + ext
	return fullNewPath
}

/*
This function print and return the files to be renamed,
searching in the received path and subdirectories.
*/
func ListCamelCasePaths(path string) []string {
	files := []string{}
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if isInBlackList(path) {
				return nil
			}
			if isCamelCase(info.Name()) {
				files = append(files, path)
			}
			return nil
		})
	if err != nil {
		return []string{}
	}
	// sort the list of files by the length of the path
	// so that the files in the deepest directories are renamed first
	// this is to avoid renaming a file that has already been renamed
	sort.Slice(files, func(i, j int) bool {
		return len(files[i]) > len(files[j])
	})

	return files
}

func ListNewNames(files []string) []string {
	newNames := []string{}
	for _, file := range files {
		newNames = append(newNames, GetSnakeCase(file))
	}
	return newNames
}

func RenameFile(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}
