package com.davidkuda.algs.algs4.fundamentals;

import java.util.Iterator;

public class IteratorLinkedStack<Item> implements Iterable<Item> {

  Node top = null;

  private class Node {
    Item item;
    Node next;

    private Node(Item i) {
      this.item = i;
    }
  }

  public boolean isEmpty() {
    return top == null;
  }

  public void push(Item item) {
    Node newTop = new Node(item);
    newTop.next = top;
    top = newTop;
  }

  public Item pop() {
    Item topItem = top.item;
    top = top.next;
    return topItem;
  }

  public Iterator<Item> iterator() {
    return new LinkedStackIterator(top);
  }

  private class LinkedStackIterator implements Iterator<Item> {

    Node current;

    public LinkedStackIterator(Node first) {
      current = first;
    }

    @Override
    public boolean hasNext() {
      return current != null;
    }

    @Override
    public Item next() {
      Item iItem = current.item;
      current = current.next;
      return iItem;
    }

    public void remove() {
      throw new UnsupportedOperationException();
    }
  }
}
