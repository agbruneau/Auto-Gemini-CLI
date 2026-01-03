//! Fibonacci Visualization Tool
//!
//! Generates charts and visualizations for benchmark results

use fib_core::{closed_form, iterative, matrix};
use std::fs;

fn main() {
    println!("📊 Fibonacci Visualization Generator");
    println!("=====================================");
    println!();

    // Create results directory
    fs::create_dir_all("results").ok();

    generate_complexity_data();
    generate_accuracy_data();
    generate_golden_ratio_data();

    println!();
    println!("✅ Data files generated in results/");
    println!("   Use your favorite plotting tool to visualize!");
}

fn generate_complexity_data() {
    println!("📈 Generating complexity comparison data...");

    let mut csv = String::from("n,iterative_ns,matrix_ns\n");

    for n in (10..=1000).step_by(10) {
        let iterations = 100;

        // Time iterative
        let start = std::time::Instant::now();
        for _ in 0..iterations {
            let _ = iterative::fib_iterative(n);
        }
        let iter_ns = start.elapsed().as_nanos() / iterations as u128;

        // Time matrix
        let start = std::time::Instant::now();
        for _ in 0..iterations {
            let _ = matrix::fib_matrix_fast(n);
        }
        let matrix_ns = start.elapsed().as_nanos() / iterations as u128;

        csv.push_str(&format!("{},{},{}\n", n, iter_ns, matrix_ns));
    }

    fs::write("results/complexity_comparison.csv", csv).ok();
    println!("   ✓ results/complexity_comparison.csv");
}

fn generate_accuracy_data() {
    println!("📈 Generating Binet accuracy data...");

    let mut csv = String::from("n,exact,binet,abs_error,rel_error\n");

    for n in 0..=100 {
        let exact = iterative::fib_iterative(n);
        let binet = closed_form::fib_binet_f64(n);
        let (abs_error, rel_error) = closed_form::binet_error_analysis(n);

        csv.push_str(&format!(
            "{},{},{},{},{}\n",
            n, exact, binet, abs_error, rel_error
        ));
    }

    fs::write("results/binet_accuracy.csv", csv).ok();
    println!("   ✓ results/binet_accuracy.csv");
}

fn generate_golden_ratio_data() {
    println!("📈 Generating golden ratio convergence data...");

    let mut csv = String::from("n,ratio,error_from_phi\n");
    let phi = closed_form::PHI;

    for n in 1..=50 {
        let ratio = closed_form::fibonacci_ratio(n);
        let error = (ratio - phi).abs();

        csv.push_str(&format!("{},{},{}\n", n, ratio, error));
    }

    fs::write("results/golden_ratio_convergence.csv", csv).ok();
    println!("   ✓ results/golden_ratio_convergence.csv");
}
