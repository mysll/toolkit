package toolkit

import "math/rand"

// Flatten returns an array a single level deep
func Flatten[T any](collection [][]T) []T {
	totalLen := 0
	for i := range collection {
		totalLen += len(collection[i])
	}

	result := make([]T, 0, totalLen)
	for i := range collection {
		result = append(result, collection[i]...)
	}

	return result
}

//Shuffle returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm
func Shuffle[T any](collections []T) []T {
	rand.Shuffle(len(collections), func(i, j int) {
		collections[i], collections[j] = collections[j], collections[i]
	})
	return collections
}

// Reverse reverses array so that the first element becomes the last, the second element becomes the second to last, and so on
func Reverse[T any](collection []T) []T {
	length := len(collection)
	half := length / 2

	for i := 0; i < half; i = i + 1 {
		j := length - 1 - i
		collection[i], collection[j] = collection[j], collection[i]
	}

	return collection
}
