## Async coding with Dart

### 1. Isolates and Event loops

Dart is single threaded
Isolate: isolated space with mem
can spawn new isolates
Eventloop

### 2. Futures

Uncompleted (closed)
Completed with data
Completed with error

Think futures as a box. Box is closed, when a box is open it can either a value or an error

```dart
// network commn creates a future
myFuture = http.get('http://example.com');


// example
import `dart:async`;

void main() {
  final myFuture = Future((), {
    print("Creating the future");
    return 12;
  });
  print("Done with main().");
}
```

### 3. Streams

- reactive programming
- Each future represents by a single value, error
- Streams can give 0 or more values and errors overtime
  sync -> int -> Iterator<int>
  async -> Futreu<int> -> Stream<int>

### 4. aysnc/await
 