package main

import "reflect"

type Date struct {
	Month int
	Day   int
	Year  int
}

/**
Standard Recipe for object.equals(object):
- check same address first, avoid a lot of code if true
- check against nil
- check type
- compare each field
  - if field is primitive type, use ==
  - if field is object, use its equals method
  - if field is array, check each value of array

Best Practices:

*/
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
