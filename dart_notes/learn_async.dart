Future<String> fetchUserData() {
  return Future.delayed(Duration(seconds: 3), () {
    return "User data retrieived";
  });
}

// blocking
// void main() async {
//   print("Requesting user data");
//   String data =
//       await fetchUserData(); // pauses the main function until the future is complete
//   print("I am thinking something else");
//   print(data);
// }

// non blocking
void main() {
  print("Requesting user data");
  // handle the Future using then()
  fetchUserData().then((result) {
    print(result);
  });
  print("I am thinking something else");
}
