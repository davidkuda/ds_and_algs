package main

import "reflect"

type Date struct {
	Month int
	Day   int
	Year  int
}

func (x *Date) equals(y *Date) bool {
	if x == y {
		return true
	}

	if y == nil {
		return true
	}

	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		return false
	}

	if x.Day != y.Day || x.Month != y.Month || x.Year != y.Year {
		return false
	}

	return true
}
