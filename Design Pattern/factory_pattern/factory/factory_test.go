package factory_test

import (
	"factory/factory"
	"testing"
)

func TestGetEngine(t *testing.T) {
	t.Logf("Get Engine")
	carEngine := factory.GetEngine("car")
	
	if carEngine != factory.Car {
		t.Errorf("Fail: carEngine must be: %v ", factory.Car)
	}
}
