fn main() {
    // let mut count = 5;
    // count = count + 1;

    // let age = "age is 42";
    // println!("{age}");
    // let mut age = 24;
    // age += 1;
    // let age = "53";
    // let age: i32 = age.parse().unwrap();
    // println!("{age}")
    //
    // let f = fahrnenheit(100.0);
    // println!("{f}");

    // printnumbers(15);
}

fn fahrnenheit(c: f64) -> f64 {
    c * 9.0 / 5.0 + 32.0
}

// rust has only for .. in syntax
// for i in 1..5 {}
// for i in 1..=5 {}
// let names = vec!["a", "b", "c"]; for name in &names {}
// for (i, name) in names.iter().enumerate() {}
fn printnumbers(n: i32) {
    // for i in 1..n {
    //     println!("{i}")
    // }

    for i in 1..=n {
        if i % 15 == 0 {
            println!("FizzBuzz");
        } else if i % 3 == 0 {
            println!("Fizz");
        } else if i % 5 == 0 {
            println!("Buzz")
        } else {
            println!("{i}")
        }
    }
}
