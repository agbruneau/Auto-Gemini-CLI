//! Fibonacci Performance Benchmark Suite CLI
//!
//! A comprehensive command-line tool for calculating and benchmarking
//! Fibonacci numbers using various algorithms.

use clap::{Parser, Subcommand};

mod commands;

#[derive(Parser)]
#[command(name = "fib-bench")]
#[command(author = "FibBenchmark Team")]
#[command(version = "0.1.0")]
#[command(about = "🔬 Fibonacci Performance Benchmark Suite", long_about = None)]
#[command(propagate_version = true)]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    /// Calculate Fibonacci number with a specific method
    Calc {
        /// The Fibonacci index to calculate
        #[arg(short, long)]
        n: u64,

        /// Algorithm to use: recursive, iterative, matrix, binet
        #[arg(short, long, default_value = "iterative")]
        method: String,

        /// Show timing information
        #[arg(short, long)]
        time: bool,
    },

    /// Run quick comparison of all algorithms
    Compare {
        /// The Fibonacci index to calculate
        #[arg(short, long)]
        n: u64,

        /// Maximum n for recursive (to avoid long waits)
        #[arg(long, default_value = "30")]
        max_recursive: u64,
    },

    /// Run the Criterion benchmarks
    Bench {
        /// Filter benchmarks by name
        #[arg(short, long)]
        filter: Option<String>,
    },

    /// Show algorithm complexity information
    Info {
        /// Algorithm to show info for (or "all")
        #[arg(short, long, default_value = "all")]
        method: String,
    },

    /// Generate a sequence of Fibonacci numbers
    Sequence {
        /// Number of Fibonacci numbers to generate
        #[arg(short, long, default_value = "20")]
        count: u64,

        /// Starting index
        #[arg(short, long, default_value = "0")]
        start: u64,
    },

    /// Analyze Binet formula accuracy
    BinetAnalysis {
        /// Maximum n to analyze
        #[arg(short, long, default_value = "100")]
        max_n: u64,
    },
}

fn main() {
    let cli = Cli::parse();

    match cli.command {
        Commands::Calc { n, method, time } => {
            commands::calc::run(n, &method, time);
        }
        Commands::Compare { n, max_recursive } => {
            commands::compare::run(n, max_recursive);
        }
        Commands::Bench { filter } => {
            commands::bench::run(filter);
        }
        Commands::Info { method } => {
            commands::info::run(&method);
        }
        Commands::Sequence { count, start } => {
            commands::sequence::run(count, start);
        }
        Commands::BinetAnalysis { max_n } => {
            commands::binet_analysis::run(max_n);
        }
    }
}
