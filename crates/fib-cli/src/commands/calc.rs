//! Calculate command implementation

use fib_core::FibMethod;
use std::time::Instant;

pub fn run(n: u64, method: &str, show_time: bool) {
    let method: FibMethod = match method.parse() {
        Ok(m) => m,
        Err(e) => {
            eprintln!("❌ Error: {}", e);
            eprintln!("Available methods: recursive, recursive_memo, iterative, matrix, binet");
            std::process::exit(1);
        }
    };

    // Warn about slow recursive for large n
    if matches!(method, FibMethod::Recursive) && n > 35 {
        eprintln!("⚠️  Warning: Recursive method is extremely slow for n > 35");
        eprintln!("    Consider using --method iterative or --method matrix");
    }

    let start = Instant::now();
    let result = method.calculate(n);
    let elapsed = start.elapsed();

    println!("┌─────────────────────────────────────────────────┐");
    println!("│ 🔢 Fibonacci Calculation                        │");
    println!("├─────────────────────────────────────────────────┤");
    println!("│ Method:     {:20}              │", method.name());
    println!("│ n:          {:20}              │", n);
    println!("│ F(n):       {}", result);

    if show_time {
        println!("├─────────────────────────────────────────────────┤");
        println!("│ ⏱️  Time: {:?}", elapsed);
    }

    println!("└─────────────────────────────────────────────────┘");
}
