package handlers

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User User  `json:"user,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

type FriendUserListResponse struct {
	Response
	UserList []FriendUser `json:"user_list"`
}

type MessageChatResponse struct {
	Response
	MessageList []Message `json:"message_list,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type User struct {
	Id            		int64  `json:"id,omitempty"`
	Name          		string `json:"name,omitempty"`
	FollowCount   		int64  `json:"follow_count,omitempty"`
	FollowerCount 		int64  `json:"follower_count,omitempty"`
	IsFollow      		bool   `json:"is_follow,omitempty"`
	Avatar      		string   `json:"avatar,omitempty"`
	BackgroundImage     string   `json:"background_image,omitempty"`
	Signature      		string   `json:"signature,omitempty"`
	TotalFavorited      int64   `json:"total_favorited,omitempty"`
	WorkCount      		int64   `json:"work_count,omitempty"`
	FavoriteCount      	int64   `json:"favorite_count,omitempty"`
}

type FriendUser struct{
	Id            		int64  `json:"id,omitempty"`
	Name          		string `json:"name,omitempty"`
	FollowCount   		int64  `json:"follow_count,omitempty"`
	FollowerCount 		int64  `json:"follower_count,omitempty"`
	IsFollow      		bool   `json:"is_follow,omitempty"`
	Avatar      		string   `json:"avatar,omitempty"`
	BackgroundImage     string   `json:"background_image,omitempty"`
	Signature      		string   `json:"signature,omitempty"`
	TotalFavorited      int64   `json:"total_favorited,omitempty"`
	WorkCount      		int64   `json:"work_count,omitempty"`
	FavoriteCount      	int64   `json:"favorite_count,omitempty"`
	Message		      	string   `json:"message,omitempty"`
	MsgType      		int64   `json:"msgType,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	FromUserId   int64  `json:"from_user_id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime	 int64 `json:"create_time,omitempty"`
}