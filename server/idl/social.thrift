namespace go social

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
}

struct RelationActionRequest {
    1: string token
    2: i64 to_user_id
    3: i32 action_type
}

struct RelationActionResponse {
    1: BaseResp base_resp
}

struct FollowListRequest {
    1: i64 user_id
    2: string token
}

struct FollowListResponse {
    1: list<User> user_list
    2: BaseResp base_resp
}

struct FollowerListRequest {
    1: i64 user_id
    2: string token
}

struct FollowerListResponse {
    1: list<User> user_list
    2: BaseResp base_resp
}

service SocialService {
    RelationActionResponse RelationAction(1:RelationActionRequest req)
    FollowListResponse FollowList(1:FollowListRequest req)
    FollowerListResponse FollowerList(1:FollowerListRequest req)
}