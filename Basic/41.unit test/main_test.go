package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cube           Cube    = Cube{4}
	volumeExpected float64 = 64
	areaExpected   float64 = 96
	aroundExpected float64 = 48
)

func TestCalculateVolume(t *testing.T) {
	t.Logf("Volume: %2.f", cube.Volume())

	if cube.Volume() != volumeExpected {
		t.Errorf("Wrong! volume must be %2.f", volumeExpected)
	}
}

func TestCalculateArea(t *testing.T) {
	t.Logf("Area: %2.f", cube.Area())

	if cube.Area() != areaExpected {
		t.Errorf("Wrong! area must be %2.f", areaExpected)
	}
}

func TestCalculateAround(t *testing.T) {
	t.Logf("Area: %2.f", cube.Around())

	if cube.Around() != aroundExpected {
		t.Errorf("Wrong! area must be %2.f", aroundExpected)
	}
}

func BenchmarkCalculateAround(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cube.Around()
	}
}

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, cube.Volume(), volumeExpected, "perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, cube.Area(), areaExpected, "perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, cube.Around(), aroundExpected, "perhitungan keliling salah")
}
