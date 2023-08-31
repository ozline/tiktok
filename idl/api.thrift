namespace go api

// Model

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

struct Video {
    1: required i64 id,
    2: required User author,
    3: required string play_url,
    4: required string cover_url,
    5: required i64 favorite_count,
    6: required i64 comment_count,
    7: required bool is_favorite,
    8: required string title,
}

struct Comment {
    1: required i64 id,
    2: required User user,
    3: required string content,
    4: required string create_date,
}

struct FriendUser {
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
    12: optional string message,
    13: required i64 msgType, // 0 => 当前请求用户接收的消息 1=>当前请求用户发送的消息
}

struct Message {
    1: required i64 id,
    2: required i64 to_user_id,
    3: required i64 from_user_id,
    4: required string content,
    5: required string create_time,
}

// Basic

struct FeedRequest {
    1: optional i64 latest_time,
    2: optional string token,
}

struct FeedResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<Video> video_list,
    4: optional i64 next_time,
}

struct UserRegisterRequest {
    1: required string username,
    2: required string password,
}

struct UserRegisterResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required i64 user_id,
    4: required string token,
}

struct UserLoginRequest {
    1: required string username,
    2: required string password,
}

struct UserLoginResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required i64 user_id,
    4: required string token,
}

struct UserRequest {
    1: required i64 user_id,
    2: required string token,
}

struct UserResponse {
    1: required i64 status_code = 0,
    2: optional i64 status_msg,
    3: required User user,
}

struct PublishActionRequest {
    1: required string token,
    // 2: required binary data,
    2: required string title,
}

struct PublishActionResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
}

struct PublishListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct PublishListResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<Video> video_list,
}

// Interaction

struct FavoriteActionRequest {
    1: required string token,
    2: required i64 video_id,
    3: required i64 action_type, // 1-点赞, 2-取消点赞
}

struct FavoriteActionResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
}

struct FavoriteListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct FavoriteListResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<Video> video_list,
}

struct CommentActionRequest {
    1: required string token,
    2: required i64 video_id,
    3: required i64 action_type,
    4: optional string comment_text,
    5: optional i64 comment_id,
}

struct CommentActionResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: optional Comment comment, // 评论成功才返回
}

struct CommentListRequest {
    1: required string token,
    2: required i64 video_id,
}

struct CommentListResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<Comment> comment_list,
}

// Social

struct RelationActionRequest {
    1: required string token,
    2: required i64 to_user_id,
    3: required i64 action_type, // 1-关注, 2-取消关注
}

struct RelationActionResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
}

struct RelationFollowListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct RelationFollowListResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<User> user_list,
}

struct RelationFollowerListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct RelationFollowerListResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<User> user_list,
}

struct RelationFriendListRequest {
    1: required i64 user_id,
    2: required string token,
}

struct RelationFriendListResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<FriendUser> user_list,
}

struct MessageChatRequest {
    1: required string token,
    2: required i64 to_user_id,
    // 3: required i64 pre_msg_time, // 上次最新消息的时间
}

struct MessageChatResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<Message> message_list,
}

struct MessageActionRequest {
    1: required string token,
    2: required i64 to_user_id,
    3: required i64 action_type, // 1-发送消息
    4: required string content,
}

struct MessageActionResponse {
    1: required i64 status_code = 0,
    2: required string status_msg,
}


service BasicService {
    FeedResponse Feed(1: FeedRequest req) (api.get="/douyin/feed/")

    UserRegisterResponse UserRegister(1: UserRegisterRequest req) (api.post="/douyin/user/register/")
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/douyin/user/login/")
    UserResponse UserInfo(1: UserRequest req) (api.get="/douyin/user/")

    PublishActionResponse PublishAction(1: PublishActionRequest req) (api.post="/douyin/publish/action/")
    PublishListResponse PublishList(1: PublishListRequest req) (api.get="/douyin/publish/list/")
}

service InteractionService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req) (api.post="/douyin/favorite/action/")
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req) (api.get="/douyin/favorite/list/")

    CommentActionResponse CommentAction(1: CommentActionRequest req) (api.post="/douyin/comment/action/")
    CommentListResponse CommentList(1: CommentListRequest req) (api.get="/douyin/comment/list/")
}

service SocialService {
    RelationActionResponse RelationAction(1: RelationActionRequest req) (api.post="/douyin/relation/action/")
    RelationFollowListResponse RelationFollowList(1: RelationFollowListRequest req) (api.get="/douyin/relation/follow/list/")
    RelationFollowerListResponse RelationFollowerList(1: RelationFollowerListRequest req) (api.get="/douyin/relation/follower/list/")
    RelationFriendListResponse RelationFriendList(1: RelationFriendListRequest req) (api.get="/douyin/relation/friend/list/")

    MessageActionResponse MessageAction(1: MessageActionRequest req) (api.post="/douyin/message/action/")
    MessageChatResponse MessageChat(1: MessageChatRequest req) (api.get="/douyin/message/chat/")
}