package squareroot

import "errors"

func SquareRoot(number int) (int, error) {
	if number < 1 {
		return 0, errors.New("input must be a positive whole number")
	}
	low := 1
	high := number

	for low <= high {
		mid := low + (high-low)/2
		square := mid * mid

		if square == number {
			return mid, nil
		}

		if square > number {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, errors.New("not a perfect square")
}
