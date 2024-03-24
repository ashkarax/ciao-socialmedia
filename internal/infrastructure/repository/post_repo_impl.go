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

	query := "SELECT post_id,caption,created_at FROM posts WHERE user_id=$1 AND post_status=$2"
	err := d.DB.Raw(query, userId, "normal").Scan(&response)
	if err.Error != nil {
		return &response, err.Error
	}
	return &response, nil
}
func (d *PostRepo) GetPostMediaById(postId *string) (*[]string, error) {
	var response []string

	query := "SELECT media_url FROM post_media WHERE post_id=$1"
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
	query := "SELECT COUNT(*) FROM posts WHERE user_id=$1"
	if err := d.DB.Raw(query, userId).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil

}
