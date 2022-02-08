package com.davidkuda.algs.cci.arrays_strings;

public class PalindromePermutation {

  boolean[] palindromometer = initPalindromometer();

  // Set all values in palindrometer to true.
  private boolean[] initPalindromometer() {
    boolean[] palindromometer = new boolean[25];
    for (int i = 0; i < palindromometer.length; i++) {
      palindromometer[i] = true;
    }
    return palindromometer;
  }

  /**
   * Check if string is a palindrome permutation.
   *
   * <p>A palindrome is a string that reads same backwards and forwards.
   *
   * @param t String to be checked
   * @return true if t is a palindrome permutation
   */
  public boolean isPalindromePermutation(String s) {
    char[] charArr = s.toCharArray();


    for (char c : charArr) {
      if (Character.isLetter(c)) {
        char lc = Character.toLowerCase(c);
        // ASCII lowercase letters from a - z -> 97 to 122
        int asciiValue = (int) lc;
        switchBool(asciiValue - 97);
      }
    }

  return checkPalindrometer();
  }

  private void switchBool(int index) {
    if (palindromometer[index] == true) {
      palindromometer[index] = false;
    } else {
      palindromometer[index] = true;
    }
  }

  private boolean checkPalindrometer() {
    int falseCount = 0;
    for (boolean b : palindromometer) {
      if (b == false) {
        falseCount++;
      }
    }
    // One character may be in the middle and the string still be a palindrome.
    if (falseCount > 1) {
      return false;
    } else  {
      return true;
    }
  }

}
