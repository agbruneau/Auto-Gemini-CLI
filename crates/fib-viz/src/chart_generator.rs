use crate::data_parser::BenchmarkData;
use plotly::common::{Mode, Title};
use plotly::layout::{Axis, AxisType, Layout};
use plotly::{Plot, Scatter};
use std::path::Path;

pub fn generate_charts(data: &BenchmarkData, output_dir: &str) {
    let dir = Path::new(output_dir);
    std::fs::create_dir_all(dir).ok();

    // 1. Complexity Chart
    let mut plot = Plot::new();

    let n_values: Vec<u64> = data.complexity.iter().map(|p| p.n).collect();
    let iter_times: Vec<u128> = data.complexity.iter().map(|p| p.iterative_ns).collect();
    let matrix_times: Vec<u128> = data.complexity.iter().map(|p| p.matrix_ns).collect();

    let trace1 = Scatter::new(n_values.clone(), iter_times)
        .name("Iterative")
        .mode(Mode::LinesMarkers);

    let trace2 = Scatter::new(n_values, matrix_times)
        .name("Matrix Exponentiation")
        .mode(Mode::LinesMarkers);

    plot.add_trace(trace1);
    plot.add_trace(trace2);

    let layout = Layout::new()
        .title(Title::new("Algorithm Complexity Comparison"))
        .x_axis(Axis::new().title(Title::new("n (Fibonacci Index)")))
        .y_axis(Axis::new().title(Title::new("Time (ns)")));

    plot.set_layout(layout);
    plot.write_html(dir.join("complexity_chart.html"));

    // 2. Binet Accuracy Chart
    let mut plot = Plot::new();

    let n_values: Vec<u64> = data.accuracy.iter().map(|p| p.n).collect();
    // Use log scale for error if possible or just raw error
    let errors: Vec<f64> = data.accuracy.iter().map(|p| p.abs_error).collect();

    let trace = Scatter::new(n_values, errors)
        .name("Absolute Error")
        .mode(Mode::LinesMarkers);

    plot.add_trace(trace);

    let layout = Layout::new()
        .title(Title::new("Binet Formula Accuracy"))
        .x_axis(Axis::new().title(Title::new("n")))
        .y_axis(Axis::new().title(Title::new("Absolute Error")));

    plot.set_layout(layout);
    plot.write_html(dir.join("binet_accuracy_chart.html"));

    // 3. Golden Ratio Convergence
    let mut plot = Plot::new();

    let n_values: Vec<u64> = data.golden_ratio.iter().map(|p| p.n).collect();
    let errors: Vec<f64> = data.golden_ratio.iter().map(|p| p.error_from_phi).collect();

    let trace = Scatter::new(n_values, errors)
        .name("Deviation from Phi")
        .mode(Mode::LinesMarkers);

    plot.add_trace(trace);

    let layout = Layout::new()
        .title(Title::new("Golden Ratio Convergence"))
        .x_axis(Axis::new().title(Title::new("n")))
        .y_axis(
            Axis::new()
                .title(Title::new("Error |Ratio - Phi|"))
                .type_(AxisType::Log),
        ); // Log scale is useful here

    plot.set_layout(layout);
    plot.write_html(dir.join("golden_ratio_chart.html"));
}
