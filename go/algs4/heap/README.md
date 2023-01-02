# Priority Queues

Stack: Remove the item most recently added.
Queue: Remove the item least recently added.
Priority Queue: Remove the largest (or smallest) item.

## Basic API

Requirement: items must be comparable

- pq.Insert(key T)
- pq.DelMax() -> key T: Return and remove the largest key
- pq.IsEmpty() -> bool
- pq.Max() -> key T: Return the highest Key
- pq.Size() -> int: number of entries in the pq
