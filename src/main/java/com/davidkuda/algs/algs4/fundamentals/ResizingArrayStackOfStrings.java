package com.davidkuda.algs.algs4.fundamentals;

/**
 * Array will always be between 25% and 100% full
 * Space used always at max times two of items in the array
 * Mostly O(1) operations, but as soon as resizing happens O(n)
 * Consider "Amortized Analysis" -> Total cost averaged over all operations
 * You will have many O(1) operations and only little O(n) operations; The
 * amortized time will nevertheless be O(1), but worst case will be O(n)
 */
public class ResizingArrayStackOfStrings {

  int N = 0;
  String[] stack;

  public ResizingArrayStackOfStrings() {
    stack = new String[1];
  }

  public void push(String item) {
    // Double the size when array is full:
    if (N == stack.length) {
      resize( 2 * stack.length );
    }
    stack[N++] = item;
  }

  public String pop() {
    // Halve the array when one-quarter full:
    if (N == stack.length / 4) {
      resize( stack.length / 2 );
    }
    return stack[--N];
  }

  private void resize(int capacity) {
    String[] copy = new String[capacity];
    for (int i = 0; i < N; i++) {
      copy[i] = stack[i];
    }
    stack = copy;
  }
}
