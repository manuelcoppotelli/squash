// Use of this source code is governed by EUPL-1.2-or-later license
// that can be found in the LICENSE file or read at
// https://joinup.ec.europa.eu/collection/eupl/eupl-text-eupl-12

// Package squash implements a stategy to turn an array of arbitrarily nested
// arrays into one flat array.
package squash

import (
	"errors"
)

// Flatten an array of integers
//
// It returns a new flattened array
func Integers(in interface{}) ([]int, error) {
	if _, ok := in.(int); !ok {
		return integers([]int{}, in)
	}

	return nil, errors.New("Only array of int type supported")
}

// perform the actual flattening through recursion
func integers(res []int, in interface{}) ([]int, error) {
	var err error

	switch val := in.(type) {
	case int:
		res = append(res, val)
	case []int:
		res = append(res, val...)
	case []interface{}:
		for idx := range val {
			res, err = integers(res, val[idx])
			if err != nil {
				return nil, errors.New("Array can only contain int or []int")
			}
		}
	default:
		return nil, errors.New("Only array of int type supported")
	}

	return res, nil
}
