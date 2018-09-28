package diffcheck

import (
	"time"
)

// Operation defines the operation of a diff item.
type Operation int8

const (
	DiffDelete Operation = -1
	DiffInsert Operation = 1
	DiffEqual  Operation = 0
)

// Diff represents one diff operation
type Diff struct {
	Type    Operation
	HashMap []string
}

// Main function of Diff Check Algorithm
func (dc *DiffClient) DiffMain(bf_array, af_array []string) []Diff {
	var deadline time.Time
	if dc.DiffTimeout > 0 {
		deadline = time.Now().Add(dc.DiffTimeout)
	}
	return dc.diffMainArrays(bf_array, af_array, deadline)
}

func (dc *DiffClient) diffMainArrays(bf_array, af_array []string, deadline time.Time) []Diff {
	if arraysEqual(bf_array, af_array) {
		var diffs []Diff
		if len(bf_array) > 0 {
			diffs = append(diffs, Diff{DiffEqual, bf_array})
		}
		return diffs
	}
	return nil
}

// DiffCommonPrefix determines the common prefix length of two strings.
func (dmp *DiffClient) DiffCommonPrefix(ar1, ar2 []string) []string {
	// Unused in this code, but retained for interface compatibility.
	return commonPrefixArray(ar1, ar2)
}

// DiffCommonSuffix determines the common suffix length of two strings.
func (dmp *DiffClient) DiffCommonSuffix(ar1, ar2 []string) []string {
	// Unused in this code, but retained for interface compatibility.
	return commonSuffixArray(ar1, ar2)
}

// commonPrefixArray returns the common array prefix of two arrays
func commonPrefixArray(ar1, ar2 []string) []string {
	n := 0
	for ; n < len(ar1) && n < len(ar2); n++ {
		if ar1[n] != ar2[n] {
			return ar1[:n]
		}
	}
	return ar1[:n]
}

// commonPrefixArray returns the common array suffix of two arrays
func commonSuffixArray(ar1, ar2 []string) []string {
	i1 := len(ar1)
	i2 := len(ar2)
	for n := 0; ; n++ {
		i1--
		i2--
		if i1 < 0 || i2 < 0 || ar1[i1] != ar2[i2] {
			return ar1[i1+1:]
		}
	}
}
