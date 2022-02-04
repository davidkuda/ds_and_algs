package com.davidkuda.algs.cci.arrays_strings;

import org.junit.Assert;
import org.junit.Test;

public class URLifierTest {
  
  @Test
  public void testUrlifier() {
    URLifier u = new URLifier();
    String actual = u.urlify("This Love Has Taken its Toll On Me", 48);
    String expected = "This%20Love%20Has%20Taken%20its%20Toll%20On%20Me";
    Assert.assertEquals(expected, actual);
  }
  
  @Test
  public void testUrlifierSimple() {
    URLifier u = new URLifier();
    String actual = u.urlify("Hello World", 13);
    String expected = "Hello%20World";
    Assert.assertEquals(expected, actual);
  }
  
}
