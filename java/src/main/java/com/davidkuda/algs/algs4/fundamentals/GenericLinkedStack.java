package com.davidkuda.algs.algs4.fundamentals;

public class GenericLinkedStack<Item> {

  private Node top = null;

  private class Node {
    Item item;
    Node next;
  }

  public boolean isEmpty() {
    return top == null;
  }

  public void push(Item i) {
    Node n = new Node();
    n.item = i;
    n.next = top;
    top = n;
  }

  public Item pop() {
    Item tempTop = top.item;
    top = top.next;
    return tempTop;
  }
}
