#include <iostream>
using namespace std;

void printArray(int[], int);

int main() {
  /*
   *int v[] = {0,1,2,3,4,5,6,7,8,9};
   */
  // place a copy in x and print it;
  /*
   *for(int x: v) {
   *  cout << x << "\n";
   *}
   */

  // if we didn't want to copy the values from v into the variable x, but rather
  // just have x refer to element.
  /*
   *int v[] = {10,11,12,13,14,15,16,17,18,19};
   *for(int& x: v) {
   *  ++x;
   *  cout << x << "\n";
   *}
   */

  int v[] = {1,2,3,4,5,6};
  printArray(v, 6);
  return 0;
}


// in a declaration sytax, unary suffix * means "reference to".
// A reference is similar to a pointer, except that we don't need to use a
// prefix * to access the value referred to by the reference.


void printArray(int args[], int length) {
  for(int i=0; i<length; i++) {
    cout << args[i] << "\n";
  }
}
