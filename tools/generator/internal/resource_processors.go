package internal

import (
	"log"
)

type fnResourceProcessor func(map[string]any) (map[string]any, error)

var fnResourceProcessors = make(map[string]fnResourceProcessor)

func ResourceProcessor(name string, in map[string]any) (map[string]any, error) {
	if fn, ok := fnResourceProcessors[name]; ok {
		log.Printf("Executing custom processor for %s", name)
		return fn(in)
	}
	return in, nil
}
