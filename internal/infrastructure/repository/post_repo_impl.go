package repository

import (
	"errors"
	"time"

	interfaceRepository "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/repository/interfaces"
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
	"gorm.io/gorm"
)

type PostRepo struct {
	DB *gorm.DB
}

func NewPostRepo(db *gorm.DB) interfaceRepository.IPostRepo {
	return &PostRepo{DB: db}
}

func (d *PostRepo) AddNewPost(postData *requestmodels.AddPostData) error {
	var PostId string

	query := "INSERT INTO Posts (user_id,caption,created_at) VALUES ($1,$2,$3) RETURNING post_id"
	err := d.DB.Raw(query, postData.UserId, postData.Caption, time.Now()).Scan(&PostId).Error
	if err != nil {
		return err
	}

	mediaInsQuery := "INSERT INTO post_media (post_id,media_url) VALUES ($1,$2)"

	for _, url := range postData.MediaURLs {
		errIns := d.DB.Exec(mediaInsQuery, PostId, url).Error
		if errIns != nil {
			return errIns
		}
	}

	return nil
}

func (d *PostRepo) GetAllActivePostByUser(userId *string) (*[]responsemodels.PostData, error) {
	var response []responsemodels.PostData

	query := "SELECT post_id,caption,created_at FROM posts WHERE user_id=$1 AND post_status=$2 ORDER BY created_at DESC"
	err := d.DB.Raw(query, userId, "normal").Scan(&response)
	if err.Error != nil {
		return &response, err.Error
	}
	return &response, nil
}
func (d *PostRepo) GetPostMediaById(postId *string) (*[]string, error) {
	var response []string

	query := "SELECT media_url FROM post_media WHERE post_id=$1 ORDER BY media_id DESC"
	err := d.DB.Raw(query, *postId).Scan(&response).Error
	if err != nil {
		return &response, err
	}

	return &response, nil
}

func (d *PostRepo) DeletePostById(postId *string, userId *string) error {
	query := "DELETE FROM posts WHERE post_id=$1 AND user_id=$2"
	res := d.DB.Exec(query, postId, userId)
	if res.RowsAffected == 0 {
		return errors.New("enter a valid post id,rows affected 0")
	}
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (d *PostRepo) DeletePostMedias(postId *string) error {
	query := "DELETE FROM post_media WHERE post_id=$1"
	res := d.DB.Exec(query, postId).Error
	if res != nil {
		return res
	}
	return nil

}

func (d *PostRepo) GetPostCountOfUser(userId *string) (uint, error) {
	var count uint
	query := "SELECT COUNT(*) FROM posts WHERE user_id=$1 AND post_status=$2"
	if err := d.DB.Raw(query, userId, "normal").Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil

}

func (d *PostRepo) LikePost(inputData *requestmodels.LikeRequest) error {
	query := "INSERT INTO post_likes (user_id,post_id,created_at) VALUES (?,?,?) ON CONFLICT (user_id, post_id) DO NOTHING;"
	err := d.DB.Exec(query, inputData.UserID, inputData.PostID, time.Now()).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PostRepo) UnLikePost(inputData *requestmodels.LikeRequest) error {
	query := "DELETE FROM post_likes WHERE user_id=? AND post_id=?"
	err := d.DB.Exec(query, inputData.UserID, inputData.PostID).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *PostRepo) GetAllActiveRelatedPostsForHomeScreen(userId *string) (*[]responsemodels.PostData, error) {
	var response []responsemodels.PostData

	query := "SELECT posts.*,CASE WHEN post_likes.user_id IS NULL THEN FALSE ELSE TRUE END AS is_liked FROM posts INNER JOIN follow_relationships ON posts.user_id = follow_relationships.following_id LEFT JOIN (SELECT post_id, user_id FROM post_likes WHERE user_id = $1) AS post_likes ON posts.post_id = post_likes.post_id WHERE follow_relationships.follower_id = $1 AND posts.post_status=$2 ORDER BY posts.created_at DESC;"
	err := d.DB.Raw(query, userId, "normal").Scan(&response)
	if err.Error != nil {
		return &response, err.Error
	}
	return &response, nil

}

func (d *PostRepo) LikeAndCommentCountsOfPost(postId *string) (string, error) {
	var likeCount string

	query := "SELECT COUNT(*) FROM post_likes WHERE post_id=?"
	err := d.DB.Raw(query, postId).Scan(&likeCount).Error
	if err != nil {
		return "", err
	}
	return likeCount, nil
}
