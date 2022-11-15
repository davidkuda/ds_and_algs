package com.davidkuda.algs.cci.arrays_strings;

import java.util.HashMap;

/** Take two strings and check if one is a permutation of the other. */
public class Permutation {

  public boolean isPermutation(String s1, String s2) {
    if (s1.length() != s2.length()) {
      return false;
    }

    HashMap<Character, Integer> s1Chars = countChars(s1);
    HashMap<Character, Integer> s2Chars = countChars(s2);

    if (s1Chars.equals(s2Chars)) {
      return true;
    } else {
      return false;
    }
  }

  private HashMap<Character, Integer> countChars(String s) {

    HashMap<Character, Integer> sChars = new HashMap<>();

    for (int i = 0; i < s.length(); i++) {
      Character curChar = s.charAt(i);

      if (sChars.get(curChar) != null) {
        int curVal = sChars.get(curChar);
        sChars.put(curChar, curVal++);
      } else {
        sChars.put(curChar, 1);
      }
    }

    return sChars;
  }
}
