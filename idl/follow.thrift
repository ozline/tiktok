namespace go follow

struct BaseResp {
    1: required i64 code,
    2: optional string msg,
}

struct User {
    1: required i64 id,
    2: required string name,
    3: optional i64 follow_count,
    4: optional i64 follower_count,
    5: required bool is_follow,
    6: optional string avatar,
    7: optional string background_image,
    8: optional string signature,
    9: optional i64 total_favorited,
    10: optional i64 work_count,
    11: optional i64 favorite_count,
}

struct FriendUser {
    1: required User user,
    2: optional string message,
    3: required i64 msgType, // 0 => 当前请求用户接收的消息 1=>当前请求用户发送的消息
}

struct ActionRequest {
    1: required string token
    2: required i64 to_user_id
    3: required i64 action_type // 1-关注, 2-取消关注
}

struct ActionResponse {
    1: required BaseResp base
}

struct FollowListRequest {
    1: required i64 user_id
    2: required string token
}

struct FollowListResponse {
    1: required BaseResp base
    2: optional list<User> user_list
}

struct FollowerListRequest {
    1: required i64 user_id
    2: required string token
}

struct FollowerListResponse {
    1: required BaseResp base
    2: optional list<User> user_list
}

struct FriendListRequest {
    1: required i64 user_id
    2: required string token
}

struct FriendListResponse {
    1: required BaseResp base
    2: optional list<FriendUser> user_list
}

service FollowService {
    ActionResponse Action(1:ActionRequest req)
    FollowListResponse FollowList(1:FollowListRequest req)
    FollowerListResponse FollowerList(1:FollowerListRequest req)
    FriendListResponse FriendList(1:FriendListRequest req)
}