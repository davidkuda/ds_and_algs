package com.davidkuda.algs4.fundamentals;

import org.junit.Test;
import org.junit.Assert;

public class StackTets {

  @Test
  public void testLinkedStack() {

    // Given:
    LinkedStackOfStrings linkedStack = new LinkedStackOfStrings();

    // When:
    linkedStack.push("First");
    linkedStack.push("Second");
    linkedStack.push("Third");

    // Then:
    Assert.assertEquals("Third", linkedStack.pop());
    Assert.assertEquals("Second", linkedStack.pop());
    Assert.assertEquals("First", linkedStack.pop());

    Assert.assertTrue(linkedStack.isEmpty());
  }

  @Test
  public void testFixedCapacityStack() {

    // Given:
    FixedCapacityStackOfStrings fixCapaStack = new FixedCapacityStackOfStrings(3);

    // When:
    fixCapaStack.push("First");
    fixCapaStack.push("Second");
    fixCapaStack.push("Third");

    // Then:
    Assert.assertEquals("Third", fixCapaStack.pop());
    Assert.assertEquals("Second", fixCapaStack.pop());
    Assert.assertEquals("First", fixCapaStack.pop());

    Assert.assertTrue(fixCapaStack.isEmpty());
  }
}
