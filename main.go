package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

func parseACF(scanner *bufio.Scanner) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	var currKey string
	var valStart bool
	var val string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		for _, v := range line {
			switch string(v) {
			case `{`:
				var err error
				data[currKey], err = parseACF(scanner)
				if err != nil {
					return nil, err
				}
				currKey = ""
			case `}`:
				return data, nil
			case `"`:
				if !valStart {
					valStart = true
				} else {
					valStart = false
					if currKey == "" {
						currKey = val
					} else {
						data[currKey] = val
						currKey = ""
					}
					val = ""
				}
				continue
			}
			if valStart {
				val += string(v)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func serializeACF(data map[string]interface{}, indent, keyValIndent int) ([]byte, error) {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for _, key := range keys {
		value := data[key]
		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			buf.WriteString(fmt.Sprintf("%v\"%v\"\n%v%v\n", strings.Repeat("\t", indent), key, strings.Repeat("\t", indent), "{"))
			nestedData, err := serializeACF(value.(map[string]interface{}), indent+1, keyValIndent)
			if err != nil {
				return nil, err
			}
			buf.Write(nestedData)
			buf.WriteString(fmt.Sprintf("%v%v\n", strings.Repeat("\t", indent), "}"))
		default:
			buf.WriteString(fmt.Sprintf("%v\"%v\"%v\"%v\"\n", strings.Repeat("\t", indent), key, strings.Repeat("\t", keyValIndent), value))
		}
	}

	return buf.Bytes(), nil
}

func main() {
	// Directory to scan
	dir := "D:/Program Files (x86)/Steam/steamapps"

	// Read the directory contents (does not recurse)
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Iterate over the files in the directory
	for _, file := range files {
		// Skip if it's a directory
		if file.IsDir() {
			continue
		}

		// Check if the file has a .acf extension
		if strings.HasSuffix(file.Name(), ".acf") {
			// Process the .acf file (e.g., print its name)
			fmt.Println("Found .acf file:", filepath.Join(dir, file.Name()))
			err = UpdateACF(filepath.Join(dir, file.Name()))
			if err != nil {
				fmt.Println("Error updating ACF:", err)
			}
			fmt.Println("----------------------------")
		}
	}
}

func UpdateACF(name string) error {
	// 0 always keep this game updated
	// 1 Only update this game when I launch it
	// 2 High Priority - Always auto-update this game before others
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data, err := parseACF(scanner)
	if err != nil {
		return err
	}

	data, err = update(data, []string{"AppState", "AutoUpdateBehavior"}, "1")
	if err != nil {
		return err
	}

	d, err := serializeACF(data, 0, 2)
	if err != nil {
		return err
	}

	err = os.WriteFile(name, d, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func update(data map[string]interface{}, targetKeys []string, value string) (map[string]interface{}, error) {
	if len(targetKeys) == 0 {
		return nil, errors.New("target list is empty")
	}
	targetKey := targetKeys[0]
	if len(targetKeys) == 1 {
		fmt.Printf("Update target value: %v -> %v\n", data[targetKey], value)
		data[targetKey] = value
		return data, nil
	} else {
		_, ok := data[targetKey]
		if !ok {
			return nil, fmt.Errorf("target key %v not found", targetKey)
		}

		targetData, ok := data[targetKey].(map[string]interface{})
		if !ok {
			return nil, errors.New("target is invalid type")
		}

		var err error
		data[targetKey], err = update(targetData, targetKeys[1:], value)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}
