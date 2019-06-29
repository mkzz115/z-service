namespace go forum
namespace php forum
namespace py forum
namespace cpp forum
namespace java forum

include "base.thrift"

struct PublishArticleReq {
    1: i64 Uid
    2: string Topic
    3: string Content
    4: i64 PublishTs
}

struct PublishArticleRes {
    1: optional i32 Status
    100: optional base.ErrorInfo ErrInfo
}

service ForumService {
    PublishArticleRes PublishArticle(1:PublishArticleReq req)
}