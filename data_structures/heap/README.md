# Priority Queues

Stack: Remove the item most recently added.
Queue: Remove the item least recently added.
Priority Queue: Remove the largest (or smallest) item.

## Further Resources

- [Fantastic Video Explanation](https://www.youtube.com/watch?v=t0Cq6tVNRBA)

## Basic API

Requirement: items must be comparable, and immutable

- pq.Insert(key T)
- pq.DelMax() -> key T: Return and remove the largest key
- pq.IsEmpty() -> bool
- pq.Max() -> key T: Return the highest Key
- pq.Size() -> int: number of entries in the pq

other operations:
- Remove an arbitrary item
- Change the priority of an item

## Heaps

- A Heap can either start at array[0] or at array[1]
- if array[1]:
    - parent of node is at k/2
    - children of node is at 2k and 2k+1
- if array[0]:
    - parent -> (k-1) / 2
    - leftChild -> k*2 + 1
    - rightChild -> k*2 + 2
- restore order by exchanging position with parent
- swim operation: move child upwards until in right position
- insert -> at the end of array, then swim new value
- sink operation: move parent downward until in right position by exchanging with the largest of its two children
- DelMax -> remove pq[0], put pq[N] at pq[0], swim until order restored
- Time Complexity of the API:
    - Insert: O(logN)
    - DelMax: O(logN)
    - Max: O(1)

## Heap Sort

- Rearrange an array into a heap
- Repeatedly place the max element to arr[N] (N = len(arr)-1)
- Build heap with the "bottom-up" approach ("divide and conquer"), start with smallest sub-heap, and work until heap.items[0]

```go
// first pass: heap construction using bottom-up method
N = len(pq.items) - 1
for k := N/2; k >= 1; k-- {
    pq.sink(k)
}

// second pass: sort array by placing largest item at the end
for N > 0 {
    pq.exchange(0, N)
    N--
    pq.sink(0)
}
```
