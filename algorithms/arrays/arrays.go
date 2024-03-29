// note: an array stands for any sequence

package arrays

func reverse(a []int) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

// rotate left k times, example:
// rotateLeft([]int{1,2,3}, 1) == []int{2, 3, 1}
// rotateLeft([]int{1,2,3}, 7) == []int{2, 3, 1}
func rotateLeft(a []int, k int) {
	k = k % len(a)
	reverse(a[:k])
	reverse(a[k:])
	reverse(a)
}

// rotate right k times, example:
// rotateRight([]int{1,2,3,4,5}, 2) == []int{4, 5, 1, 2, 3}
// rotateRight([]int{1,2,3,4,5}, 7) == []int{4, 5, 1, 2, 3}
func rotateRight(a []int, k int) {
	k = k % len(a)
	reverse(a)
	reverse(a[:k])
	reverse(a[k:])
}
