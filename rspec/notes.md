There are several definitions of objects that are not real. The general
term is test double. This term encompasses: dummy, fake, stub, mock.
* Dummy objects are passed around but never actually used. Usually they
  are just used to fill parameter lists.
* Fake objects actually have working implementations, but usually take
  some shortcut which makes them not suitable for production.
* Stubs provide canned answers to calls made during the test, usually
  not responding at all to anything outside what's programmed for the
  test.
* Mocks are objects pre-programmed with expectations which form a
  specification of the calls they are expected to receive.

  Mocks vs Stubs = Behavioral testing vs state testing

  stub - for replacing a method with code that returns a specified
  result.
  mock - a stub with an assertion that the method gets called.
Ref: http://stackoverflow.com/a/17810004
