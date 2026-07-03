const std = @import("std");

pub fn main() void {
    std.debug.print("Hello, {s}!\n", .{"World"});
}

const expect = std.testing.expect;

test "always succeeds" {
    try expect(true);
}

const constant: i32 = 5;
var variable: u32 = 5000;

const inferred_constant = @as(i32, 5);
var inferred_variable = @as(u32, 5000);

const a:i32 = undefined;
const b:u32 = undefined;