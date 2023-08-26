namespace go interaction

include "user.thrift"

struct BaseResp {
    1: required i64 code,
    2: optional string msg,
}

struct Video {
    1: required i64 id,
    2: required user.User author,
    3: required string play_url,
    4: required string cover_url,
    5: required i64 favorite_count,
    6: required i64 comment_count,
    7: required bool is_favorite,
    8: required string title,
}

struct Comment {
    1: required i64 id,
    2: required user.User user,
    3: required string content,
    4: required string create_date,
}

struct FavoriteActionRequest {
    1: required i64 video_id,
    2: required i64 action_type,
    3: required string token,
}

struct FavoriteActionResponse {
    1: required BaseResp base,
}

struct FavoriteListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct FavoriteListResponse {
    1: required BaseResp base,
    2: required list<Video> video_list,
}

struct VideoFavoritedCountRequest {
    1: required i64 video_id,
    2: required string token,
}

struct VideoFavoritedCountResponse {
    1: required BaseResp base,
    2: required i64 like_count,
}

struct UserTotalFavoritedRequest {
    1: required string token,
    2: required i64 user_id,
}

struct UserTotalFavoritedResponse {
    1: required BaseResp base,
    2: required i64 total_favorited,
}

struct UserFavoriteCountRequest {
    1: required i64 user_id,
    2: required string token,
}

struct UserFavoriteCountResponse {
    1: required BaseResp base,
    2: required i64 like_count,
}

struct IsFavoriteRequest {
    1: required i64 user_id,
    2: required i64 video_id,
    3: required string token,
}

struct IsFavoriteResponse {
    1: required BaseResp base,
    2: required bool is_favorite,
}

struct CommentActionRequest {
    1: required i64 video_id,
    2: required i64 action_type,
    3: required string token,
    4: optional string comment_text,
    5: optional i64 comment_id,
}

struct CommentActionResponse {
    1: required BaseResp base,
    2: optional Comment comment,
}

struct CommentListRequest {
    1: required i64 video_id,
    2: required string token,
}

struct CommentListResponse {
    1: required BaseResp base,
    2: required list<Comment> comment_list,
}

struct CommentCountRequest {
    1: required i64 video_id,
    2: optional string token,
}

struct CommentCountResponse {
    1: required BaseResp base,
    2: required i64 comment_count,
}

service InteractionService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
    FavoriteListResponse FavoriteList(1 : FavoriteListRequest req)
    VideoFavoritedCountResponse VideoFavoritedCount(1 : VideoFavoritedCountRequest req)
    UserFavoriteCountResponse UserFavoriteCount(1 : UserFavoriteCountRequest req)
    UserTotalFavoritedResponse UserTotalFavorited(1 : UserTotalFavoritedRequest req)
    IsFavoriteResponse IsFavorite(1 : IsFavoriteRequest req)
    CommentActionResponse CommentAction(1 : CommentActionRequest req)
    CommentListResponse CommentList(1 : CommentListRequest req)
    CommentCountResponse CommentCount(1 : CommentCountRequest req)
}