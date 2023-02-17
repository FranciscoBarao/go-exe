package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumValue(t *testing.T) {

	list_values := []int{1, 3, 2, -7, 5}
	target := 7
	result, _ := TwoSumValue(list_values, target)
	expected := []int{2, 4}
	assert.Equal(t, result, expected, "#1 -> Result should be [2,4]")

	list_values = []int{1, 3, 2, -7, 5}
	target = 10
	_, err := TwoSumValue(list_values, target)
	expected2 := "404"
	assert.Equal(t, err.Error(), expected2, "#2 -> Result should not exist ")
}
