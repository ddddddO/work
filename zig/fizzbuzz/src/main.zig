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

test "success" {
    const allocator = std.heap.page_allocator;

    {
        const fizz_num = 3;
        const fizz_ret = fizzbuzz(allocator, fizz_num);
        defer allocator.free(fizz_ret);
        try std.testing.expect(std.mem.eql(u8, "FIZZ: 3", fizz_ret));
    }

    {
        const buzz_num = 5;
        const buzz_ret = fizzbuzz(allocator, buzz_num);
        defer allocator.free(buzz_ret);
        try std.testing.expect(std.mem.eql(u8, "BUZZ: 5", buzz_ret));
    }

    {
        const fizzbuzz_num = 15;
        const fizzbuzz_ret = fizzbuzz(allocator, fizzbuzz_num);
        defer allocator.free(fizzbuzz_ret);
        try std.testing.expect(std.mem.eql(u8, "FIZZ-BUZZ: 15", fizzbuzz_ret));
    }

    {
        const zero_num = 0;
        const zero_ret = fizzbuzz(allocator, zero_num);
        defer allocator.free(zero_ret);
        try std.testing.expect(std.mem.eql(u8, "FIZZ-BUZZ: 0", zero_ret));
    }

    {
        const num = 128;
        const ret = fizzbuzz(allocator, num);
        defer allocator.free(ret);
        try std.testing.expect(std.mem.eql(u8, "", ret));
    }
}
