void main() {
  print('Hello, Dart!');
  // null safety
  int? i = 0;
  print(i);

  // late variables

  // a final variable can only be set once
  // final variable is determined at the runtime
  // const variables are implicitly final
  // constants are compile time constants.
  //  use final for values that can be determined
  // at runtime and need to be assigned once.
  //  Use const for values that are known at compile-time
  // and are constant expressions

  var image = {
    'tags': ['saturn'],
    'url': '//path/to/saturn.jpg'
  };
// print(image["tags"]);

  image.forEach((key, value) {
    print('$key: $value');
  });
}
