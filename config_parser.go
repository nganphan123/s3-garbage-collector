package main

import (
	"fmt"
	"os"

	"github.com/nganphan123/s3-garbage-collector/types"
	"sigs.k8s.io/yaml"
)

func parser() {
	configContent, err := os.ReadFile("./example/mock_config.yaml")
	if err != nil {
		fmt.Print(err)
		return
	}

	var config types.DeleteConfig
	err = yaml.Unmarshal(configContent, &config)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("%v", config)
}
