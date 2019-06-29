namespace go account
namespace php account
namespace py account
namespace cpp account
namespace java account

include "base.thrift"

struct LogInReq {
    1: string Account
    2: string Password
    3: string Sign
}

struct LogInRes {
    1: optional i64 Uid
    2: optional i32 Status
    3: optional i32 LoginStatus
    100: optional base.ErrorInfo ErrInfo
}

struct LogOutReq {
    1: i64 Uid
}

struct LogOutRes {
    1: optional i32 LoginStatus
    100: optional base.ErrorInfo ErrInfo
}

struct RegisterReq {
    1: string Account
    2: string Password
    3: string Email
    4: string Addr
}

struct RegisterRes {
    1: optional i32 Status
    100: optional base.ErrorInfo ErrInfo
}

struct ChangePwReq {
    1: i64 Uid
    2: string OriginalPw
    3: string NewPw
}

struct ChangePwRes {
    1: optional i32 Status
    100: optional base.ErrorInfo ErrInfo
}


service AccountService {
    LogInRes UserLogIn(1:LogInReq req)
    LogOutRes UserLogOut(1:LogOutReq req)
    RegisterRes UserRegister(1:RegisterReq req)
    ChangePwRes UserChangePw(1:ChangePwReq req)
}