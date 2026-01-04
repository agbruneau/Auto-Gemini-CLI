# fib-core

The core library for the Fibonacci Benchmark Suite. This crate implements various algorithms for calculating Fibonacci numbers, ranging from basic recursive definitions to optimized matrix exponentiation and SIMD-accelerated batch processing.

## Algorithms

| Algorithm | Time Complexity | Space Complexity | Description |
|-----------|-----------------|------------------|-------------|
| Recursive | O(2^n) | O(n) | Naive implementation (slow). |
| Recursive Memo | O(n) | O(n) | Recursive with memoization. |
| Iterative | O(n) | O(1) | Standard iterative loop. |
| Matrix | O(log n) | O(1) | Matrix exponentiation. |
| Fast Doubling | O(log n) | O(log n) | Optimized matrix approach. |
| Binet | O(1) | O(1) | Closed-form approximation (float). |

## Usage

```rust
use fib_core::{FibMethod, iterative, matrix};

fn main() {
    // Direct usage
    let f100 = iterative::fib_iterative(100);

    // Via enum strategy
    let method = FibMethod::FastDoubling;
    let result = method.calculate(100);

    println!("F(100) = {}", result);
}
```

## Features

*   **SIMD**: Enable `simd` feature for AVX2/AVX512 batch processing support.
