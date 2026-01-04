//! # fib-core
//!
//! Core Fibonacci algorithm implementations with multiple complexity levels.
//!
//! ## Algorithms Provided
//!
//! | Algorithm | Time Complexity | Space Complexity | Best For |
//! |-----------|-----------------|------------------|----------|
//! | Recursive | O(2^n) | O(n) | Demonstration only |
//! | Recursive Memo | O(n) | O(n) | Small n with caching |
//! | Iterative | O(n) | O(1) | General use |
//! | Matrix | O(log n) | O(1) | Large n values |
//! | Binet | O(1) | O(1) | Approximation (n ≤ 78) |
//!
//! ## Example
//!
//! ```rust
//! use fib_core::{iterative, matrix};
//!
//! let fib_20 = iterative::fib_iterative(20);
//! assert_eq!(fib_20, 6765);
//!
//! let fib_100 = matrix::fib_matrix_fast(100);
//! assert_eq!(fib_100, 354224848179261915075);
//! ```

pub mod allocator;
pub mod closed_form;
pub mod iterative;
pub mod matrix;
pub mod memory;
pub mod recursive;

#[cfg(feature = "simd")]
pub mod simd;

// Re-export main functions for convenience
pub use closed_form::{binet_error_analysis, fib_binet_f64};
pub use iterative::{fib_iterative, fib_iterative_batch, fib_iterative_branchless};
pub use matrix::{fib_matrix_fast, fib_matrix_modulo};
pub use recursive::{fib_recursive, fib_recursive_memo};

#[cfg(feature = "simd")]
pub use simd::{fib_simd_batch, SimdBatchCalculator, SimdFeatures};

/// Enum representing available Fibonacci algorithms
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum FibMethod {
    /// Naive recursive - O(2^n)
    Recursive,
    /// Recursive with memoization - O(n)
    RecursiveMemo,
    /// Iterative - O(n)
    Iterative,
    /// Iterative branchless - O(n)
    IterativeBranchless,
    /// Matrix exponentiation - O(log n)
    Matrix,
    /// Binet formula - O(1) with precision limits
    Binet,
}

impl FibMethod {
    /// Calculate Fibonacci using the specified method
    ///
    /// # Arguments
    /// * `n` - The Fibonacci index to calculate
    ///
    /// # Returns
    /// The nth Fibonacci number
    ///
    /// # Panics
    /// May panic for large n with certain methods due to overflow or stack limits
    pub fn calculate(&self, n: u64) -> u128 {
        match self {
            FibMethod::Recursive => fib_recursive(n),
            FibMethod::RecursiveMemo => fib_recursive_memo(n),
            FibMethod::Iterative => fib_iterative(n),
            FibMethod::IterativeBranchless => fib_iterative_branchless(n),
            FibMethod::Matrix => fib_matrix_fast(n),
            FibMethod::Binet => fib_binet_f64(n) as u128,
        }
    }

    /// Get the name of the method
    pub fn name(&self) -> &'static str {
        match self {
            FibMethod::Recursive => "recursive",
            FibMethod::RecursiveMemo => "recursive_memo",
            FibMethod::Iterative => "iterative",
            FibMethod::IterativeBranchless => "iterative_branchless",
            FibMethod::Matrix => "matrix",
            FibMethod::Binet => "binet",
        }
    }

    /// Get the time complexity of the method
    pub fn time_complexity(&self) -> &'static str {
        match self {
            FibMethod::Recursive => "O(2^n)",
            FibMethod::RecursiveMemo => "O(n)",
            FibMethod::Iterative => "O(n)",
            FibMethod::IterativeBranchless => "O(n)",
            FibMethod::Matrix => "O(log n)",
            FibMethod::Binet => "O(1)",
        }
    }

    /// Get the space complexity of the method
    pub fn space_complexity(&self) -> &'static str {
        match self {
            FibMethod::Recursive => "O(n)",
            FibMethod::RecursiveMemo => "O(n)",
            FibMethod::Iterative => "O(1)",
            FibMethod::IterativeBranchless => "O(1)",
            FibMethod::Matrix => "O(1)",
            FibMethod::Binet => "O(1)",
        }
    }
}

impl std::str::FromStr for FibMethod {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s.to_lowercase().as_str() {
            "recursive" => Ok(FibMethod::Recursive),
            "recursive_memo" | "memo" => Ok(FibMethod::RecursiveMemo),
            "iterative" => Ok(FibMethod::Iterative),
            "iterative_branchless" | "branchless" => Ok(FibMethod::IterativeBranchless),
            "matrix" => Ok(FibMethod::Matrix),
            "binet" => Ok(FibMethod::Binet),
            _ => Err(format!("Unknown method: {}", s)),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const FIRST_20_FIBS: [u128; 21] = [
        0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765,
    ];

    #[test]
    fn test_all_methods_match() {
        for (n, expected) in FIRST_20_FIBS.iter().enumerate() {
            let n = n as u64;
            assert_eq!(fib_iterative(n), *expected, "iterative failed at n={}", n);
            assert_eq!(fib_matrix_fast(n), *expected, "matrix failed at n={}", n);

            if n <= 15 {
                assert_eq!(fib_recursive(n), *expected, "recursive failed at n={}", n);
            }

            assert_eq!(
                fib_recursive_memo(n),
                *expected,
                "recursive_memo failed at n={}",
                n
            );
        }
    }

    #[test]
    fn test_large_fibonacci() {
        // F(100) = 354224848179261915075
        let expected = 354224848179261915075u128;
        assert_eq!(fib_iterative(100), expected);
        assert_eq!(fib_matrix_fast(100), expected);
    }

    #[test]
    fn test_fib_method_enum() {
        let method = FibMethod::Iterative;
        assert_eq!(method.calculate(10), 55);
        assert_eq!(method.name(), "iterative");
        assert_eq!(method.time_complexity(), "O(n)");
    }
}
