package com.davidkuda.algs.algs4.fundamentals;

import java.util.Iterator;

import org.junit.Assert;
import org.junit.Test;

public class TestIteratorLinkedStack {

  @Test
  public void testIteratorLinkedStack1() {
    IteratorLinkedStack<String> s = new IteratorLinkedStack<String>();

    String[] words = {"first", "second", "third"};
    
    for (String word : words) {
      s.push(word);
    }

    int counter = 0;

    Iterator<String> it = s.iterator();

    while (it.hasNext()) {
      String expectedWord = words[words.length - 1 - counter];
      String actualWord = it.next();
      Assert.assertEquals(expectedWord, actualWord);
      counter++;
    }
  }

  @Test
  public void testIteratorLinkedStack2() {
    IteratorLinkedStack<String> s = new IteratorLinkedStack<String>();

    String[] words = {"first", "second", "third"};
    
    for (String word : words) {
      s.push(word);
    }

    int counter = 0;
    for (String word : s) {
      String expectedWord = words[words.length - 1 - counter++];
      Assert.assertEquals(expectedWord, word);
    }
  }
}
