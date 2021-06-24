package calc

import d "gps-worker/domain"

func Sort(arr []*d.Position) []*d.Position {
	len := len(arr)

	if len == 1 {
		return arr
	}

	middle := int(len / 2)

	left := make([]*d.Position, middle)
	right := make([]*d.Position, len-middle)

	for i := 0; i < len; i++ {
		if i < middle {
			left[i] = arr[i]
		} else {
			right[i-middle] = arr[i]
		}
	}

	return merge(Sort(left), Sort(right))
}

func merge(left, right []*d.Position) (result []*d.Position) {
	result = make([]*d.Position, len(left)+len(right))

	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0].Distance < right[0].Distance {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}

		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
