use assert_cmd::Command;
use predicates::prelude::*;

#[test]
fn test_help_command() {
    let mut cmd = Command::cargo_bin("fib-bench").unwrap();
    // Help can be on stdout or stderr depending on clap version/config
    // We just check that it runs successfully and mentions "Usage"
    cmd.arg("--help")
        .assert()
        .success()
        .stdout(predicate::str::contains("Usage"));
}

#[test]
fn test_calc_command() {
    let mut cmd = Command::cargo_bin("fib-bench").unwrap();
    cmd.arg("calc")
        .arg("--n")
        .arg("10")
        .assert()
        .success()
        .stdout(predicate::str::contains("55"));
}

#[test]
fn test_calc_command_iterative_method() {
    let mut cmd = Command::cargo_bin("fib-bench").unwrap();
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
    let mut cmd = Command::cargo_bin("fib-bench").unwrap();
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
    let mut cmd = Command::cargo_bin("fib-bench").unwrap();
    cmd.arg("info")
        .assert()
        .success()
        .stdout(predicate::str::contains("Algorithm"));
}

#[test]
fn test_invalid_command() {
    let mut cmd = Command::cargo_bin("fib-bench").unwrap();
    cmd.arg("invalid_cmd").assert().failure();
}
