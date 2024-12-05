package steam_acf

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"
)

type SteamACF interface {
	Update(target []string, value string) error
	Serialize() ([]byte, error)
}

type steamACF struct {
	data map[string]interface{}
}

func Parse(data []byte) (SteamACF, error) {
	parsedSteamACFData, err := parseRecursive(bufio.NewScanner(bytes.NewReader(data)))
	if err != nil {
		return nil, err
	}
	return &steamACF{data: parsedSteamACFData}, nil
}

func (sa *steamACF) Update(target []string, value string) error {
	updatedData, err := updateRecursive(sa.data, target, value)
	if err != nil {
		return err
	}
	sa.data = updatedData
	return nil
}

func (sa *steamACF) Serialize() ([]byte, error) {
	return serializeRecursive(sa.data, 0, 2)
}

func parseRecursive(scanner *bufio.Scanner) (map[string]interface{}, error) {
	var (
		data             = make(map[string]interface{})
		currKey, currVal string
		valStart         bool
		err              error
	)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		for _, v := range line {
			switch string(v) {
			case `{`:
				data[currKey], err = parseRecursive(scanner)
				if err != nil {
					return nil, err
				}
				currKey = ""
			case `}`:
				return data, nil
			case `"`:
				if !valStart {
					valStart = true
					continue
				}
				if currKey == "" {
					currKey = currVal
				} else {
					data[currKey] = currVal
					currKey = ""
				}
				currVal = ""
				valStart = false
			default:
				if valStart {
					currVal += string(v)
				}
			}
		}
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func updateRecursive(data map[string]interface{}, targetKeys []string, value string) (map[string]interface{}, error) {
	if len(targetKeys) == 0 {
		return nil, errors.New("target list is empty")
	}
	targetKey := targetKeys[0]
	if len(targetKeys) == 1 {
		data[targetKey] = value
		return data, nil
	}
	dataValue, ok := data[targetKey]
	if !ok {
		return nil, fmt.Errorf("target key %v not found", targetKey)
	}
	targetDataValue, ok := dataValue.(map[string]interface{})
	if !ok {
		return nil, errors.New("target is invalid type")
	}
	var err error
	data[targetKey], err = updateRecursive(targetDataValue, targetKeys[1:], value)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func serializeRecursive(data map[string]interface{}, indent, keyValIndent int) ([]byte, error) {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, key := range keys {
		switch value := data[key].(type) {
		case map[string]interface{}:
			buf.WriteString(fmt.Sprintf("%v\"%v\"\n%v%v\n", strings.Repeat("\t", indent), key, strings.Repeat("\t", indent), "{"))
			nestedData, err := serializeRecursive(value, indent+1, keyValIndent)
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
