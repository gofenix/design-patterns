package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	dr := Director{Builder: CarBuilder{}}
	benzi := dr.Builder.SetType("SUV").AddBrand("奔驰").PaintColor("white").Build()
	benzi.Move()

	bmw := dr.Builder.SetType("sporting").AddBrand("宝马").PaintColor("red").Build()
	bmw.Move()
}
