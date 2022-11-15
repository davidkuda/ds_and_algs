package com.davidkuda.algs.algs4.fundamentals;

public class FixedCapacityStackOfStrings {

  private String[] stack;
  private int N = 0;

  public FixedCapacityStackOfStrings(int capacity) {
    stack = new String[capacity];
  }

  public boolean isEmpty() {
    return N == 0;
  }
  
  public void push(String element) {
    stack[N++] = element;
  }

  public String pop() {
    return stack[--N];   
  }
}
