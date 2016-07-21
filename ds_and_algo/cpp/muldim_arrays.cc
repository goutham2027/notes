#include <iostream>
using namespace std;

#define WIDTH 5
#define HEIGHT 3

//int jimmy [HEIGHT][WIDTH];
int jimmy [HEIGHT * WIDTH];
int n,m;

int main() {
  // multi dim array
  /*
  for(n=0; n<HEIGHT; n++) {
    for(m=0; m<WIDTH; m++) {
      jimmy[n][m] = (n+1) * (m+1);
    }
  }

  for(n=0; n<HEIGHT; n++) {
    for(m=0; m<WIDTH; m++) {
      cout << jimmy[n][m] << " ";
    }
    cout << "\n";
  }
  */

  // single dimension arrays
  for(n=0; n<HEIGHT; n++) {
    for(m=0; m<WIDTH; m++) {
      jimmy[n*WIDTH+m] = (n+1) * (m+1);
    }
  }


  for(n=0; n<HEIGHT; n++) {
    for(m=0; m<WIDTH; m++) {
      cout << jimmy[n*WIDTH+m] << " ";
    }
      cout << "\n";
  }
  cout << "The end!!";
}
