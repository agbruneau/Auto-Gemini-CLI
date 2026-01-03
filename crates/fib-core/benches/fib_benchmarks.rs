//! Criterion benchmarks for Fibonacci implementations
//!
//! Run with: `cargo bench`

use criterion::{black_box, criterion_group, criterion_main, BenchmarkId, Criterion};
use fib_core::{closed_form, iterative, matrix, recursive};

/// Benchmark comparing algorithm complexities
fn complexity_comparison(c: &mut Criterion) {
    let mut group = c.benchmark_group("complexity_comparison");
    group.sample_size(100);

    // Test small n values where all algorithms are feasible
    for n in [10, 15, 20, 25].iter() {
        // Recursive is only practical for small n
        if *n <= 25 {
            group.bench_with_input(BenchmarkId::new("recursive", n), n, |b, &n| {
                b.iter(|| recursive::fib_recursive(black_box(n)))
            });
        }

        group.bench_with_input(BenchmarkId::new("recursive_memo", n), n, |b, &n| {
            b.iter(|| recursive::fib_recursive_memo(black_box(n)))
        });

        group.bench_with_input(BenchmarkId::new("iterative", n), n, |b, &n| {
            b.iter(|| iterative::fib_iterative(black_box(n)))
        });

        group.bench_with_input(BenchmarkId::new("matrix", n), n, |b, &n| {
            b.iter(|| matrix::fib_matrix_fast(black_box(n)))
        });

        group.bench_with_input(BenchmarkId::new("binet", n), n, |b, &n| {
            b.iter(|| closed_form::fib_binet_f64(black_box(n)))
        });
    }

    group.finish();
}

/// Benchmark scaling behavior for larger n values
fn large_n_scaling(c: &mut Criterion) {
    let mut group = c.benchmark_group("large_n");
    group.sample_size(50);

    for n in [100, 500, 1000, 5000, 10000].iter() {
        group.bench_with_input(BenchmarkId::new("iterative", n), n, |b, &n| {
            b.iter(|| iterative::fib_iterative(black_box(n)))
        });

        group.bench_with_input(BenchmarkId::new("matrix", n), n, |b, &n| {
            b.iter(|| matrix::fib_matrix_fast(black_box(n)))
        });

        group.bench_with_input(BenchmarkId::new("doubling", n), n, |b, &n| {
            b.iter(|| matrix::fib_doubling(black_box(n)))
        });
    }

    group.finish();
}

/// Benchmark iterative variants
fn iterative_variants(c: &mut Criterion) {
    let mut group = c.benchmark_group("iterative_variants");
    group.sample_size(100);

    for n in [50, 100, 500].iter() {
        group.bench_with_input(BenchmarkId::new("standard", n), n, |b, &n| {
            b.iter(|| iterative::fib_iterative(black_box(n)))
        });

        group.bench_with_input(BenchmarkId::new("branchless", n), n, |b, &n| {
            b.iter(|| iterative::fib_iterative_branchless(black_box(n)))
        });
    }

    group.finish();
}

/// Benchmark batch operations
fn batch_operations(c: &mut Criterion) {
    let mut group = c.benchmark_group("batch");
    group.sample_size(50);

    let small_batch: Vec<u64> = (1..=10).collect();
    let medium_batch: Vec<u64> = (1..=50).collect();
    let large_batch: Vec<u64> = (1..=100).collect();

    group.bench_function("batch_10", |b| {
        b.iter(|| iterative::fib_iterative_batch(black_box(&small_batch)))
    });

    group.bench_function("batch_50", |b| {
        b.iter(|| iterative::fib_iterative_batch(black_box(&medium_batch)))
    });

    group.bench_function("batch_100", |b| {
        b.iter(|| iterative::fib_iterative_batch(black_box(&large_batch)))
    });

    group.finish();
}

/// Benchmark cache vs direct calculation
fn cache_vs_direct(c: &mut Criterion) {
    let mut group = c.benchmark_group("cache_vs_direct");
    group.sample_size(100);

    let cache = iterative::FibonacciCache::new(100);
    let queries: Vec<u64> = vec![10, 25, 50, 75, 100];

    group.bench_function("direct_lookups", |b| {
        b.iter(|| {
            queries
                .iter()
                .map(|&n| iterative::fib_iterative(black_box(n)))
                .collect::<Vec<_>>()
        })
    });

    group.bench_function("cached_lookups", |b| {
        b.iter(|| {
            queries
                .iter()
                .map(|&n| cache.get(black_box(n)))
                .collect::<Vec<_>>()
        })
    });

    group.finish();
}

/// Benchmark modular arithmetic
fn modular_arithmetic(c: &mut Criterion) {
    let mut group = c.benchmark_group("modular");
    group.sample_size(50);

    let modulo = 1_000_000_007u128;

    for n in [1000, 10000, 100000].iter() {
        group.bench_with_input(BenchmarkId::new("matrix_mod", n), n, |b, &n| {
            b.iter(|| matrix::fib_matrix_modulo(black_box(n), black_box(modulo)))
        });
    }

    group.finish();
}

criterion_group!(
    benches,
    complexity_comparison,
    large_n_scaling,
    iterative_variants,
    batch_operations,
    cache_vs_direct,
    modular_arithmetic,
);

criterion_main!(benches);
