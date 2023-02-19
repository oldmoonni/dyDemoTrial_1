namespace go video

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

struct FeedRequest {
    1: i64 latest_time
    2: string token
}

struct FeedResponse {
    1: list<Video> video_list
    2: i64 next_time
    3: BaseResp base_resp
}

struct PublishListRequest {
    1: i64 user_id
    2: string token
}

struct PublishListResponse {
    1: list<Video> video_list
    2: BaseResp base_resp
}

service VideoService {
    FeedResponse VideoFeed(1:FeedRequest req)
    PublishListResponse VideoPublishList(1:PublishListRequest req)
}