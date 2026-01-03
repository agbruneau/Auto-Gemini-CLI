//! Sequence command - generate Fibonacci sequences

use fib_core::iterative::fib_iterative;

pub fn run(count: u64, start: u64) {
    println!(
        "📐 Fibonacci Sequence (F({}) to F({}))",
        start,
        start + count - 1
    );
    println!();

    let max_digits = format!("{}", fib_iterative(start + count - 1)).len();

    for i in 0..count {
        let n = start + i;
        let fib = fib_iterative(n);

        // Show the golden ratio approximation for n > 0
        let ratio_str = if n > 0 && i > 0 {
            let prev = fib_iterative(n - 1);
            if prev > 0 {
                format!("{:.10}", fib as f64 / prev as f64)
            } else {
                "∞".to_string()
            }
        } else {
            "-".to_string()
        };

        println!(
            "  F({:4}) = {:>width$}    φ ≈ {}",
            n,
            fib,
            ratio_str,
            width = max_digits
        );
    }

    println!();
    println!("φ (golden ratio) = 1.6180339887...");
}
