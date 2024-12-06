package steam_acf

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

const (
	serializeRecursiveIndent       = 0
	serializeRecursiveKeyValIndent = 2
)

var (
	writeRegex = regexp.MustCompile(`^[^\n{}"]*$`)
)

type SteamACF interface {
	Get(target []string) (string, error)
	Update(target []string, value string) (string, error)
	Serialize() []byte
	String() string
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

func (sa *steamACF) Get(target []string) (string, error) {
	return "", errors.New("Not implemented")
}

func (sa *steamACF) Update(target []string, value string) (string, error) {
	if !writeRegex.MatchString(value) {
		return "", fmt.Errorf("invalid value: value match with regex %v", writeRegex.String())
	}
	previousValue, updatedData, err := updateRecursive(sa.data, target, value, nil)
	if err != nil {
		return "", err
	}
	sa.data = updatedData
	return fmt.Sprint(previousValue), nil
}

func (sa *steamACF) Serialize() []byte {
	return serializeRecursive(sa.data, serializeRecursiveIndent, serializeRecursiveKeyValIndent)
}

func (sa *steamACF) String() string {
	return string(sa.Serialize())
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

// TODO finish
//func getRecursive(data map[string]interface{}, tk []string, ptk []string) ([]byte, error) {
//	if len(tk) == 0 {
//		return serializeRecursive(data, serializeRecursiveIndent, serializeRecursiveKeyValIndent), nil
//	}
//	targetKey := tk[0]
//	ptk = append(ptk, targetKey)
//	dataValue, hasChild := data[targetKey]
//	if !hasChild {
//		return nil, fmt.Errorf("target key '%v' is not found", strings.Join(ptk, "."))
//	}
//	_, isMap := dataValue.(map[string]interface{})
//	if len(tk) != 1 {
//		if !isMap {
//			return nil, fmt.Errorf("target key '%v' is not found", strings.Join(append(ptk, tk[1]), "."))
//		}
//		var err error
//		dataValue, err = getRecursive(dataValue.(map[string]interface{}), tk[1:], ptk)
//		if err != nil {
//			return nil, err
//		}
//		return dataValue, nil
//	}
//	if isMap {
//		return serializeRecursive(dataValue.(map[string]interface{}), serializeRecursiveIndent, serializeRecursiveKeyValIndent), nil
//	}
//	return []byte(fmt.Sprintf(dataValue)), nil
//}

func updateRecursive(data map[string]interface{}, tk []string, value string, ptk []string) (interface{}, map[string]interface{}, error) {
	if len(tk) == 0 {
		return nil, nil, errors.New("target keys is empty")
	}
	targetKey := tk[0]
	ptk = append(ptk, targetKey)
	dataValue, hasChild := data[targetKey]
	if !hasChild {
		return nil, nil, fmt.Errorf("target key '%v' is not found", strings.Join(ptk, "."))
	}
	_, isMap := dataValue.(map[string]interface{})
	if len(tk) != 1 {
		if !isMap {
			return nil, nil, fmt.Errorf("target key '%v' is not found", strings.Join(append(ptk, tk[1]), "."))
		}
		var err error
		dataValue, data[targetKey], err = updateRecursive(dataValue.(map[string]interface{}), tk[1:], value, ptk)
		if err != nil {
			return nil, nil, err
		}
		return dataValue, data, nil
	}
	if isMap {
		return nil, nil, fmt.Errorf("target key '%v' is not a value", strings.Join(ptk, "."))
	}
	data[targetKey] = value
	return dataValue, data, nil
}

func serializeRecursive(data map[string]interface{}, indent, keyValIndent int) []byte {
	var (
		keys            = make([]string, len(data))
		indentStr       = strings.Repeat("\t", indent)
		keyValIndentStr = strings.Repeat("\t", keyValIndent)
	)
	var i int
	for key := range data {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, key := range keys {
		switch value := data[key].(type) {
		case map[string]interface{}:
			buf.WriteString(fmt.Sprintf("%v\"%v\"\n%v{\n", indentStr, key, indentStr))
			buf.Write(serializeRecursive(value, indent+1, keyValIndent))
			buf.WriteString(fmt.Sprintf("%v}\n", indentStr))
		default:
			buf.WriteString(fmt.Sprintf("%v\"%v\"%v\"%v\"\n", indentStr, key, keyValIndentStr, value))
		}
	}
	return buf.Bytes()
}
