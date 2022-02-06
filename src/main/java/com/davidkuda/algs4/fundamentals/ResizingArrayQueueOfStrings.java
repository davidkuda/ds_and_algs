package com.davidkuda.algs4.fundamentals;

public class ResizingArrayQueueOfStrings {

  String[] queue = new String[1];
  int head = 0;
  int tail = 0;

  public int getSize() {
    return tail - head;
  }

  public boolean isEmpty() {
    return getSize() == 0;
  }

  public void enqueue(String item) {
    if (queue.length == tail) {
      resize(queue.length * 2);
    }

    queue[tail] = item;
    tail++;
  }

  public String dequeue() {

    String item = queue[head];
    queue[head] = null;
    head++;
    return item;
  }

  private void resize(int size) {
    String[] oldQueue = queue;
    queue = new String[size];

    int newCounter = 0;
    for (int i = 0; i < oldQueue.length; i++) {
      if (oldQueue[i] == null) {
        continue;
      }
      queue[newCounter] = oldQueue[i];
      newCounter++;
    }
    head = 0;
    tail = newCounter;
  }
}
