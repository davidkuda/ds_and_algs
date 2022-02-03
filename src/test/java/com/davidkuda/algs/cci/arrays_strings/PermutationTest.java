package com.davidkuda.algs.cci.arrays_strings;

import org.junit.Assert;
import org.junit.Test;

public class PermutationTest {

  @Test
  public void testDifferentLength() {

    String s1 = "Hello";
    String s2 = "World!";
    boolean p = Permutation.isPermutation(s1, s2);
    Assert.assertFalse(p);
  }

  @Test
  public void testSameWord() {
    Assert.assertTrue(Permutation.isPermutation("Stratocaster", "Stratocaster"));
  }

  @Test
  public void testPermutation() {
    Assert.assertTrue(Permutation.isPermutation("Berlin", "linBer"));
  }
}
