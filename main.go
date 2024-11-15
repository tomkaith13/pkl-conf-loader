package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apple/pkl-go/pkl"
)

type ConfigElem struct {
	Name     string `pkl:"name" json:"name"`
	Value    any    `pkl:"val" json:"val"`
	ElemType string `pkl:"type" json:"type"`
}
type Configs struct {
	Foo     string       `pkl:"foo" json:"foo"`
	Configs []ConfigElem `pkl:"configs" json:"configs"`
}

// type DocConfigs struct {
// 	DocConfigs map[string]map[string]Configs `pkl:"doc_confs"`
// }

func ReadConfigs() {
	evaluator, err := pkl.NewEvaluator(context.Background(), pkl.PreconfiguredOptions)
	if err != nil {
		panic(err)
	}
	defer evaluator.Close()

	textOutput, err := evaluator.EvaluateOutputText(
		context.Background(),
		pkl.FileSource("./configuration.pkl"))
	if err != nil {
		panic(err)
	}
	fmt.Println(textOutput)

	var cfg Configs

	// Alert!!
	//  This does not work due to panic when trying to parse ConfigElem, so  resorting to json
	//
	// if err = evaluator.EvaluateModule(
	// 	context.Background(),
	// 	pkl.FileSource("./configuration.pkl"), &cfg); err != nil {
	// 	panic(err)
	// }
	err = json.Unmarshal([]byte(textOutput), &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("doc config:  %+v", cfg)
}

func main() {
	ReadConfigs()
}
