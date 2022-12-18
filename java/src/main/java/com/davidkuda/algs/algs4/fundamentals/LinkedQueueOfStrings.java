package com.davidkuda.algs.algs4.fundamentals;

public class LinkedQueueOfStrings {

  private Node first;
  private Node last;
  private int size;

  // Same as in Queue:
  private class Node {
    String item;
    Node next;
  }

  public int getSize() {
    return size;
  }

  public boolean isEmpty() {
    return first == null;
  }

  public void enqueue(String item) {
    Node n = new Node();
    n.item = item;
    n.next = null;

    if (isEmpty()) {
      first = n;
      last = n;
      size++;
    } else {
      last.next = n;
      last = n;
      size++;
    }
  }

  // Queue.dequeue() is identical to Stack.pop()
  public String dequeue() {
    String item = first.item;
    first = first.next;
    size--;

    if (isEmpty()) {
      last = null;
    }

    return item;
  }
}
