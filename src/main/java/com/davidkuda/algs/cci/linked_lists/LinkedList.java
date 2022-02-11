package com.davidkuda.algs.cci.linked_lists;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;

public class LinkedList<Item> {

  Node first;
  Node last;
  private int size;

  class Node {
    Item item;
    Node next;
  }

  public int getSize() {
    return size;
  }

  public boolean isEmpty() {
    return size == 0;
  }

  public void insertAtBeginning(Item item) {
    if (isEmpty()) {
      // Note: This check is only required if there is the field "last".
      first = new Node();
      first.item = item;
      last = first;
      size++;
      return;
    }

    Node oldFirst = first;
    first = new Node();
    first.item = item;
    first.next = oldFirst;
    size++;
  }

  public void removeFromBeginning() {
    // "first" becomes an orphan; Java memory management will reclaim the space
    // ("Garbage Collection")
    first = first.next;
    size--;

    if (isEmpty()) {
      last = null;
    }
  }

  public void insertAtEnd(Item item) {
    Node oldLast = last;
    last = new Node();
    last.item = item;
    oldLast.next = last;
    size++;
  }

  public void removeDups() {
    HashSet<Item> set = new HashSet<Item>();

    Node current = first;
    Node previous = null;

    while (current != null) {
      if (set.contains(current.item)) {
        previous.next = current.next;
      } else {
        set.add(current.item);
        previous = current;
      }
    
      current = current.next;
    }
  }

  private void removeNode() {}

  private void traverse() {
    for (Node x = first; x != null; x = x.next) {
      System.out.println(x.item);
    }
  }
  
}