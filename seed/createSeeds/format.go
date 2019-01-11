package createSeeds

import (
	"fmt"

	"github.com/weet/model"
)

func CreateSeedFormat() {

	format_infos := []map[string]string{
		map[string]string{
			"Name": "友達",
		},
		map[string]string{
			"Name": "恋愛",
		},
		map[string]string{
			"Name": "婚活",
		},
		map[string]string{
			"Name": "ルームメイト",
		},
	}

	for _, info := range format_infos {
		createFormat(model.Format{
			Name: info["Name"],
		})
	}
}

func createFormat(format model.Format) {
	format, err := model.CreateFormat(format)
	if err != nil {
		panic(err)
	}
	fmt.Printf("format created\n")
}
