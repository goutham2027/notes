// pluralsight course
https://app.pluralsight.com/library/courses/go-create-test-applications/table-of-contents

### Intro to testing

- testing and profiling tools in the standard library
- complete API for controlling test exection
- little support for assertions in the standard library

We can write 3 types of tests using Go tooling

- Test (assertion style tests)
  - Unit
  - Integration
  - End to End
- Benchmark tests
  - Performance
- Example tests
  - Documentation

Running tests `go test`

#### Testing related packages

standard library packages

- testing
- testing/quick
  - to simplify blackbox testing
- testing/iotest

  - contains several readers and writers

- net/http/httptest
  - simulate requests and responses and write assertions

3rd party packages

- Testify
  - github.com/stretchrcom/testify
  - assertion framework
- Ginkgo
  - BDD
  - github.com/onsi/ginkgo
- GoConvey
  - viewing test results
  - goconvey.co
- httpexpect
  - end to end testing for webservices
  - testing rest apis
  - github.com/gavv/httpexpect
- gomock
  - mocking
  - code.google.com/p/gomock
- gosqlmock
  - github.com/DATA-DOG/go-sqlmock
  - in-memory sql mocking

### Creating and Running tests

#### Naming conventions

- add `_test` to filenames
- `_test` files are excluded when binaries are created
- prefix tests with `Test` eg: `TestAddition`
- Accept one param: `*testing.T`. eg: `TestAddition(t *testing.T)`
- Same package for whitebox tests
- Add `_test` suffix to package for blackbox tests

immediate failure

- exit the tests as soon as a test is exited with a failure

```golang
t.FailNow()
t.Fatal(args ...interface{})
t.Fatalf(format string, args ...interface{})
```

Non-immediate failure

- do not exit tests when a test is failed

```golang
t.Fail()
t.Error(args ...interface{})
t.Errorf(format string, args ...interface{})
```

### Running tests

```bash
go test # run all tests in current directory
go test {pkg1} {pkg2} # test specified packages
go test ./... # Run tests in current package and descendants
go test -v # Generate verbose output
go test -run {regexp} # Run only tests matching {regexp}

```

#### Test coverage

```bash
go test -cover

go test -coverprofile cover.out

go tool cover -func cover.out

go tool cover -html cover.out

go test -coverprofile count.out -covermode count # relative scale
```

#### Table-driven tests

To test different inputs in a single test

#### Other useful functions

- Log and Logf
- Helper
- Skip, Skipf, SkipNow
- Run
- Parallel

### Benchmarking and Profiling

- Use `Benchmark` prefix
- Accept one parameter `*testing.B`

```golang
func BenchmarkFoo(b *testing.B) {}
```

- `b.N` - number of times to run the test
- `b.StartTimer`, `b.StopTimer`, `b.ResetTimer`
- `b.RunParallel`

#### Running benchmarking tests

```golang
go test -bench
go test -bench -benchtime 10s
go test -v -bench .
```

#### Profiling

```bash
go test -benchmem # reports memory allocation stats for benchmarks
go test -trace {trace.out} # record execution trace to {trace.out} for analysis
# generate profile of requested type
# block
# cover
# cpu
# mem
# mutex
go test -{type}profile {file}

go tool pprof profile.out
```
