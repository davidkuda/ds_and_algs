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
    Node oldLast = last;

    last = new Node();
    last.item = item;
    last.next = null;

    if (isEmpty()) {
      first = last;
    } else  {
      oldLast.next = last;
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
