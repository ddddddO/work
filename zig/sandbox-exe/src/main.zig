const std = @import("std");
const log = std.log;

pub fn main() anyerror!void {
    std.log.info("Request...", .{});

    const allocator = std.heap.page_allocator;
    const host = "www.google.com";
    // const host = "www.yahoo.co.jp";
    const tcp_conn = try std.net.tcpConnectToHost(allocator, host, 80);

    _ = try tcp_conn.write("GET / HTTP/1.1\r\n");
    const host_header = "Host: " ++ host ++ "\r\n";
    _ = try tcp_conn.write(host_header);
    _ = try tcp_conn.write("\r\n");

    const writer = std.io.getStdOut().writer();
    while (true) {
        var response_buffer: [2048]u8 = undefined;
        const len = tcp_conn.read(&response_buffer) catch 0;
        if (len == 0) {
            log.debug("Response end.", .{});
            break;
        }
        const response = response_buffer[0..len];
        try writer.writeAll(response);

        const end_response_1 = std.mem.eql(u8, "\r\n", response_buffer[len - 2 .. len]);
        if (end_response_1) {
            log.debug("Response end..", .{});
            break;
        }
        const end_response_2 = std.mem.eql(u8, "\r\n ", response_buffer[len - 3 .. len]);
        if (end_response_2) {
            log.debug("Response end...", .{});
            break;
        }
    }
}

// test "basic test" {
//     try std.testing.expectEqual(10, 3 + 7);
// }
