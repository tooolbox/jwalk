package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tooolbox/jwalk"
)

const input = `{
	"key1": [{
		"id": 1
	}, {
		"id": 2
	}],
	"key2": [1,2],
	"key3": {
		"id": 1,
		"name": null
	}
}`

func main() {
	i, err := jwalk.Parse([]byte(input))
	if err != nil {
		log.Fatal(err)
	}

	switch v := i.(type) {
	case jwalk.ObjectWalker:
		v.Walk(func(key string, value interface{}) (interface{}, error) {
			fmt.Println(key + ":")
			switch v := value.(type) {
			case jwalk.ObjectsWalker:
				v.Walk(func(obj jwalk.ObjectWalker) error {
					fmt.Println("\t-")
					obj.Walk(func(key string, value interface{}) (interface{}, error) {
						if v, ok := value.(jwalk.Value); ok {
							fmt.Println("\t", key+":", v.Interface())
							if v.Int() == 2 {
								fmt.Println("Found 2")
								return jwalk.JSONMarshaler(3), nil
							}
						}
						return value, nil
					})
					return nil
				})
			case jwalk.Value:
				fmt.Println("\t", v.Interface())
			case jwalk.ObjectWalker:
				v.Walk(func(key string, value interface{}) (interface{}, error) {
					if v, ok := value.(jwalk.Value); ok {
						fmt.Println("\t", key+":", v.Interface())
					}
					return value, nil
				})
			}
			return value, nil
		})
	}

	b, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(b))
}
