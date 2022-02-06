package com.davidkuda.algs4.fundamentals;

import org.junit.Assert;
import org.junit.Test;

public class QueueTests {

  @Test
  public void testLinkedQueueOfStrings() {

    // Given:
    LinkedQueueOfStrings queue = new LinkedQueueOfStrings();

    // When:
    queue.enqueue("First");
    queue.enqueue("Second");
    queue.enqueue("Third");

    // Then:
    Assert.assertEquals("First", queue.dequeue());
    Assert.assertEquals("Second", queue.dequeue());
    Assert.assertEquals("Third", queue.dequeue());
  }

  @Test
  public void testResizingArrayQueueOfStrings() {

    // Given:
    ResizingArrayQueueOfStrings queue = new ResizingArrayQueueOfStrings();

    // When:
    queue.enqueue("First");
    queue.enqueue("Second");
    queue.enqueue("Third");

    // Then:
    Assert.assertEquals("First", queue.dequeue());
    Assert.assertEquals("Second", queue.dequeue());
    Assert.assertEquals("Third", queue.dequeue());
  }
}
