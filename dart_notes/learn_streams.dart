import 'dart:async';

// number stream emits an integer every second
// we listen to this stream and print each number as it arrives
// onDone callback is executed

// streams are useful in Flutter for handling things like
// user input, timers
void main() {
  // creating stream of ints
  Stream<int> numberStream =
      Stream.periodic(Duration(seconds: 1), (x) => x).take(5);

  print("Stream started");

  // Listening to the stream
  numberStream.listen((number) {
    print("Number from stream: $number");
  }, onDone: () {
    print("Stream completed");
  });

  print("after listening");
}
