//! Custom allocator for tracking memory usage
//!
//! This module provides a `TrackingAllocator` that wraps the system allocator
//! and keeps track of total allocated bytes. This is useful for profiling
//! and understanding the memory footprint of different algorithms.

use std::alloc::{GlobalAlloc, Layout, System};
use std::sync::atomic::{AtomicUsize, Ordering};

/// A wrapper around the system allocator that tracks memory usage
pub struct TrackingAllocator {
    allocator: System,
    allocated_bytes: AtomicUsize,
    allocation_count: AtomicUsize,
}

impl TrackingAllocator {
    pub const fn new() -> Self {
        Self {
            allocator: System,
            allocated_bytes: AtomicUsize::new(0),
            allocation_count: AtomicUsize::new(0),
        }
    }

    /// Reset the statistics
    pub fn reset(&self) {
        self.allocated_bytes.store(0, Ordering::SeqCst);
        self.allocation_count.store(0, Ordering::SeqCst);
    }

    /// Get total bytes currently allocated (net)
    /// Note: This simple implementation only tracks growing usage loosely
    /// as deallocations might not exactly match allocations in complex scenarios,
    /// but for benchmarking specific runs it gives a good delta.
    ///
    /// Actually, for a precise "currently allocated", we subtract on dealloc.
    /// For "total traffic", we might want another counter.
    /// Let's track:
    /// 1. Current usage (alloc - dealloc)
    pub fn get_current_usage(&self) -> usize {
        self.allocated_bytes.load(Ordering::SeqCst)
    }

    pub fn get_allocation_count(&self) -> usize {
        self.allocation_count.load(Ordering::SeqCst)
    }
}

unsafe impl GlobalAlloc for TrackingAllocator {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        let ptr = self.allocator.alloc(layout);
        if !ptr.is_null() {
            self.allocated_bytes
                .fetch_add(layout.size(), Ordering::SeqCst);
            self.allocation_count.fetch_add(1, Ordering::SeqCst);
        }
        ptr
    }

    unsafe fn dealloc(&self, ptr: *mut u8, layout: Layout) {
        self.allocator.dealloc(ptr, layout);
        self.allocated_bytes
            .fetch_sub(layout.size(), Ordering::SeqCst);
    }
}
