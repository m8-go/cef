package cef

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

func mkCEFVersion(v []byte) (int, error) {
	ss := bytes.Split(v, []byte(":"))

	if len(ss) != 2 {
		return 0, errors.New("")
	}

	ver, err := strconv.Atoi(string(ss[1]))
	if err != nil {
		return 0, err
	}

	return ver, nil
}

func mkCollection(b []byte) []string {
	var collection []string

	str := string(b)
	ss := strings.Split(str, " ")

	for _, s := range ss {
		if !strings.ContainsRune(s, '=') {
			collection[len(collection)-1] = collection[len(collection)-1] + s

			continue
		}

		collection = append(collection, s)
	}

	return collection
}
