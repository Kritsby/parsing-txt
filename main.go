package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type TestRecord struct {
	Data []Data `json:"data"`
}

type Data struct {
	Code string `json:"code"`
	Name Name   `json:"name"`
}

type Name struct {
	Ru string `json:"ru"`
	En string `json:"en"`
}

type result struct {
	Code string `json:"code"`
	Ru   string `json:"ru"`
	En   string `json:"en"`
}

func main() {
	file, err := os.Open("country.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	bytes1, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println(err)
		return
	}

	var testRecord TestRecord
	err = json.Unmarshal(bytes1, &testRecord)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(testRecord)

	file1, err := os.Create("country.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var resultSlcie []result

	for _, data := range testRecord.Data {
		resultSlcie = append(resultSlcie, result{
			Code: data.Code,
			Ru:   data.Name.Ru,
			En:   data.Name.En,
		})
	}

	b, err := json.MarshalIndent(resultSlcie, "", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file1.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	//cat country.json | jq -cr '.[]' | sed 's/\\[tn]//g' > output.json
}
