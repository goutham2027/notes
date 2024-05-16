import 'dart:async';

void main() {
  // create a streamcontroller
  final StreamController<int> controller = StreamController<int>();

  // subscribe to the controller's stream
  controller.stream
      .where((number) => number % 2 == 0) // Filter to only event numbers
      .listen((number) => print("Received: $number"),
          onError: (error) => print("Error: $error"),
          onDone: () => print("Stream completed"));

  // adding data to the stream
  List<int> numbers = [1, 2, 3, 4, 5, 6];
  for (var n in numbers) {
    controller.sink.add(n);
  }
  controller.close();
}
