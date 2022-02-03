package com.davidkuda.algs.cci.arrays_strings;

import org.junit.Assert;
import org.junit.Test;

public class PermutationTest {

  Permutation pObj = new Permutation();

  @Test
  public void testDifferentLength() {

    String s1 = "Hello";
    String s2 = "World!";
    boolean p = pObj.isPermutation(s1, s2);
    Assert.assertFalse(p);
  }

  @Test
  public void testSameWord() {
    Assert.assertTrue(pObj.isPermutation("Stratocaster", "Stratocaster"));
  }

  @Test
  public void testPermutation() {
    Assert.assertTrue(pObj.isPermutation("Berlin", "linBer"));
  }
}
