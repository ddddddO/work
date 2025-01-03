const std = @import("std");
const Allocator = std.mem.Allocator;

pub fn main() anyerror!void {
    const allocator = std.heap.page_allocator;
    const args = try std.process.argsAlloc(allocator);
    defer allocator.free(args);

    run(allocator, args) catch |err| {
        switch (err) {
            ArgumentError.InvalidCount => {
                const prog = args[0];
                std.log.err("usage: {s} [max number]", .{prog});
            },
            ArgumentError.Invalid => {
                std.log.err("specify positive integer for argument.", .{});
            },
            else => {
                std.log.err("run error. {s}", .{err});
            },
        }
        allocator.free(args);
        std.os.exit(1);
    };
}

const ArgumentError = error{
    InvalidCount,
    Invalid,
};

fn run(
    allocator: Allocator,
    args: [][:0]u8,
) anyerror!void {
    if (args.len != 2) {
        return ArgumentError.InvalidCount;
    }

    const max = std.fmt.parseUnsigned(u32, args[1], 10) catch {
        return ArgumentError.Invalid;
    };

    const writer = std.io.getStdOut().writer();

    try writer.print("{s}-MAX:{d} start!\n", .{ "fizzbuzz", max });

    var i: u32 = 0;
    while (i <= max) : (i += 1) {
        const ret = try fizzbuzz(allocator, i);
        defer allocator.free(ret);
        if (ret.len == 0) continue;
        try writer.print("{s}\n", .{ret});
    }
}

fn fizzbuzz(allocator: Allocator, i: u32) ![]u8 {
    if ((i % 3 == 0) and (i % 5 == 0)) {
        return std.fmt.allocPrint(allocator, "FIZZ-BUZZ: {d}", .{i});
    }

    if (i % 3 == 0) {
        return std.fmt.allocPrint(allocator, "FIZZ: {d}", .{i});
    }

    if (i % 5 == 0) {
        return std.fmt.allocPrint(allocator, "BUZZ: {d}", .{i});
    }

    return "";
}

test "success" {
    const allocator = std.testing.allocator;

    {
        const fizz_num = 3;
        const fizz_ret = try fizzbuzz(allocator, fizz_num);
        defer allocator.free(fizz_ret);
        try std.testing.expect(std.mem.eql(u8, "FIZZ: 3", fizz_ret));
    }

    {
        const buzz_num = 5;
        const buzz_ret = try fizzbuzz(allocator, buzz_num);
        defer allocator.free(buzz_ret);
        try std.testing.expect(std.mem.eql(u8, "BUZZ: 5", buzz_ret));
    }

    {
        const fizzbuzz_num = 15;
        const fizzbuzz_ret = try fizzbuzz(allocator, fizzbuzz_num);
        defer allocator.free(fizzbuzz_ret);
        try std.testing.expect(std.mem.eql(u8, "FIZZ-BUZZ: 15", fizzbuzz_ret));
    }

    {
        const zero_num = 0;
        const zero_ret = try fizzbuzz(allocator, zero_num);
        defer allocator.free(zero_ret);
        try std.testing.expect(std.mem.eql(u8, "FIZZ-BUZZ: 0", zero_ret));
    }

    {
        const num = 128;
        const ret = try fizzbuzz(allocator, num);
        defer allocator.free(ret);
        try std.testing.expect(std.mem.eql(u8, "", ret));
    }
}
