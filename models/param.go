package models

// ParamSignUp 定义注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 定义登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResponseLogin struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

//type ParamPostVote struct {
//	PostId    int64 `json:"post_id" binding:"required"`                // 帖子id
//	Direction int8  `json:"direction" binding:"required,oneof=1,0,-1"` //赞成1，反对-1，取消0
//}

// ParamVoteData 投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`               // 贴子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1" ` // 赞成票(1)还是反对票(-1)取消投票(0)
}
