#include <iostream>
using namespace std;

int main() {
  // array of 6 characters
  char v[6] =  {'a', 'b', 'c', 'd', 'e', 'f'};
  // pointer to character
  char* p;
  char x;

  // A pointer variable can hold the address of an object of the appropriate type
  // p points to v's fourth element.
  p = &v[3];
  // *p is the object that p points to
  x = *p;
  cout << x;

}

// in an expression, unary * means contents of and prefix unary & means address of.
