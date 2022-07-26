const std = @import("std");

pub fn main() anyerror!void {
    const stdout = std.io.getStdOut().writer();
    try fizzbuzz(stdout, 123);
}

fn fizzbuzz(writer: anytype, max: u16) anyerror!void {
    try writer.print("{s}-MAX:{d} start!\n", .{ "fizzbuzz", max });

    var i: u16 = 0;
    while (i <= max) : (i += 1) {
        if ((i % 3 == 0) and (i % 5 == 0)) {
            try writer.print("FIZZ-BUZZ: {d}\n", .{i});
            continue;
        }
        if (i % 3 == 0) {
            try writer.print("FIZZ: {d}\n", .{i});
            continue;
        }
        if (i % 5 == 0) {
            try writer.print("BUZZ: {d}\n", .{i});
            continue;
        }
    }
}
