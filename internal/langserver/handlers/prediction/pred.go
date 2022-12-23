package prediction

import (
	"encoding/json"
	"fmt"
	"os"
)

type Prediction struct {
	internalMapping map[string]any
}

var Pred *Prediction

func (p Prediction) loadJsonPrediction() error {
	config, err := os.ReadFile(fmt.Sprintf("./prediction/processed.json"))
	if err != nil {
		return err
	}

	err = json.Unmarshal(config, &p.internalMapping)
	if err != nil {
		return err
	}
	return nil
}

func (p Prediction) Top3PredResult(parentResource string) ([]string, error) {
	if p.internalMapping == nil {
		err := p.loadJsonPrediction()
		if err != nil {
			return nil, err
		}
	}
	if value, ok := p.internalMapping[parentResource]; ok {
		return value.([]string), nil
	}
	return make([]string, 0), nil
}

func InitializePrediction() error {
	if Pred == nil {
		err := Pred.loadJsonPrediction()
		if err != nil {
			return err
		}
	}
	return nil
}
