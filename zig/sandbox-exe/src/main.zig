const std = @import("std");
const log = std.log;

const Allocator = std.mem.Allocator;
const Stream = std.net.Stream;

pub fn main() anyerror!void {
    std.log.info("Request...", .{});

    const allocator = std.heap.page_allocator;
    const host = "www.google.com";
    // const host = "www.yahoo.co.jp";

    const client = HttpClient.init(allocator);
    defer client.deinit();
    try client.req().get(host);
}

pub const HttpClient = struct {
    allocator: Allocator,

    pub fn init(allocator: Allocator) HttpClient {
        log.debug("INIT!!", .{});

        return HttpClient{
            .allocator = allocator,
        };
    }

    pub fn deinit(self: HttpClient) void {
        _ = self;
        log.debug("DEINIT!", .{});
    }

    pub fn req(self: HttpClient) Request {
        _ = self;
        return Request{
            .allocator = self.allocator,
        };
    }
};

pub const Request = struct {
    allocator: Allocator,

    // TODO: return Response
    fn get(self: Request, target: []const u8) !void {
        const tcp_conn = try std.net.tcpConnectToHost(self.allocator, target, 80);
        defer tcp_conn.close();

        log.debug("IN Request: {s}", .{target});

        _ = try tcp_conn.write("GET / HTTP/1.1\r\n");
        const host_header = try std.fmt.allocPrint(self.allocator, "Host: {s}\r\n", .{target});
        defer self.allocator.free(host_header);

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
};

// pub const Response = struct {
//     buf: []u8,

//     pub fn init(allocator: Allocator) Response {
//         return Response{

//         }
//     }
// }

// test "basic test" {
//     try std.testing.expectEqual(10, 3 + 7);
// }
