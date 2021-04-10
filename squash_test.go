// Use of this source code is governed by EUPL-1.2-or-later license
// that can be found in the LICENSE file or read at
// https://joinup.ec.europa.eu/collection/eupl/eupl-text-eupl-12
package squash_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/manuelcoppotelli/squash"
)

func TestIntegersNotInt(t *testing.T) {
	res, err := squash.Integers("")
	if err == nil {
		t.Errorf("Function was expected to return an error instead of %v", res)
	}
}

func TestIntegersNotArray(t *testing.T) {
	res, err := squash.Integers(3)
	if err == nil {
		t.Errorf("Function was expected to return an error instead of %v", res)
	}
}

func TestIntegersNotIntArray(t *testing.T) {
	res, err := squash.Integers([]string{"test"})
	if err == nil {
		t.Errorf("Function was expected to return an error instead of %v", res)
	}
}

func TestIntegersArray(t *testing.T) {
	exp := []int{3}

	res, err := squash.Integers([]int{3})
	if err != nil {
		t.Errorf("Function returned an error: %v", err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v returned %v", exp, res)
	}
}

func TestIntegersArrayMultiple(t *testing.T) {
	exp := []int{3, 5}

	res, err := squash.Integers([]int{3, 5})
	if err != nil {
		t.Errorf("Function returned an error: %v", err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v returned %v", exp, res)
	}
}

func TestIntegersInterface(t *testing.T) {
	exp := []int{1, 2}
	res, err := squash.Integers([]interface{}{1, 2})
	if err != nil {
		t.Errorf("Function returned an error: %v", err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v returned %v", exp, res)
	}
}

func TestIntegersInterfaceArray(t *testing.T) {
	exp := []int{1, 2, 3}
	res, err := squash.Integers([]interface{}{1, 2, []int{3}})
	if err != nil {
		t.Errorf("Function returned an error: %v", err)
	}
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("Expected %v returned %v", exp, res)
	}
}

func ExampleIntegers() {
	// Input:
	// [[1, 2, [3]], 4]
	res, err := squash.Integers([]interface{}{[]interface{}{1, 2, []int{3}}, 4})

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	// Output:
	// [1 2 3 4]
}
