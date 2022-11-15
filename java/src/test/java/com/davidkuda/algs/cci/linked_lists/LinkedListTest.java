package com.davidkuda.algs.cci.linked_lists;

import org.junit.Assert;
import org.junit.Test;

public class LinkedListTest {

  @Test
  public void testGetSize() {
    LinkedList<Integer> l = new LinkedList<Integer>();    

    Assert.assertEquals(0, l.getSize());

    l.insertAtBeginning(108);

    Assert.assertEquals(1, l.getSize());

    l.insertAtBeginning(108);
    Assert.assertEquals(2, l.getSize());

    l.insertAtBeginning(108);
    Assert.assertEquals(3, l.getSize());
    

    l.removeFromBeginning();
    Assert.assertEquals(2, l.getSize());

    l.removeFromBeginning();
    Assert.assertEquals(1, l.getSize());

    l.removeFromBeginning();
    Assert.assertEquals(0, l.getSize());
  }

  @Test
  public void testRemoveDups() {

    // given:
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

    // when:
    l.removeDups();

    // then:
    Assert.assertEquals("First", l.first.item);
    Assert.assertEquals("Second", l.first.next.item);
    Assert.assertEquals("Third", l.first.next.next.item);
    Assert.assertEquals("Fourth", l.first.next.next.next.item);
    Assert.assertEquals("Fifth", l.first.next.next.next.next.item);
  }

  @Test
  public void testIteration() {
    // given:
    LinkedList<String> l = new LinkedList<String>();

    // when:
    l.insertAtBeginning("First");
    l.insertAtEnd("Second");
    l.insertAtEnd("Third");
    
    // then:
    for (String i : l) {
      System.out.println(i);
    }
  }
  
}
