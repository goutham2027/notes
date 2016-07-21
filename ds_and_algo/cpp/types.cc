#include <iostream>
using namespace std;

int main() {
  int sum = 0;
  for(int i=0; i<1000; i++) {
    sum += i;
    /*
     *cout << i << "\n";
     */
  }
  cout << "sum of 1000 nums is " << sum;
}
