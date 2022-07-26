const std = @import("std");

pub fn main() anyerror!void {
    const allocator = std.heap.page_allocator;
    const args = try std.process.argsAlloc(allocator);
    defer allocator.free(args);

    const prog = args[0];
    if (args.len != 2) {
        std.log.err("usage: {s} [max number]", .{prog});
        return; // TODO: exit status 1
    }

    const max = std.fmt.parseUnsigned(u32, args[1], 10) catch |err| {
        std.log.err("specify positive integer for argument. error: {s}", .{err});
        return; // TODO: exit status 1
    };
    const stdout = std.io.getStdOut().writer();
    try fizzbuzz(stdout, max);
}

fn fizzbuzz(writer: anytype, max: u32) anyerror!void {
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
