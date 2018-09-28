package diffcheck

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	a = "b137bfcd6bc4e62ad4265e752b91ecd37b0417b4"
	b = "0c5883bf7e55c4cfd90c2b889343c2c41ff9fc6f"
	c = "5e36b9830705588eaff106c20d6f3d86804dcc8d"
	d = "49ce7d85a5fd5cca9a23cd705c2e7e0f582124e7"
	e = "8ea373e70a63a60800b1ff4722c01fbc13a0837e"
)

func TestDiffCommonPrefix(t *testing.T) {
	type TestCase struct {
		Name     string
		Array1   []string
		Array2   []string
		Expected []string
	}

	dc := New()

	for i, tc := range []TestCase{
		{"Whole", []string{a, b, c, d}, []string{a, b, c, e}, []string{a, b, c}},
		{"len(Array1) < len(Array2)", []string{a, b, c}, []string{a, c, e, e}, []string{a}},
		{"len(Array1) > len(Array2)", []string{a, d, a, c, e}, []string{a, d, a}, []string{a, d, a}},
		{"Null", []string{}, []string{a, e}, []string{}},
	} {
		actual := dc.DiffCommonPrefix(tc.Array1, tc.Array2)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Test case #%d, %s", i, tc.Name))
	}
}

func TestDiffCommonSuffix(t *testing.T) {
	type TestCase struct {
		Name     string
		Array1   []string
		Array2   []string
		Expected []string
	}
	a := "b137bfcd6bc4e62ad4265e752b91ecd37b0417b4"
	b := "0c5883bf7e55c4cfd90c2b889343c2c41ff9fc6f"
	c := "5e36b9830705588eaff106c20d6f3d86804dcc8d"
	d := "49ce7d85a5fd5cca9a23cd705c2e7e0f582124e7"
	e := "8ea373e70a63a60800b1ff4722c01fbc13a0837e"

	dc := New()

	for i, tc := range []TestCase{
		{"Whole", []string{a, b, c, d}, []string{c, b, c, d}, []string{b, c, d}},
		{"len(Array1) < len(Array2)", []string{a, b, e}, []string{a, c, e, e}, []string{e}},
		{"len(Array1) > len(Array2)", []string{a, d, a, c, e, a}, []string{a, d, e, a}, []string{e, a}},
		{"Null", []string{}, []string{a, e}, []string{}},
	} {
		actual := dc.DiffCommonSuffix(tc.Array1, tc.Array2)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Test case #%d, %s", i, tc.Name))
	}
}
