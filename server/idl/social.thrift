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
	6:string Avatar
	7:string BackgroundImage
	8:string Signature
	9:i64 TotalFavorited
	10:i64 WorkCount
	11:i64 FavoriteCount
}

struct FriendUser {
    1: i64 Id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
	6: string Avatar
	7: string BackgroundImage
	8: string Signature
	9: i64 TotalFavorited
	10: i64 WorkCount
	11: i64 FavoriteCount
    12: string message
    13: i64 msgType
}

struct Message {
    1: i64 Id
    2: i64 ToUserId
    3: i64 FromUserId
    4: string Content
    5: i64 CreateTime
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

struct FriendListRequest {
    1: i64 user_id
    2: string token
}

struct FriendListResponse {
    1: list<FriendUser> user_list
    2: BaseResp base_resp
}

struct MessageChatRequest {
    1: string token
    2: i64 toUserId
    3: i64 preMsgTime
}

struct MessageChatResponse {
    1: list<Message> message_list
    2: BaseResp base_resp
}

struct MessageActionRequest {
    1: string token
    2: i64 toUserId
    3: i32 actionType
    4: string content
}

struct MessageActionResponse {
    1: BaseResp base_resp
}

service SocialService {
    RelationActionResponse RelationAction(1:RelationActionRequest req)
    FollowListResponse FollowList(1:FollowListRequest req)
    FollowerListResponse FollowerList(1:FollowerListRequest req)
    FriendListResponse FriendList(1:FriendListRequest req)
    MessageChatResponse MessageChat(1:MessageChatRequest req)
    MessageActionResponse MessageAction(1:MessageActionRequest req)
}