package com.davidkuda.algs.algs4.fundamentals;

public class LinkedStackOfStrings {

  private Node top;

  private class Node {
    String item;
    Node next;
  }

  public boolean isEmpty() {
    return top == null;
  }

  public void push(String element) {
    Node newNode = new Node();
    newNode.item = element;
    newNode.next = top;
    top = newNode;
  }

  public String pop() {
    String topValue = top.item;
    top = top.next;
    return topValue;
  } 
}
