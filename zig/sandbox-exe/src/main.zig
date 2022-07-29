const std = @import("std");
const log = std.log;

const Allocator = std.mem.Allocator;
const Stream = std.net.Stream;

const HttpClient = @import("http-client").HttpClient;

pub fn main() anyerror!void {
    const allocator = std.heap.page_allocator;
    const host = "www.google.com";
    // const host = "www.yahoo.co.jp";

    const client = HttpClient.init(allocator);
    const res = try client.req()
        .setHeader("Accept: text/html")
        .get(host);

    defer res.deinit();

    const writer = std.io.getStdOut().writer();
    try writer.print("Status Line: {s}\n", .{res.statusLine()});
    try writer.print("Status Code: {s}\n", .{res.statusCode()});
    try writer.print("Status: {s}\n", .{res.status()});
}
