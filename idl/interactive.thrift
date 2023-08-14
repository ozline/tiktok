namespace go interactive

include "user.thrift"

struct BaseResp {
    1: required i64 code,
    2: optional string msg,
}

struct Comment{
    1: required i64 id
    2: required user.User user
    3: required string content
    4: required string create_date
}

struct FavoriteActionRequest {
    1: required string video_id,
    2: required string action_type,
    3: required string token,
}

struct FavoriteActionResponse {
    1: required BaseResp base,
}

struct FavoriteListRequest {
    1: required string user_id,
    2: required string token,
}

struct FavoriteListResponse {
    1: required BaseResp base,
}

struct CommentActionRequest {
    1: required string video_id,
    2: required string action_type,
    3: required string token,
    4: optional string comment_text,
    5: optional string comment_id, 
}

struct CommentActionResponse {
    1: required BaseResp base,
    2: optional Comment comment
}

struct CommentListRequest {
    1: required string video_id,
    2: required string token,
}

struct CommentListResponse {
    1: required BaseResp base,
    2: required list<Comment> comment_list
}

service InteractiveService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
    FavoriteListResponse FavoriteList(1 : FavoriteListRequest req)
    CommentActionResponse CommentAction(1 : CommentActionRequest req)
    CommentListResponse CommentList(1 : CommentListRequest req)
}