use assert_cmd::Command;
use predicates::prelude::*;

#[test]
fn test_help_command() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    // Help can be on stdout or stderr depending on clap version/config
    // We just check that it runs successfully and mentions "Usage"
    cmd.arg("--help")
        .assert()
        .success()
        .stdout(predicate::str::contains("Usage"));
}

#[test]
fn test_calc_command() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("calc")
        .arg("--n")
        .arg("10")
        .assert()
        .success()
        .stdout(predicate::str::contains("55"));
}

#[test]
fn test_calc_command_iterative_method() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("calc")
        .arg("--n")
        .arg("10")
        .arg("--method")
        .arg("iterative")
        .assert()
        .success()
        .stdout(predicate::str::contains("55"));
}

#[test]
fn test_calc_command_fast_doubling_method() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("calc")
        .arg("--n")
        .arg("10")
        .arg("--method")
        .arg("fast_doubling")
        .assert()
        .success()
        .stdout(predicate::str::contains("55"));
}

#[test]
fn test_info_command() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("info")
        .assert()
        .success()
        .stdout(predicate::str::contains("Algorithm"));
}

#[test]
fn test_invalid_command() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("invalid_cmd").assert().failure();
}

#[test]
fn test_compare_command() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("compare")
        .arg("--n")
        .arg("10")
        .assert()
        .success()
        .stdout(predicate::str::contains("Iterative"))
        .stdout(predicate::str::contains("Matrix"));
}

#[test]
fn test_sequence_command() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("sequence")
        .arg("--count")
        .arg("5")
        .assert()
        .success()
        .stdout(predicate::str::contains("0"))
        .stdout(predicate::str::contains("1"))
        .stdout(predicate::str::contains("2"))
        .stdout(predicate::str::contains("3"));
}

#[test]
fn test_sequence_command_offset() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("sequence")
        .arg("--count")
        .arg("3")
        .arg("--start")
        .arg("5")
        .assert()
        .success()
        .stdout(predicate::str::contains("5")) // F(5)
        .stdout(predicate::str::contains("8")) // F(6)
        .stdout(predicate::str::contains("13")); // F(7)
}

#[test]
fn test_binet_analysis_command() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("binet-analysis")
        .arg("--max-n")
        .arg("10")
        .assert()
        .success()
        .stdout(predicate::str::contains("Accuracy"));
}

#[test]
fn test_memory_command() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("memory")
        .arg("--n")
        .arg("100")
        .assert()
        .success()
        .stdout(predicate::str::contains("Memory"));
}

#[test]
fn test_calc_invalid_method() {
    let mut cmd = Command::new(env!("CARGO_BIN_EXE_fib-bench"));
    cmd.arg("calc")
        .arg("--n")
        .arg("10")
        .arg("--method")
        .arg("invalid_method_name")
        .assert()
        .failure();
}
