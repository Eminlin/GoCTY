package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	//取出所有key 方便后续进行匹配
	c := make(map[string]json.RawMessage)
	b, err := ioutil.ReadFile("testFile/cty.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	e := json.Unmarshal(b, &c)
	if e != nil {
		fmt.Println(err)
		return
	}
	k := make([]string, len(c))
	i := 0
	for s, _ := range c {
		k[i] = s
		i++
	}
	// fmt.Printf("%+v\n", k)
	// fmt.Println(len(k))
}
