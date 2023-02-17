package exercises

import (
	"errors"
	"sort"
)

/*
Logic:

Sort the slice so we can binary search
Go through the Slice once
Check for each value, the second value that we would need (target - value)
After that, we can use searchInts which binary searches for that value in the slice, returning the index in which the value is OR the index where it should be added
Finally we just validate if that index is correct and return
*/
func TwoSumValue(listValues []int, target int) ([]int, error) {

	if len(listValues) == 0 {
		return nil, errors.New("empty slice")
	}

	result := make([]int, 2)

	sort.Ints(listValues) // Sorting first allows binary search which is the best approach

	for index, value := range listValues {
		valueToSearch := target - value // If this value exists in the remaining array, we achieved the result

		indexValueToSearch := sort.SearchInts(listValues, valueToSearch) // SearchInts searches for x in a sorted slice of ints and returns the index it is in.
		if index == indexValueToSearch {
			continue
		}
		if isIndexValid(listValues, indexValueToSearch) && valueToSearch == listValues[indexValueToSearch] {
			result[0] = index
			result[1] = indexValueToSearch
			return result, nil
		}
	}

	return nil, errors.New("404")
}

func isIndexValid(slice []int, index int) bool {
	return index < len(slice)
}
