fn main() {
    // let big: i64 = 3_000_000_000;
    // println!("{big}");

    //     let t = (1, "two", 3.0);
    //     let (a, b, c) = t;
    //     println!("{}", t.0);
    //
    //     let arr: [i32; 3] = [10, 20, 30];
    //     println!("{} {}", arr[0], arr.len());
    //

    // ex-1
    // 1.  Make three i32 variables: 7, 8, and 10.
    // 2.  Find the average and keep the decimals. Print it. The output must be 8.333333333333334.
    // 3.  Try let x = sum / 3.0; with an i32 sum. Read the compiler error. Then correct the code with as.
    // 4.  Make let x: u8 = 255;.
    // 5.  Print x.checked_add(1), x.wrapping_add(1), and x.saturating_add(1).
    //
    //     let (a, b, c) = (7, 8, 10);
    //     let sum: i32 = a + b + c;
    //     let avg: f64 = sum as f64 / 3.0;
    //
    //     println!("{}", avg)

    // let x: u8 = 255;
    // println!("{:?}", x.checked_add(1));
    // println!("{}", x.wrapping_add(1));
    // println!("{}", x.saturating_add(1));

    // ex-2
    // 1.  Write a function min_max.
    // 2.  It has one parameter: an array of 5 i32 values.
    // 3.  It returns the lowest and highest value as a tuple (i32, i32).
    // 4.  Call it from main with [4, -2, 9, 0, 7].
    // 5.  Use destructuring to get low and high.
    // 6.  Print both. The output must be low = -2, high = 9.
    //
    // let (low, high) = min_max([4, -2, 9, 0, 7]);
    // println!("low: {}, high: {}", low, high);

    // ex-3
    // 1.  Add your to_fahrenheit function from session 1.
    // 2.  Print one line for each value: its index, the Celsius value, and the Fahrenheit value.
    // 3.  Change the type to [f64; 5] and add a sixth value. Read the compiler error.
    // 4.  Try arr[5] with an array of 5 values. Does the error come at compile time or at run time?

    let arr: [f64; 5] = [-5.0, 0.0, 21.5, 37.0, 100.0];
    for (idx, c) in arr.iter().enumerate() {
        println!("{}. {}C = {}F", idx, c, to_fahrnenheit((*c)));
    }
}

fn to_fahrnenheit(c: f64) -> f64 {
    c * 9.0 / 5.0 + 32.0
}

fn min_max(arr: [i32; 5]) -> (i32, i32) {
    let (mut low, mut high) = (arr[0], arr[0]);

    for i in arr {
        if i > high {
            high = i;
        }
        if i < low {
            low = i;
        }
    }
    // return (low, high);
    // the last line without semicolon (;) is the return value
    (low, high)
}
