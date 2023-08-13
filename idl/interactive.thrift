namespace go interactive

include "user.thrift"

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct Comment{
    1: i64 id
    2: user.User user
    3: string content
    4: string create_date
}

struct FavoriteActionRequest {
    1: string video_id,
    2: string action_type,
    3: string token,
}

struct FavoriteActionResponse {
    1: BaseResp base,
}

struct FavoriteListRequest {
    1: string user_id,
    2: string token,
}

struct FavoriteListResponse {
    1: BaseResp base,
}

struct CommentActionRequest {
    1: string video_id,
    2: string action_type,
    3: string token,
    4: optional string comment_text,
    5: optional string comment_id, 
}

struct CommentActionResponse {
    1: BaseResp base,
    2: Comment comment
}

struct CommentListRequest {
    1: string video_id,
    2: string token,
}

struct CommentListResponse {
    1: BaseResp base,
    2: list<Comment> comment_list
}

service InteractiveService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
    FavoriteListResponse FavoriteList(1 : FavoriteListRequest req)
    CommentActionResponse CommentAction(1 : CommentActionRequest req)
    CommentListResponse CommentList(1 : CommentListRequest req)
}