namespace go user

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct User {
    1:i64 Id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
	6:string Avatar
	7:string BackgroundImage
	8:string Signature
	9:i64 TotalFavorited
	10:i64 WorkCount
	11:i64 FavoriteCount
}

struct UserRegisterRequest {
    1:string name
    2:string password
}

struct UserRegisterResponse {
    1:i64 user_id
    2:string token
    3:BaseResp base_resp
}

struct UserLoginRequest {
    1:string name
    2:string password
}

struct UserLoginResponse {
    1:i64 user_id
    2:string token
    3:BaseResp base_resp
}

struct UserInfoRequest {
    1:i64 user_id
    2:string token
}

struct UserInfoResponse {
    1:User user
    2:BaseResp base_resp
}

service UserService {
    UserRegisterResponse UserRegister(1:UserRegisterRequest req)
    UserLoginResponse UserLogin(1:UserLoginRequest req)
    UserInfoResponse UserInfo(1:UserInfoRequest req)
}

