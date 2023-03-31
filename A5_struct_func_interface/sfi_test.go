package a5structfuncinterface

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectanbgle := Rectangle{Width: 10.0, Height: 10.0}
	got := rectanbgle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestArea_noStruct(t *testing.T) {
	checkArea := func(shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectanbgle := Rectangle{10, 10}
		want := 100.0
		checkArea(rectanbgle, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := math.Pi * 100
		checkArea(circle, want)
	})
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: math.Pi * 100},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}
	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%v %v got %v want %v", tt.name, tt.shape, got, tt.hasArea)
			}
		})
	}
}
