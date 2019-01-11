package createSeeds

import (
	"fmt"

	"github.com/weet/model"
)

func CreateSeedSamples() {

	sample_infos := []map[string]string{
		map[string]string{},
		map[string]string{},
		map[string]string{},
	}

	for _, info := range sample_infos {
		createSample(model.Sample{
			Name:    info["Name"],
			Comment: info["Comment"],
		})
	}
}

func createSample(sample model.Sample) {
	sample, err := model.CreateSample(sample)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sample created\n")
}
