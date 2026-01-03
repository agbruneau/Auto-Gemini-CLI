//! Memory profiling utilities

use crate::allocator::TrackingAllocator;

/// Snapshot of memory statistics
#[derive(Debug, Clone, Copy)]
pub struct MemoryStats {
    pub current_bytes: usize,
    pub allocations: usize,
}

impl MemoryStats {
    pub fn now(allocator: &TrackingAllocator) -> Self {
        Self {
            current_bytes: allocator.get_current_usage(),
            allocations: allocator.get_allocation_count(),
        }
    }

    /// Calculate delta from a previous snapshot
    pub fn delta(&self, start: &MemoryStats) -> Self {
        Self {
            current_bytes: self.current_bytes.saturating_sub(start.current_bytes),
            allocations: self.allocations.saturating_sub(start.allocations),
        }
    }
}
