namespace go user

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct User {
    1: i64 id,
    2: string name,
    3: i64 follow_count,
    4: i64 follower_count,
    5: bool is_follow
    6: string avatar,
    7: string background_image,
    8: string signature,
    9: i64 total_favorited,
    10: i64 work_count,
    11: i64 favorited_count,
}

struct RegisterRequest {
    1: string username,
    2: string password,
}

struct RegisterResponse {
    1: BaseResp base,
    2: i64 user_id,
    3: string token,
}

struct LoginRequest {
    1: string username,
    2: string password,
}

struct LoginResponse {
    1: BaseResp base,
    2: User user,
    3: string token,
}

struct InfoRequest {
    1: i64 user_id,
    2: string token,
}

struct InfoResponse {
    1: BaseResp base,
    2: User user,
}

service UserService {
    RegisterResponse Register(1: RegisterRequest req),
    LoginResponse Login(1: LoginRequest req),
    InfoResponse Info(1: InfoRequest req),
}