package logic

import (
	"go.uber.org/zap"
	"web-app/dao/mysql"
	"web-app/dao/redis"
	"web-app/models"
	"web-app/pkg/snowflake"
)

func CreatePost(p *models.Post) error {
	//生成post_id
	p.PostId = snowflake.GenID()
	//创建帖子
	if err := mysql.CreatePost(p); err != nil {
		return err
	}
	return redis.CreatePostCache(p)

}

func GetPostDetail(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询帖子
	post, err := mysql.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailById failed", zap.Int64("post_id", pid), zap.Error(err))
		return
	}
	//查询帖子的作者
	author, err := mysql.GetUserById(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetUserById failed", zap.Int64("post.AuthorId", post.AuthorId), zap.Error(err))
		return
	}
	// 查询帖子所属社区详情
	community, err := mysql.GetCommunityDetailById(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("post.CommunityId", post.CommunityId), zap.Error(err))
		return
	}
	// 拼接数据
	data = &models.ApiPostDetail{
		Username:  author.Username,
		Post:      post,
		Community: community,
	}
	return
}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	// 查找所有的帖子
	postList, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(postList))
	for _, post := range postList {
		// 查找帖子作者
		author, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Int64("post.AuthorId", post.AuthorId), zap.Error(err))
			continue
		}
		// 查询帖子所属社区详情
		community, err := mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("post.CommunityId", post.CommunityId), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			Username:  author.Username,
			Post:      post,
			Community: community,
		}
		data = append(data, postDetail)
	}
	return
}

// todo
func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 去redis 取出帖子列表
	ids, err := redis.GetPostIdsInOrder(p)
	zap.L().Debug("redis.GetPostIdsInOrder", zap.Any("ids", ids))
	if err != nil {
		zap.L().Error("redis.GetPostIdsInOrder failed ", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIdsInOrder len(ids) == 0 ")
		return
	}
	// 根据帖子列表查询所有帖子数据
	postList, err := mysql.GetPostListByIds(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListByIds(ids) failed", zap.Error(err))
		return
	}
	// 查询帖子每个帖子对应的分数
	voteScore := redis.GetPostVoteData(ids)

	for idx, post := range postList {
		// 查找帖子作者
		author, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Int64("post.AuthorId", post.AuthorId), zap.Error(err))
			continue
		}
		// 查询帖子所属社区详情
		community, err := mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("post.CommunityId", post.CommunityId), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			Username:  author.Username,
			VoteScore: voteScore[idx],
			Post:      post,
			Community: community,
		}
		data = append(data, postDetail)
	}

	return
}

func GetCommunityPostList(p *models.ParamCommunityPostList) (data []*models.ApiPostDetail, err error) {

	// 去redis 取出帖子列表
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	zap.L().Debug("redis.GetPostIdsInOrder", zap.Any("ids", ids))
	if err != nil {
		zap.L().Error("redis.GetPostIdsInOrder failed ", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIdsInOrder len(ids) == 0 ")
		return
	}
	// 根据帖子列表查询所有帖子数据
	postList, err := mysql.GetPostListByIds(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListByIds(ids) failed", zap.Error(err))
		return
	}
	// 查询帖子每个帖子对应的分数
	voteScore := redis.GetPostVoteData(ids)

	for idx, post := range postList {
		// 查找帖子作者
		author, err := mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorId) failed", zap.Int64("post.AuthorId", post.AuthorId), zap.Error(err))
			continue
		}
		// 查询帖子所属社区详情
		community, err := mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("post.CommunityId", post.CommunityId), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			Username:  author.Username,
			VoteScore: voteScore[idx],
			Post:      post,
			Community: community,
		}
		data = append(data, postDetail)
	}

	return
}
