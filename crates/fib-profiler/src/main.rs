//! Fibonacci Profiler
//!
//! Profiling and memory analysis tools for Fibonacci benchmarks

use std::time::{Duration, Instant};
use fib_core::{iterative, matrix};

fn main() {
    println!("🔬 Fibonacci Performance Profiler");
    println!("==================================");
    println!();

    // Profile different methods
    profile_iterative();
    profile_matrix();
    profile_memory_usage();
    profile_scaling();
}

fn profile_iterative() {
    println!("📊 Iterative Method Profile");
    println!("─────────────────────────────");

    let test_values = [10, 100, 1000, 10000, 100000];

    for n in test_values {
        let iterations = 1000;
        let mut total_time = Duration::ZERO;

        for _ in 0..iterations {
            let start = Instant::now();
            let _ = iterative::fib_iterative(n);
            total_time += start.elapsed();
        }

        let avg_time = total_time / iterations;
        println!("  n = {:6}: avg {:?}", n, avg_time);
    }
    println!();
}

fn profile_matrix() {
    println!("📊 Matrix Method Profile");
    println!("─────────────────────────");

    let test_values = [10, 100, 1000, 10000, 100000];

    for n in test_values {
        let iterations = 1000;
        let mut total_time = Duration::ZERO;

        for _ in 0..iterations {
            let start = Instant::now();
            let _ = matrix::fib_matrix_fast(n);
            total_time += start.elapsed();
        }

        let avg_time = total_time / iterations;
        println!("  n = {:6}: avg {:?}", n, avg_time);
    }
    println!();
}

fn profile_memory_usage() {
    println!("📊 Memory Analysis");
    println!("──────────────────");

    println!("  Iterative: O(1) - uses only 2 u128 values = 32 bytes");
    println!("  Matrix:    O(1) - uses 2x2 matrix of u128 = 64 bytes");
    println!("  Recursive Memo: O(n) - allocates Vec<u128>");
    println!();

    // Demonstrate recursive memo memory
    for n in [100, 1000, 10000] {
        let mem_bytes = (n + 1) * 16; // u128 = 16 bytes
        println!("  Recursive Memo (n={}): ~{} bytes allocated", n, mem_bytes);
    }
    println!();
}

fn profile_scaling() {
    println!("📊 Scaling Analysis (Iterative vs Matrix)");
    println!("──────────────────────────────────────────");

    let test_values = [100, 1000, 10000, 100000];
    let iterations = 100;

    println!("  {:>10} │ {:>15} │ {:>15} │ {:>10}", "n", "Iterative", "Matrix", "Speedup");
    println!("  ───────────┼─────────────────┼─────────────────┼───────────");

    for n in test_values {
        // Time iterative
        let mut iter_time = Duration::ZERO;
        for _ in 0..iterations {
            let start = Instant::now();
            let _ = iterative::fib_iterative(n);
            iter_time += start.elapsed();
        }
        let iter_avg = iter_time / iterations;

        // Time matrix
        let mut matrix_time = Duration::ZERO;
        for _ in 0..iterations {
            let start = Instant::now();
            let _ = matrix::fib_matrix_fast(n);
            matrix_time += start.elapsed();
        }
        let matrix_avg = matrix_time / iterations;

        let speedup = iter_avg.as_nanos() as f64 / matrix_avg.as_nanos() as f64;

        println!(
            "  {:>10} │ {:>15?} │ {:>15?} │ {:>10.2}x",
            n, iter_avg, matrix_avg, speedup
        );
    }
    println!();
    println!("💡 Matrix method shows O(log n) advantage for large n");
}
