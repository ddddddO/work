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
    try run(allocator, stdout, max);
}

fn run(allocator: anytype, writer: anytype, max: u32) anyerror!void {
    try writer.print("{s}-MAX:{d} start!\n", .{ "fizzbuzz", max });

    var i: u32 = 0;
    while (i <= max) : (i += 1) {
        const ret = fizzbuzz(allocator, i);
        defer allocator.free(ret);
        if (ret.len == 0) {
            continue;
        }
        try writer.print("{s}\n", .{ret});
    }
}

// TODO: catch -> try に変更してerrをthrowする。
fn fizzbuzz(allocator: anytype, i: u32) []u8 {
    if ((i % 3 == 0) and (i % 5 == 0)) {
        const ret = std.fmt.allocPrint(
            allocator,
            "FIZZ-BUZZ: {d}",
            .{i},
        ) catch |err| {
            std.log.err("error: {s}\n", .{err});
            return "";
        };
        return ret;
    }

    if (i % 3 == 0) {
        const ret = std.fmt.allocPrint(
            allocator,
            "FIZZ: {d}",
            .{i},
        ) catch |err| {
            std.log.err("error: {s}\n", .{err});
            return "";
        };
        return ret;
    }

    if (i % 5 == 0) {
        const ret = std.fmt.allocPrint(
            allocator,
            "BUZZ: {d}",
            .{i},
        ) catch |err| {
            std.log.err("error: {s}\n", .{err});
            return "";
        };
        return ret;
    }

    return "";
}
