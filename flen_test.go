package flen

import (
	"os"
	"reflect"
	"testing"
)

var (
	sampleFuncLens *FuncLens
	sampleCode     = `
func single() {
	println("hello single")
}

func double() {
	println("hello double")
	println("hello double")
}

func trouble() {
	println("hello trouble")
	println("hello trouble")
	println("hello trouble")
}
`
)

func TestFuncLensPrint(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		flens *FuncLens
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.flens.Print()
	}
}

func TestFuncLensLen(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		flens *FuncLens
		// Expected results.
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.flens.Len(); got != tt.want {
			t.Errorf("%q. FuncLens.Len() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFuncLensLess(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		flens *FuncLens
		// Parameters.
		i int
		j int
		// Expected results.
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.flens.Less(tt.i, tt.j); got != tt.want {
			t.Errorf("%q. FuncLens.Less() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFuncLensSwap(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		flens *FuncLens
		// Parameters.
		i int
		j int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.flens.Swap(tt.i, tt.j)
	}
}

func TestFuncLensComputeHistogram(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		flens *FuncLens
		// Expected results.
		want []int
	}{}
	for _, tt := range tests {
		if got := tt.flens.computeHistogram(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FuncLens.computeHistogram() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFuncLensQuery(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		flens *FuncLens
		// Parameters.
		lowerLimit int
		upperLimit int
		// Expected results.
		want FuncLens
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.flens.Query(tt.lowerLimit, tt.upperLimit); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FuncLens.Query() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFuncLensGetZeroLenFuncs(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		flens *FuncLens
		// Expected results.
		want FuncLens
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.flens.GetZeroLenFuncs(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FuncLens.GetZeroLenFuncs() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFuncLensGetExternallyImplementedFuncs(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		flens *FuncLens
		// Expected results.
		want FuncLens
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.flens.GetExternallyImplementedFuncs(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FuncLens.GetExternallyImplementedFuncs() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFuncLensDisplayHistogram(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		flens *FuncLens
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.flens.DisplayHistogram()
	}
}

func TestFuncLensComputePercentiles(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver.
		flens *FuncLens
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.flens.ComputePercentiles()
	}
}

func TestGenerateFuncLens(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		pkg     string
		options *Options
		// Expected results.
		want    FuncLens
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := GenerateFuncLens(tt.pkg, tt.options)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. GenerateFuncLens() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GenerateFuncLens() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetPkgPath(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		pkgname string
		// Expected results.
		want  string
		want1 *os.PathError
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, got1 := getPkgPath(tt.pkgname)
		if got != tt.want {
			t.Errorf("%q. getPkgPath() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. getPkgPath() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

/*
func TestMain(m *testing.M) {
	os.Exit()
}
*/
