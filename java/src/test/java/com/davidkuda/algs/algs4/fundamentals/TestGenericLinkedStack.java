package com.davidkuda.algs.algs4.fundamentals;

import org.junit.Assert;
import org.junit.Test;

public class TestGenericLinkedStack {

  @Test
  public void testGenericLinkedStack1() {
    // given:
    GenericLinkedStack<String> s = new GenericLinkedStack<String>();

    // when:
    s.push("first");
    s.push("second");
    s.push("third");

    // then:
    Assert.assertEquals("third", s.pop());
    Assert.assertEquals("second", s.pop());
    Assert.assertEquals("first", s.pop());
  }

  @Test
  public void testGenericLinkedStack2() {
    GenericLinkedStack<String> s = new GenericLinkedStack<String>();

    s.push("first");
    s.push("second");

    Assert.assertEquals("second", s.pop());

    s.push("third");

    Assert.assertEquals("third", s.pop());
    Assert.assertEquals("first", s.pop());
  }

  @Test
  public void testGenericLinkedStack3() {
    GenericLinkedStack<String> s = new GenericLinkedStack<String>();

    String[] words = {
      "Kicking", "around", "on", "a", "piece", "of", "ground", "in", "your", "hometown"
    };

    for (String word : words) {
      s.push(word);
    };

    int counter = 0;
    while (!s.isEmpty()) {
      String expectedWord = words[words.length - 1 - counter];
      String actualWord = s.pop();
      Assert.assertEquals(expectedWord, actualWord);
      counter++;
    };
  }
}
