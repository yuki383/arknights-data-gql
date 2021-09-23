package operators

import (
	"encoding/json"
	"os"
)

type Operator struct {
	Name string `json:"name"`
	Class string `json:"class"`
	Available []string `json:"available"`
	Rarity int `json:"rarity"`
}

const operatorFilePath = "./data/operators.json"

var operators []Operator

func init() {
	data, err := os.ReadFile(operatorFilePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &operators)
	if err != nil {
		panic(err)
	}
}

func GetOperators() []Operator {
	return operators
}


