namespace go follow

struct BaseResp {
    1: i64 code
    2: string msg
}

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: string avatar
    7: string background_image
    8: string signature
    9: i64 total_favorited
    10: i64 work_count
    11: i64 favorite_count
}

struct FriendUser {
    1: User user
    2: string message
    3: i64 msgType // 0 => 当前请求用户接收的消息 1=>当前请求用户发送的消息
}

struct ActionRequest {
    1: string token
    2: i64 to_user_id
    3: i64 action_type // 1-关注, 2-取消关注
}

struct ActionResponse {
    1: BaseResp base
}

struct FollowListRequest {
    1: i64 user_id
    2: string token
}

struct FollowListResponse {
    1: BaseResp base
    2: list<User> user_list
}

struct FollowerListRequest {
    1: i64 user_id
    2: i64 token
}

struct FollowerListResponse {
    1: BaseResp base
    2: list<User> user_list
}

struct FriendListRequest {
    1: i64 user_id
    2: string token
}

struct FriendListResponse {
    1: BaseResp base
    2: list<FriendUser> user_list
}

service FollowService {
    ActionResponse Action(1:ActionRequest req)
    FollowListResponse FollowList(1:FollowListRequest req)
    FollowerListResponse FollowerList(1:FollowerListRequest req)
    FriendListResponse FriendList(1:FriendListRequest req)
}