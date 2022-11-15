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
    queue.enqueue("Fourth");
    queue.enqueue("Fifth");
    queue.enqueue("Sixth");
    queue.enqueue("Seventh");
    queue.enqueue("Eigth");
    queue.enqueue("Ninth");

    // Then:
    Assert.assertEquals("First", queue.dequeue());
    Assert.assertEquals("Second", queue.dequeue());
    Assert.assertEquals("Third", queue.dequeue());
    Assert.assertEquals("Fourth", queue.dequeue());
    Assert.assertEquals("Fifth", queue.dequeue());
    Assert.assertEquals("Sixth", queue.dequeue());
    Assert.assertEquals("Seventh", queue.dequeue());
    Assert.assertEquals("Eigth", queue.dequeue());
    Assert.assertEquals("Ninth", queue.dequeue());
  }
}
