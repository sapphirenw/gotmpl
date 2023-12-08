package main

import (
	"fmt"
	"html/template"
	"strings"
)

func defaultFuncMap() *template.FuncMap {
	return &template.FuncMap{
		"dict":     dict,
		"truncstr": truncstr,
		"substr":   substr,
		"join":     join,
	}
}

// allows for arbitrary values to be passed to the template with custom names
func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

// truncates a string `str` to `num` characters
func truncstr(str string, num int) string {
	if len(str) > num {
		return str[:num] + "..."
	}
	return str
}

func substr(str string, num int) string {
	if len(str) > num {
		return str[:num]
	}
	return str
}

func join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}
