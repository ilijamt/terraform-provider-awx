package internal

import (
	"log"
)

type fnResourceProcessor func(map[string]any) (map[string]any, error)

var fnResourceProcessors = map[string]fnResourceProcessor{
	"Config": seedReadAction,
	"Ping":   seedReadAction,
}

func seedReadAction(in map[string]any) (map[string]any, error) {
	if _, ok := in["actions"]; !ok {
		in["actions"] = map[string]any{"GET": map[string]any{}}
	}
	return in, nil
}

func ResourceProcessor(name string, in map[string]any) (map[string]any, error) {
	in = NormalizeResourcePayload(in)
	if fn, ok := fnResourceProcessors[name]; ok {
		log.Printf("Executing custom processor for %s", name)
		return fn(in)
	}
	return in, nil
}
