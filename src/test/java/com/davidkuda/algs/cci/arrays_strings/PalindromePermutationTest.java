package com.davidkuda.algs.cci.arrays_strings;

import org.junit.Assert;
import org.junit.Test;

public class PalindromePermutationTest {

  @Test
  public void testIsPermutation() {
    // given:
    PalindromePermutation p = new PalindromePermutation();

    // when:
    boolean b = p.isPalindromePermutation("David K David");

    // then:
    Assert.assertTrue(b);
  }

  @Test
  public void testIsNotPermutation() {
    // given:
    PalindromePermutation p = new PalindromePermutation();

    // when:
    boolean b = p.isPalindromePermutation("Dave K David");

    // then:
    Assert.assertFalse(b);
  }

  @Test
  public void furtherTests() {
    PalindromePermutation p = new PalindromePermutation();

    Assert.assertFalse(p.isPalindromePermutation("Dave K David"));
    Assert.assertTrue(p.isPalindromePermutation("Dave K vead"));
    Assert.assertTrue(p.isPalindromePermutation("this is sparta is this sparta?"));
  }

  
}
