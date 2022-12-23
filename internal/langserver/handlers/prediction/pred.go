package prediction

import (
	"encoding/json"
	"fmt"
	"os"
)

type Prediction struct {
	internalMapping map[string]any
}

func (p Prediction) loadJsonPrediction() error {
	config, err := os.ReadFile(fmt.Sprintf("./prediction/pred.json"))
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
		mapping := value.(map[string]string)

		i := 0
		keys := make([]string, len(mapping))
		for key := range mapping {
			keys[i] = key
			i++
		}
		return keys, nil
	}
	return nil, fmt.Errorf("error finding predicted resources")
}
