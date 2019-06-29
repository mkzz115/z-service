namespace go order
namespace php order
namespace py order
namespace cpp order
namespace java order

include "base.thrift"

struct GetWishListReq {
    1: i64 Uid
}

struct WishListData {

}

struct GetWishListRes {
    1: optional list<WishListData> items
    100: optional base.ErrorInfo ErrInfo
}

service OrderService {
    GetWishListRes GetWishList(1:GetWishListReq req)
}