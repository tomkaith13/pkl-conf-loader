package main

import (
	"context"
	"fmt"

	"github.com/apple/pkl-go/pkl"
)

type ConfigElem struct {
	Name     string `pkl:"name"`
	Value    string `pkl:"val"`
	ElemType string `pkl:"type"`
}
type Configs struct {
	Foo string `pkl:"foo"`
	// Configs []ConfigElem `pkl:"configs"`
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
	if err = evaluator.EvaluateModule(
		context.Background(),
		pkl.FileSource("./configuration.pkl"), &cfg); err != nil {
		panic(err)
	}

	fmt.Printf("doc config:  %+v", cfg)
}

func main() {
	ReadConfigs()
}
