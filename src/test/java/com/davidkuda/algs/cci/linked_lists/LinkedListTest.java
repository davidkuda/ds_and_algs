package com.davidkuda.algs.cci.linked_lists;

import org.junit.Assert;
import org.junit.Test;

public class LinkedListTest {

  @Test
  public void testRemoveDups() {
    LinkedList<String> l = new LinkedList<String>();

    l.insertAtBeginning("First");
    l.insertAtEnd("Second");
    l.insertAtEnd("Second");
    l.insertAtEnd("Third");

    Assert.assertEquals(4, l.getSize());

    l.removeDups();
    Assert.assertEquals(3, l.getSize());
  }
  
}
