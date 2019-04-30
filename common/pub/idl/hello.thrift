namespace go hello
namespace php hello
namespace py hello
namespace cpp hello
namespace java hello

struct HelloReq {
    1: required i64 uid
}

struct HelloRes {
    1: optional i64 ack
    2: optional i32 errMsg
}


service HelloService {
    HelloRes GetHello(1: HelloReq req)
}
