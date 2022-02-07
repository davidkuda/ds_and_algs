package com.davidkuda.algs.cci.arrays_strings;

import org.junit.Assert;
import org.junit.Test;

public class PalindromePermutationTest {

  @Test
  public void testIsPermutation() {
    // given:
    PalindromePermutation p = new PalindromePermutation();

    // when:
    boolean b = p.isPermutation("David K David");

    // then:
    Assert.assertTrue(b);
  }

  @Test
  public void testIsNotPermutation() {
    // given:
    PalindromePermutation p = new PalindromePermutation();

    // when:
    boolean b = p.isPermutation("Dave K David");

    // then:
    Assert.assertFalse(b);
  }

  
}
