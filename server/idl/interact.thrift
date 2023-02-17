namespace go interact

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

struct Video {
    1: i64 id
    2: User author
    3: string play_url
    4: string cover_url
    5: i64 favorite_count
    6: i64 comment_count
    7: bool is_favorite
    8: string title
}

struct Comment {
	1: i64 id
	2: User user
	3: string Content
	4: string CreateDate
}

struct FavoriteActionRequest {
    1: string token
    2: i64 video_id
    3: i32 action_type
}

struct FavoriteActionResponse {
    1: BaseResp base_resp
}

struct FavoriteListRequest {
    1: i64 user_id
    2: string token
}

struct FavoriteListResponse {
    1:list<Video> video_list
    1: BaseResp base_resp
}

struct CommentActionRequest {
    1: string token
    2: i64 video_id
    3: i32 action_type
    4: string comment_text
    5: i64 comment_id
}

struct CommentActionResponse {
    1: Comment comment
    2: BaseResp base_resp
}

struct CommentListRequest {
    1: string token
    2: i64 video_id
}

struct CommentListResponse {
    1: list<Comment> comment_list
    2: BaseResp base_resp
}

service InteractService {
    FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest req)
    FavoriteListResponse FavoriteList(1:FavoriteListRequest req)
    CommentActionResponse CommentAction(1:CommentActionRequest req)
    CommentListResponse CommentList(1:CommentListRequest req)
}