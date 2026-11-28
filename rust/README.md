
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
