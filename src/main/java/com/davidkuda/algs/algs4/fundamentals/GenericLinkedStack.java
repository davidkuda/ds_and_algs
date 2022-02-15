package com.davidkuda.algs.algs4.fundamentals;

public class GenericLinkedStack<Item> {

  int size = 0;
  Node top;

  private class Node {
    Item item;
    Node next;
  }

  public boolean isEmpty() {
    return size == 0;
  }

  public void push(Item i) {
    Node n = new Node();
    n.item = i;
    n.next = top;
    top = n;
    size++;
  }

  public Item pop() {
    Item tempTop = top.item;
    top = top.next;
    size--;
    return tempTop;
  }
}
