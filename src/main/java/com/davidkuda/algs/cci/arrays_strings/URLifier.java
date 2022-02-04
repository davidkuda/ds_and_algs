package com.davidkuda.algs.cci.arrays_strings;

public class URLifier {

  /**
   * Iterate over a String and replace all spaces with "%20".
   * @param s The String as input
   * @param len The length of the eventual CharArray
   * @return String
   */
  public String urlify(String s) {

    int numOfSpaces = countNumOfChar(s, ' ');
    int newLen = s.length() + numOfSpaces * 2;

    char[] charArr = new char[newLen];

    int counter = 0;
    for (char c : s.toCharArray()) {

      if (c == ' ') {
        charArr[counter] = '%';
        counter++;
        charArr[counter] = '2';
        counter++;
        charArr[counter] = '0';
        counter++;
        continue;
      }

      charArr[counter] = c;
      counter++;
    }

    String sTransformed = new String(charArr);
    System.out.println("sTransformed: " + sTransformed);
    return sTransformed;
  }

  public int countNumOfChar(String s, char targetChar) {
    int counter = 0;
    for (char c : s.toCharArray()) {
      if (c == targetChar)  {
        counter++;
      }
    }
    return counter;
  }
}
