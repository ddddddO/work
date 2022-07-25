const std = @import("std");

pub fn main() anyerror!void {
    // Note that info level log messages are by default printed only in Debug
    // and ReleaseSafe build modes.
    std.log.info("All your codebase are belong to us!! {s}", .{"hello!"});

    var nn = validFunc(5);
    std.log.info("n is: {d}", .{nn});

    // var in = invalidFunc(5);
    // std.log.info("n is: {d}", .{in});
    var tmp_n: u16 = 5;
    var nnn = validFunc2(&tmp_n);
    std.log.info("n is: {d}", .{nnn});
}

test "basic test" {
    try std.testing.expectEqual(10, 3 + 7);
}

fn validFunc(n: u16) u16 {
    var nn: u16 = 0;
    nn += n;
    return nn;
}

// fn invalidFunc(n: u16) u16 {
//     n += n;
//     return n;
// }

fn validFunc2(n: *u16) u16 {
    n.* += n.*;
    return n.*;
}
