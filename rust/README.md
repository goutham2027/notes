
`rustup` installs and updates rust

`cargo` is the build tool and package manager

```bash
rustc --version
cargo --version
```

Add `. "$HOME/.cargo/env"` to `.bashrc` or `.zshrc`. To restart zsh, use `zsh -l`

starting a new project
`cargo new <project_name>`

`cargo.toml` hold the project settings.

main is the starting point as in `golang`
`src/main.rs`

Compiles and runs the project
`cargo run`

`cargo build` only compiles

```rust
fn main() {
    println!("Hello World");
}
```

println! -> is a macro not a function

what is a macro?
- The `!` after the name show that is a macro
- A macro is code that writes more code. The compiler expands the macro into normal Rust code before it compiles the program.

### Variables
- In rust variables cannot change by default
```rust
let count = 5;
count = count + 1; // gives compiler error

let mut count = 5;
count = count + 1;
```

Shadowing
- Declaring variables with the same name
- mut changes the value, type stays the same
- shadowing makes a new variable, type can change
```rust
let age = "age is 42";
println!("{age}");
let age: i32 = 24;
// let age = "53";
// unwrap stops the program if the text is not a number
// let age: i32 = age.parse().unwrap()
```

### Functions
```rust
fn functionname(c: f64) -> f64 {}
```

### Conditionals
```rust
// 5 is not included
for i in 1..5 {}

// 5 is included
for i in 1..=5 {}

let names = vec!["a", "b", "c"];
for name in &names {}

for (i, name) in names.iter().enumerate() {}
```

### Types
- Rust checks every type at the compile time.

- signed integer - i8, i32, i64 - i32 is the default
- unsigned integer - u8, u32, u64 - no negative values
- float - f32, f64 - f64 is the default
- boolean - true, false, type bool
- character 'a' - type char. It is 4 bytes and holds one unicode character
- checked_add returns Option: None on overflow.
- wrapping_add wraps to 0.
- saturating_add stays at 255.

```
// _ makes numbers easy to read
let big: i64 = 3_000_000_000;
println!(i64);

7/2.0 is a compile error
```

### Tuples
- holds a fixed number of values, the types can differ
```rust
let t = (1, "two", 3.0);
let (a, b, c) = t
    println!("{}", t.0)
```

### Arrays
Arrays hold a fixed number of values. The type is same for all values.
The length is part of the type
```rust
let arr: [i32; 3] = [10, 20, 30];
println!("{} {}", arr[0], arr.len());
```

- use `as` for type conversion `avg as f64`
