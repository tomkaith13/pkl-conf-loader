package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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

// v2
type ConfigElemV2 struct {
	Value    any    `json:"val"`
	ElemType string `json:"type"`
}

type DocConfigsV2 struct {
	DocConfigsV2 map[string]map[string]map[string]ConfigElemV2 `json:"doc_confs"`
}

// v3
type ConfigElemV3 struct {
	Value any `json:"val"`
}

type DocConfigsV3 struct {
	DocConfigsV3 map[string]map[string]map[string]ConfigElemV3 `json:"doc_confs"`
}

// v4
type DocConfigsV4 struct {
	DocConfigsV4 map[string]map[string]map[string]any `json:"doc_confs"`
}

const ConfigurationFile string = "./configuration.pkl"
const ConfigurationFileV2 string = "./configurationv2.pkl"
const ConfigurationFileV3 string = "./configurationv3.pkl"
const ConfigurationFileV4 string = "./configurationv4.pkl"

func ReadConfigs(isFile bool, textPkl string) {
	evaluator, err := pkl.NewEvaluator(context.Background(), pkl.PreconfiguredOptions)
	if err != nil {
		panic(err)
	}
	defer evaluator.Close()

	var textOutput string
	if isFile {
		textOutput, err = evaluator.EvaluateOutputText(
			context.Background(),
			pkl.FileSource(ConfigurationFile))
		if err != nil {
			panic(err)
		}
		fmt.Println(textOutput)
	} else {
		textOutput, err = evaluator.EvaluateOutputText(
			context.Background(),
			pkl.TextSource(textPkl))
		if err != nil {
			panic(err)
		}
		fmt.Println(textOutput)
	}

	var cfg DocConfigs

	//  ====================================================================================
	// Alert!!
	// This does not work due to panic when trying to parse ConfigElem, so  resorting to json
	// panic: cannot decode Pkl value of type `base#ConfigElem` into Go type `interface {}`.
	// Define a custom mapping for this using `pkl.RegisterMapping`
	//
	// var pklConf DocConfigs
	// if err = evaluator.EvaluateModule(
	// 	context.Background(),
	// 	pkl.FileSource("./configuration.pkl"), &pklConf); err != nil {
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
			csvStr := conf.Value.(string)
			csvList := strings.Split(csvStr, ",")
			fmt.Println(csvList)
		case "boolean":
			// set some flag
			fmt.Println(conf.Value)
		case "size":
			sizeStr := conf.Value.(string)
			if strings.Contains(sizeStr, "ib") {

				// Do size based processing
				size, err := units.RAMInBytes(sizeStr)
				if err != nil {
					panic(err)
				}
				fmt.Println("from human size to bytes: ", strconv.FormatInt(size, 10))
				continue
			}

			// Do size based processing
			hsize, err := units.FromHumanSize(conf.Value.(string))
			if err != nil {
				panic(err)
			}
			fmt.Println("from human size:", hsize)

			fmt.Println(conf.Value)
		case "duration":
			// Do duration based processing
			// sadly, there is way to transform duration correctly from PKL to stdlib durations
			// things like 10.min and 10.d will cause an error.
			// Compare https://pkl-lang.org/package-docs/pkl/0.27.0/base/index.html#DurationUnit
			// to https://cs.opensource.google/go/go/+/refs/tags/go1.23.3:src/time/format.go;l=1601
			durationStr := conf.Value.(string)
			if strings.Contains(durationStr, "min") {
				dur := strings.Split(durationStr, "min")
				mins, err := strconv.Atoi(dur[0])
				if err != nil {
					panic(err)
				}

				duration := time.Duration(mins) * time.Minute
				fmt.Println("duration:", duration)
				continue
			} else if strings.Contains(durationStr, "d") {
				dur := strings.Split(durationStr, "d")
				days, err := strconv.Atoi(dur[0])
				if err != nil {
					panic(err)
				}

				duration := time.Duration(days) * time.Hour * 24
				fmt.Println("duration:", duration)
				continue

			}

			duration, err := time.ParseDuration(conf.Value.(string))
			if err != nil {
				panic(err)
			}
			fmt.Println("duration:", duration)

		}
	}
}

func ReadConfigsV2(isFile bool, textPkl string) {
	evaluator, err := pkl.NewEvaluator(context.Background(), pkl.PreconfiguredOptions)
	if err != nil {
		panic(err)
	}
	defer evaluator.Close()

	var textOutput string
	if isFile {
		textOutput, err = evaluator.EvaluateOutputText(
			context.Background(),
			pkl.FileSource(ConfigurationFileV2))
		if err != nil {
			panic(err)
		}
		fmt.Println(textOutput)
	} else {
		textOutput, err = evaluator.EvaluateOutputText(
			context.Background(),
			pkl.TextSource(textPkl))
		if err != nil {
			panic(err)
		}
		fmt.Println(textOutput)
	}
	var cfg DocConfigsV2

	err = json.Unmarshal([]byte(textOutput), &cfg)
	if err != nil {
		panic(err)
	}

	cap_conf, ok := cfg.DocConfigsV2["test"]["cap1"]
	if !ok {
		panic("Unable to decode")
	}
	fmt.Printf("doc config:  %+v\n", cap_conf)

	for name, conf := range cap_conf {

		fmt.Println("name:", name)
		switch conf.ElemType {
		case "csv":
			// do more processing on list
			csvStr := conf.Value.(string)
			csvList := strings.Split(csvStr, ",")
			fmt.Println(csvList)
		case "boolean":
			// set some flag
			fmt.Println(conf.Value)
		case "size":
			sizeStr := conf.Value.(string)
			if strings.Contains(sizeStr, "ib") {

				// Do size based processing
				size, err := units.RAMInBytes(sizeStr)
				if err != nil {
					panic(err)
				}
				fmt.Println("from human size to bytes: ", strconv.FormatInt(size, 10))
				continue
			}

			// Do size based processing
			hsize, err := units.FromHumanSize(conf.Value.(string))
			if err != nil {
				panic(err)
			}
			fmt.Println("from human size:", hsize)

			fmt.Println(conf.Value)
		case "duration":
			// Do duration based processing
			// sadly, there is way to transform duration correctly from PKL to stdlib durations
			// things like 10.min and 10.d will cause an error.
			// Compare https://pkl-lang.org/package-docs/pkl/0.27.0/base/index.html#DurationUnit
			// to https://cs.opensource.google/go/go/+/refs/tags/go1.23.3:src/time/format.go;l=1601
			durationStr := conf.Value.(string)
			if strings.Contains(durationStr, "min") {
				dur := strings.Split(durationStr, "min")
				mins, err := strconv.Atoi(dur[0])
				if err != nil {
					panic(err)
				}

				duration := time.Duration(mins) * time.Minute
				fmt.Println("duration:", duration)
				continue
			} else if strings.Contains(durationStr, "d") {
				dur := strings.Split(durationStr, "d")
				days, err := strconv.Atoi(dur[0])
				if err != nil {
					panic(err)
				}

				duration := time.Duration(days) * time.Hour * 24
				fmt.Println("duration:", duration)
				continue

			}

			duration, err := time.ParseDuration(conf.Value.(string))
			if err != nil {
				panic(err)
			}
			fmt.Println("duration:", duration)
		case "array":
			fmt.Println("array:", conf.Value)
		default:
			break
		}
	}

}
func ReadConfigsV3(isFile bool, textPkl string) {
	evaluator, err := pkl.NewEvaluator(context.Background(), pkl.PreconfiguredOptions)
	if err != nil {
		panic(err)
	}
	defer evaluator.Close()

	var textOutput string
	if isFile {
		textOutput, err = evaluator.EvaluateOutputText(
			context.Background(),
			pkl.FileSource(ConfigurationFileV3))
		if err != nil {
			panic(err)
		}
		fmt.Println(textOutput)
	} else {
		textOutput, err = evaluator.EvaluateOutputText(
			context.Background(),
			pkl.TextSource(textPkl))
		if err != nil {
			panic(err)
		}
		fmt.Println(textOutput)
	}
	var cfg DocConfigsV3
	err = json.Unmarshal([]byte(textOutput), &cfg)
	if err != nil {
		panic(err)
	}

	cap_conf, ok := cfg.DocConfigsV3["test"]["cap1"]
	if !ok {
		panic("Unable to decode")
	}
	fmt.Printf("doc config:  %+v\n", cap_conf)

	// now you need to hardcode each value decoding logic.
	// Like DOC_TTL should be decoded as a Duration and nothin else!

}

func ReadConfigsV4(isFile bool, textPkl string) {
	evaluator, err := pkl.NewEvaluator(context.Background(), pkl.PreconfiguredOptions)
	if err != nil {
		panic(err)
	}
	defer evaluator.Close()

	var textOutput string
	if isFile {
		textOutput, err = evaluator.EvaluateOutputText(
			context.Background(),
			pkl.FileSource(ConfigurationFileV4))
		if err != nil {
			panic(err)
		}
		fmt.Println(textOutput)
	} else {
		textOutput, err = evaluator.EvaluateOutputText(
			context.Background(),
			pkl.TextSource(textPkl))
		if err != nil {
			panic(err)
		}
		fmt.Println(textOutput)
	}
	var cfg DocConfigsV4
	err = json.Unmarshal([]byte(textOutput), &cfg)
	if err != nil {
		panic(err)
	}

	cap_conf, ok := cfg.DocConfigsV4["test"]["cap1"]
	if !ok {
		panic("Unable to decode")
	}
	fmt.Printf("doc config:  %+v\n", cap_conf)

	// now you need to hardcode each value decoding logic.
	// Like DOC_TTL should be decoded as a Duration and nothin else!

}

func main() {
	ReadConfigs(true, "")
	ReadConfigsV2(true, "")
	ReadConfigsV3(true, "")
	ReadConfigsV4(true, "")
}
