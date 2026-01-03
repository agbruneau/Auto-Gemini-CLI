use serde::{Deserialize, Serialize};
use std::fs;
use std::path::Path;

#[derive(Debug, Serialize, Deserialize)]
pub struct ComplexityPoint {
    pub n: u64,
    pub iterative_ns: u128,
    pub matrix_ns: u128,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct AccuracyPoint {
    pub n: u64,
    pub exact: u128,
    pub binet: f64,
    pub abs_error: f64,
    pub rel_error: f64,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct GoldenRatioPoint {
    pub n: u64,
    pub ratio: f64,
    pub error_from_phi: f64,
}

#[derive(Debug)]
pub struct BenchmarkData {
    pub complexity: Vec<ComplexityPoint>,
    pub accuracy: Vec<AccuracyPoint>,
    pub golden_ratio: Vec<GoldenRatioPoint>,
}

impl BenchmarkData {
    pub fn load(results_dir: &str) -> std::io::Result<Self> {
        let dir = Path::new(results_dir);

        let complexity_path = dir.join("complexity_comparison.json");
        let accuracy_path = dir.join("binet_accuracy.json");
        let golden_path = dir.join("golden_ratio_convergence.json");

        let complexity_json = fs::read_to_string(complexity_path)?;
        let complexity: Vec<ComplexityPoint> = serde_json::from_str(&complexity_json)?;

        let accuracy_json = fs::read_to_string(accuracy_path)?;
        let accuracy: Vec<AccuracyPoint> = serde_json::from_str(&accuracy_json)?;

        let golden_json = fs::read_to_string(golden_path)?;
        let golden_ratio: Vec<GoldenRatioPoint> = serde_json::from_str(&golden_json)?;

        Ok(BenchmarkData {
            complexity,
            accuracy,
            golden_ratio,
        })
    }
}
