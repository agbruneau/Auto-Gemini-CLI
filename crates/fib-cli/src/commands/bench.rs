//! Benchmark command implementation

pub fn run(filter: Option<String>) {
    println!("📊 Running Criterion Benchmarks...");
    println!();

    if let Some(ref f) = filter {
        println!("Filter: {}", f);
    }

    println!("To run full benchmarks, use:");
    println!();
    println!("  cargo bench");
    println!();
    
    if filter.is_some() {
        println!("  cargo bench -- {}", filter.unwrap());
    }

    println!();
    println!("Benchmark results will be saved to: target/criterion/");
    println!("Open target/criterion/report/index.html for the full report.");
}
