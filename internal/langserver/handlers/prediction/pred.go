package prediction

import (
	"encoding/json"
	"fmt"
	"os"
)

type Prediction struct {
	internalMapping map[string][]string
}

var InternalPred *Prediction

func LoadJsonPrediction(filePath string, predictionItem *Prediction) error {
	defaultFilePath := "./prediction/processed.json"
	if len(filePath) > 0 {
		defaultFilePath = filePath
	}
	config, err := os.ReadFile(fmt.Sprintf(defaultFilePath))
	if err != nil {
		return err
	}

	err = json.Unmarshal(config, &predictionItem.internalMapping)
	if err != nil {
		return err
	}
	return nil
}

func (p Prediction) Top3PredResult(parentResource string) ([]string, error) {
	if value, ok := p.internalMapping[parentResource]; ok {
		return value, nil
	}
	return make([]string, 0), nil
}

func InitializePrediction() error {
	if InternalPred == nil {
		InternalPred = &Prediction{
			internalMapping: map[string][]string{},
		}
		err := LoadJsonPrediction("", InternalPred)
		if err != nil {
			return err
		}
	}
	return nil
}
