package fileLib

import (
	"fmt"
	"io/ioutil"
)

var _VERSION string = "0.0.1"

func main() {
	fmt.Println(_VERSION)
}

func ReadFileIntoSlice(file string) ([]string, error) {
	var data, err = ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var strings = make([]string, 0)
	var curr = ""
	for _,item := range data{
		if string(item) != "\n"{
			curr = curr + string(item)
		} else {
			strings = append(strings, curr)
			curr = ""
		}
	}
	if len(curr) > 0 {
		strings = append(strings, curr)
	}

	return strings, err
}