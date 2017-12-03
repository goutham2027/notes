## unittesting with python - pluralsight course

### module - 1 unittesting with python
System Under Test
a unittest checks the behavior of an element of code
  - method/function
  - module/class

for each system under test there will be more test cases with
different behaviors. each test case should run independently.

a test runner is a program that executes test cases and reports the
results. In unittest module the command line test runner is builtin.
`python -m unittest`
It will automatically dicover tests.

a test suite is number of test cases executed by test runner.

test fixture: is a piece of code that can construct and configure
the system under test to get it ready for tested and then cleanup
afterwards.

it's not a unit test if it uses:
  - the file system
  - a db
  - the network

unittests run in memory.

Exercise - Phone Numbers
- Given a list of names and phone numberss, make a Phonebook that allows
you to look up numbers by name.

- Determine if a given Phonebook is consistent.
  - in a consistent phone list no number is a prefix of another
    eg:  Bob 91125426, alice 97625992, emergency 911
    Bob and emergency are inconsistent.


running unittests

`python -m unittest -v`

`python -m unittest -q
<test_phonebook.PhonebookTest.test_create_phonebook>`


assertEqual - convention
1st arg - expected value
2nd arg - value we get


test case design:
  Name
  it has arrange, act, assert
  arrange - entries into phonebook
  act - call inconsistent
  assert - check the output

  arrange - setup the object to be tested and collaborators
  act - exercise functionality on the object
  assert - make claims about the object and its collaborators


failure in one test case will not execute the rest in the series.


when to have more than one assert?
```
def test_phonebook_adds_names_and_numbers(self):
  self.phonebook.add("sue", "12345")
  self.assertIn("sue", self.phonebook.get_names())
  self.assertIn("12345", self.phonebook.get_phone_numbers())
```
the only reason this test will fail is if there is an error in add
implementation.


### module - 2 Why and When should you write unit tests?

documenting the units
- executable specification

- helps with the design
  * decompose the problem into units that are independently testable
    - loosely coupling
  * design the interface separately from doing the
    implementation.


- detecting regressions
 regression - something worked before and doesn't any more.
  * a unit test should fail and point out which unit failed and why.


- limitations of unit testing
  * testing can't find all the errors
  * unit testing won't find integration errors
  * cannot test non functional requirements eg: performance, security


test last
test first
test driven


### module - 3 Using Pytest for Unit Testing in Python
unittest alternatives = python nose, pytest

`python -m pytest`
to skip a test
pytest.skip("wip") works both as a decorator and as a function.

pytest test fixtures - a test fixture is a code to construct and
configure the system under test to get it ready to be tested and to
cleanup afterwards

pytest offers a different way to share test fixtures between test cases.

philosophy is we create test fixture functions that can provide
resources to test cases and test cases will request a resource that they
need. Its a kind of dependency injection.

pytest teardown - add finalizer

python -m pytest --fixtures # to list all fixtures
built-in test fixtures
tmpdir

### module - 4 Testable documentation with Doctest

- checking examples in docstrings
- regression testing
- tutorial documentation

Docstrings can get out of date. Doctests helps to keep them updated.

Yatzy
- Roll 5 dice (and re-roll some)
- choose a category to score the roll in
- each category is only used once
- the final score is the sum of the score in each category.

dice 1 - 1
dice 2 - 2
dice 3 - 2
dice 4 - 3
dice 5 - 3

2 pairs: 2 2s and 2 3s
  score = 2 + 2 + 3 + 3 = 10
3s
  score 3 + 3 = 6
small straight
  score 0


to run doctests
`python -m doctest`

`python -m pytest --doctest-modules`

handling output that changes:
  dictionaries
  floating point values
  object id


ellipsis - ... used for unwanted strings in doctest

for tracebacks it's straightforward to use it but for functions
```
>>> call the function #doctest: +ELLIPSIS
[(8, function full_house ...)]
```
for pytest this ELLIPSIS module is there by default, we can remove.

downside of ellipsis is it will match more than what we need.


docttest directives
directives control how doctest matches output.


when to put doctests in a File.
- approval testing

it's important to understand what role is your doctest playing.

if doctests are there to document few examples - then put them in
doctstring

if doctests for regression protection then put them in a separate
file.

### module - 5 Test Doubles: Mocks, Fakes and Stubs
Test Doubles - Mock objects

What is a Test Double?

- analogy - stunt Double: impersonate the actual star

- class under test doesn't know it isn't talking to the real object

- allow you to control what happens to your class under test


test double - from x unit patterns

  dummy object
  test stub
  test spy
  mock object
  fake object


  order of impportance
  * stub
  * fake
  * mock
  * test spy
  * dummy object

  stub: a kind of test double that stands in for collaborating object
  that is difficult to use in a test case.

  a stub is any test double that has a very simple implementation with
  no logic or behavior.

  eg: Racing Car Example

  alarm when pressure is too high or too low.
  Alarm
   check

  Sensor
    sample_pressure

```python
  class Alarm():
    def __init__(self, sensor=None):
      self._sensor = sensor or Sensor()
      ...
      ...

    def check(self):
      pressure = self._sensor.sample_pressure()



  Test
  def test_check_too_low_pressure_sounds_alarm():
    alarm = Alarm(sensor=TestSensor(15))
    alarm.check()

   def test_2():
    alarm = Alarm(sensor=TestSensor(13))
    alarm.check()



  def TestSensor:
    def __init__(self, pressure):
      self._pressure = pressure

    def sample_pressure(selfj):
      return self._pressure
```

unittest module comes builtin with mock module for stubs

```python
from unittest.mock import Mock
from sensor import Sensor

def test_check_pressure():
  test_sensor = Mock(Sensor)
  test_sensor.sample_pressure = 18
  alarm = Alarm(test_sensor)
  alarm.check()
  assert ...
```

mock is different from stub, but most libraries don't have a clear
distinguish between them.

-------------------

Fake: its like a stub where we replace the entire subsytem or component where it's
inconvenient to use it as is.

unlike stub, fake has some logic/implementation that works.

fake is a test double that contains both logic and functionality unlike
stub but which we use just for the purposes of the test, in production
we use different implementation.

```python
import io

class FakeFileWrapper:
  def __init__(self, text):
    self.text = text

  def open(self):
    return io.StringIO(self.text)


```

common things to replace with Fakes
- File
- Database
- WebServer

----------
Mock: mostly people use mocks, stubs and test doubles interchangeably.
But there is an important difference between stub and mock.

mocks are like stubs which returns hard coded value with no their own
behavior. Unlike stubs, mocks make assertions on what happened in the
test case. Mocks can cause a test fail stub wont.

3 kinds of assert
- check the return value or an exception
- check a state change (use a public API)
- check a method call - (use a mock or spy)

with mock we can check if a particular method call was made to a
particular object with particular arguments and fail the test if the
method was not called as specified.


example

  MyService
    handle(request, token)

  SSORegistry
    register(id): token
    is_valie(token)
    unregister(token)

  MyServiceTest
    test_valid_token
    test_invalid_token


-----------
Spy: used in similar way of mock. listen on the coverstation between the
class we are testing and the collaborating class.


using unittest to create a mock
```python
def test():
  registry = Mock(SingleSignOnRegistry)
  def is_valid(token):
    if not token == valid_token:
      raise Exception("Got the wrong token")
     return true
  registry.is_valid = Mock(side_effect=is_valid)
  my_service = MyService(registry)

  response = my_service.handle_request("do_stuff", token=valid_token)
  registry.is_valid.assert_called_with(valid_token)
```

---------
dummy object: basic kind of test doubles, empty or None
you are forced to provide argument eventhough it's not required. instead
make the argument has default value as None or empty object.

- a stub returns a hard coded answer to any query
- a fake is a real implementation, yet unsuitable for production
- a mock is as a stub, and additionally verifies interactions
- a test spy lets you query afterwards to find out what happened
- a dummy is for when the interface requires an argument.
