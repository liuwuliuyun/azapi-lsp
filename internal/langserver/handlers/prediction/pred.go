package prediction

type Prediction struct {
	internalMapping map[string][]string
}

var InternalPred *Prediction

func (p Prediction) Top3PredResult(parentResource string) ([]string, error) {
	if value, ok := p.internalMapping[parentResource]; ok {
		return value, nil
	}
	return make([]string, 0), nil
}

func InitializePrediction() error {
	if InternalPred == nil {
		InternalPred = &Prediction{}
		InternalPred.internalMapping = getMapping()
	}
	return nil
}
