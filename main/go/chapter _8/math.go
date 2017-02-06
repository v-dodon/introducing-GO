package math

// Finds the average of  a series of numbers
func Average(xs []float64) float64 {
	if(len(xs) == 0){
		return 0
	}
	total := float64(0)
	for _, x:= range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Finds the minimum value of a series of numbers
func Min(xs []float64) float64 {
	if(len(xs) == 0){
		panic("Your array is empty")
	}
	min := xs[0]
	for _, x:= range xs {
		if(x < min){
			min = x
		}
	}
	return min
}

// Finds the maximum value of a series of numbers
func Max(xs []float64) float64 {
	max := float64(0)
	for _, x:= range xs {
		if(x > max){
			max = x
		}
	}
	return max
}