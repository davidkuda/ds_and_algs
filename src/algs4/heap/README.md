# Priority Queues

Stack: Remove the item most recently added.
Queue: Remove the item least recently added.
Priority Queue: Remove the largest (or smallest) item.

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

- parent of node is at k/2
- children of node is at 2k and 2k+1
- restore order by exchanging position with parent
- swim operation: move child upwards until in right position
- insert -> at the end of array, then swim new value
- sink operation: move parent downward until in right position by exchanging with the largest of its two children
- DelMax -> remove pq[0], put pq[N] at pq[0], swim until order restored
- Time Complexity of the API:
    - Insert: O(logN)
    - DelMax: O(logN)
    - Max: O(1)

```go

```
