#include <iostream>
using namespace std;

double square(double x) {
  return x*x;
}

void print_square(double x) {
  cout << "the square of " << x << " is " << square(x) << "\n";
}

int main() {
  /*
   *cout << "hello world";
   */
  print_square(2.1233);
  return 0;
};
