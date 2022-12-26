package prediction

import (
	"reflect"
	"testing"
)

func TestPrediction_Top3PredResult(t *testing.T) {

	InternalPred = &Prediction{
		internalMapping: map[string][]string{},
	}
	err := LoadJsonPrediction("./processed.json", InternalPred)
	if err != nil {
		t.Fatal(err)
	}
	prevResource := "azurerm_resource_group"
	predictedResourceList, err2 := InternalPred.Top3PredResult(prevResource)
	if err2 != nil {
		t.Fatal(err2)
	}
	expectedResourceList := []string{"azurerm_storage_account", "azurerm_virtual_network", "azurerm_data_factory"}
	if !reflect.DeepEqual(predictedResourceList, expectedResourceList) {
		t.Fatalf("expected values are %s, however got %s", expectedResourceList, predictedResourceList)
	}
}
