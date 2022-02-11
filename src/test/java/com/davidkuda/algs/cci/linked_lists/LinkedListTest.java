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
    l.insertAtEnd("Third");
    l.insertAtEnd("Fourth");
    l.insertAtEnd("Fifth");
    l.insertAtEnd("First");
    l.insertAtEnd("Third");

    l.removeDups();

    Assert.assertEquals("First", l.first.item);
    Assert.assertEquals("Second", l.first.next.item);
    Assert.assertEquals("Third", l.first.next.next.item);
    Assert.assertEquals("Fourth", l.first.next.next.next.item);
    Assert.assertEquals("Fifth", l.first.next.next.next.next.item);
  }
  
}
