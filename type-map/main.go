package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
)

func main() {
	m := map[string]string{
		"name": "Weerasak",
		"club": "Liverpool",
	}

	m["age"] = "33"

	js, _ := json.Marshal(m)
	fmt.Printf("%s\n", js)

	jst := `
{
		"age":"33",
		"club":"Liverpool",
		"name":"Weerasak"
}
`
	var m2 map[string]string
	if err := json.Unmarshal([]byte(jst), &m2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m2)
}

func PrintSortMap(m map[string]string) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
