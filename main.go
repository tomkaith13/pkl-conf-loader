package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/apple/pkl-go/pkl"
	"github.com/docker/go-units"
)

type ConfigElem struct {
	Name     string `pkl:"name" json:"name"`
	Value    any    `pkl:"val" json:"val"`
	ElemType string `pkl:"type" json:"type"`
}

type Configs struct {
	// Foo     string       `pkl:"foo" json:"foo"`
	Configs []ConfigElem `pkl:"configs" json:"configs"`
}

type DocConfigs struct {
	Foo        string                        `pkl:"foo" json:"foo"`
	DocConfigs map[string]map[string]Configs `pkl:"doc_confs" json:"doc_confs"`
}

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

	var cfg DocConfigs

	//  ====================================================================================
	// Alert!!
	//  This does not work due to panic when trying to parse ConfigElem, so  resorting to json
	//
	// if err = evaluator.EvaluateModule(
	// 	context.Background(),
	// 	pkl.FileSource("./configuration.pkl"), &cfg); err != nil {
	// 	panic(err)
	// }
	//  ====================================================================================

	err = json.Unmarshal([]byte(textOutput), &cfg)
	if err != nil {
		panic(err)
	}

	fmt.Printf("doc config:  %+v\n", cfg)
	cap_conf, ok := cfg.DocConfigs["test"]["cap1"]
	if !ok {
		panic("Unable to decode")
	}

	for _, conf := range cap_conf.Configs {
		switch conf.ElemType {
		case "csv":
			// do more processing on list
			fmt.Println(conf.Value)
		case "boolean":
			// set some flag
			fmt.Println(conf.Value)
		case "size":
			// Do size based processing
			hsize, err := units.FromHumanSize(conf.Value.(string))
			if err != nil {
				panic(err)
			}
			// Do size based processing
			size, err := units.RAMInBytes(conf.Value.(string))
			if err != nil {
				panic(err)
			}
			fmt.Println("from human size:", hsize)
			fmt.Println("from human size to bytes: ", strconv.FormatInt(size, 10))
			fmt.Println(conf.Value)
		case "duration":
			// Do duration based processing
			// sadly, there is way to transform duration correctly from PKL to stdlib durations
			// things like 10.min and 10d will cause an error.
			// Compare https://pkl-lang.org/package-docs/pkl/0.27.0/base/index.html#DurationUnit
			// to https://cs.opensource.google/go/go/+/refs/tags/go1.23.3:src/time/format.go;l=1601
			duration, err := time.ParseDuration(conf.Value.(string))
			if err != nil {
				panic(err)
			}
			fmt.Println("duration:", duration)

		}
	}
}

func main() {
	ReadConfigs()
}
