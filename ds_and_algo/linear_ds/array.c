#include <stdio.h>

void print_an_int_arr(int arr[]) {
  for(int i=0; i<10; i++) {
    if arr[i] == "\0" {
      break;
    }
    printf("%d \n", arr[i]);
  }

}

int main() {
  // static arrays
  int a[10];
  int b[5];

  /* inserting elements into an array */
  for(int i=0; i<10; i++) {
    a[i] = i;
  }

  for(int i=0; i<5; i++) {
    b[i] = i;
  }

  // passing array - it passes the array's first element address.
  print_an_int_arr(a);
  print_an_int_arr(b);
}


