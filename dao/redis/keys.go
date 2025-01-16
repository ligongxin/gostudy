package redis

// 定义redis中用到的key

const (
	KeyPrefix              = "go:"
	KeyPostTime            = "post:time"   // zSet 记录发帖时间
	KeyPostScore           = "post:score"  // zSet 帖子分数
	KeyPostVotedZSetPrefix = "post:voted:" // zSet 记录帖子及投票类型 参数是post_id
	KeyUserToken           = "user:token:" // string 用户的token，需要传user_id
)
