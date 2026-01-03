//! Compare command - compare all algorithms

use fib_core::{iterative, matrix, recursive, closed_form};
use std::time::Instant;

pub fn run(n: u64, max_recursive: u64) {
    println!("╔═══════════════════════════════════════════════════════════════════╗");
    println!("║            🏁 Fibonacci Algorithm Comparison for n = {:6}       ║", n);
    println!("╠═══════════════════════════════════════════════════════════════════╣");
    println!("║ Algorithm           │ Result                        │ Time       ║");
    println!("╠═════════════════════╪═══════════════════════════════╪════════════╣");

    // Recursive (only for small n)
    if n <= max_recursive {
        let start = Instant::now();
        let result = recursive::fib_recursive(n);
        let elapsed = start.elapsed();
        println!("║ Recursive           │ {:29} │ {:10?} ║", result, elapsed);
    } else {
        println!("║ Recursive           │ (skipped - n > {})           │ N/A        ║", max_recursive);
    }

    // Recursive with memo
    let start = Instant::now();
    let result = recursive::fib_recursive_memo(n);
    let elapsed = start.elapsed();
    println!("║ Recursive+Memo      │ {:29} │ {:10?} ║", result, elapsed);

    // Iterative
    let start = Instant::now();
    let result = iterative::fib_iterative(n);
    let elapsed = start.elapsed();
    println!("║ Iterative           │ {:29} │ {:10?} ║", result, elapsed);

    // Iterative branchless
    let start = Instant::now();
    let result = iterative::fib_iterative_branchless(n);
    let elapsed = start.elapsed();
    println!("║ Iterative Branchless│ {:29} │ {:10?} ║", result, elapsed);

    // Matrix
    let start = Instant::now();
    let result = matrix::fib_matrix_fast(n);
    let elapsed = start.elapsed();
    println!("║ Matrix              │ {:29} │ {:10?} ║", result, elapsed);

    // Matrix doubling
    let start = Instant::now();
    let result = matrix::fib_doubling(n);
    let elapsed = start.elapsed();
    println!("║ Matrix Doubling     │ {:29} │ {:10?} ║", result, elapsed);

    // Binet (with accuracy warning)
    let start = Instant::now();
    let binet_result = closed_form::fib_binet_f64(n);
    let elapsed = start.elapsed();
    if n <= closed_form::MAX_ACCURATE_N {
        println!("║ Binet (f64)         │ {:29.0} │ {:10?} ║", binet_result, elapsed);
    } else {
        println!("║ Binet (f64) ⚠️       │ {:29.0} │ {:10?} ║", binet_result, elapsed);
    }

    println!("╚═════════════════════╧═══════════════════════════════╧════════════╝");

    if n > closed_form::MAX_ACCURATE_N {
        println!("\n⚠️  Note: Binet formula loses precision for n > {}", closed_form::MAX_ACCURATE_N);
    }
}
