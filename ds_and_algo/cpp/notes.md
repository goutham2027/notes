c++ offers a variety of fundamental types.

bool:
char: 1 byte
int: 4 bytes
double:

sizeof operator - to get size of a type

double d1=2.3;
double d2 {2.3};
complex<double>z = 1;
complex<double>z2 {d1,d2};
complex<double> z3 = {1,2} // = is optional with {}

vector<int> v {1,2,3,4,5} // a vector of ints

A constant cannot be left uninitialized and a variable should only be
left uninitialized in extremely rare cases.

when defining a variable, we don't actually need to state its type
explicitly when it can be deduced from the initializer.
auto i = 123;
we use auto where we don't have a specific reason to mention the type
explicitly.

c++ supports 2 notions of immutability:
 const: it is never modified
 constexpr: to be evaluated at compile time

c++ offers a simpler for-statement, called a range-for statement, for
loops that traverse a sequnce in the simplest way
int v[] = {0,1,2,3,4,5,6,7,8,9}
for(auto x: v) // for each x in v
  cout << x << "\n";
