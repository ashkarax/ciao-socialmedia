package usecase

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	interfaceRepository "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/repository/interfaces"
	interfaceUseCase "github.com/ashkarax/ciao-socialmedia/internal/infrastructure/usecase/interfaces"
	requestmodels "github.com/ashkarax/ciao-socialmedia/internal/models/request_models"
	responsemodels "github.com/ashkarax/ciao-socialmedia/internal/models/response_models"
	aws "github.com/ashkarax/ciao-socialmedia/pkg/aws_s3"
	"github.com/go-playground/validator/v10"
)

type PostUseCase struct {
	PostRepo interfaceRepository.IPostRepo
}

func NewPostUseCase(postRepo interfaceRepository.IPostRepo) interfaceUseCase.IPostUseCase {
	return &PostUseCase{PostRepo: postRepo}
}

func (r *PostUseCase) AddNewPost(postData *requestmodels.AddPostData) (*responsemodels.AddPostResp, error) {
	var respPostData responsemodels.AddPostResp
	BucketFolder := "ciao-socialmedia/posts/"

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(postData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {
				case "Caption":
					respPostData.Caption = "should contain less than 60 letters"
				case "UserId":
					respPostData.UserId = "No userId got"
				case "Media":
					respPostData.Media = "you can't add a post without a image/video"

				}
			}
			fmt.Println(err)
			return &respPostData, err
		}
	}
	numFiles := len(postData.Media)
	if numFiles < 1 || numFiles > 5 {
		return &respPostData, errors.New("you can only add 5 image/video in a post")
	}

	for _, media := range postData.Media {
		if media.Size > 5*1024*1024 { // 5 MB limit
			return &respPostData, errors.New("file size exceeds the limit (5MB)")
		}
	}

	allowedTypes := map[string]struct{}{
		"image/jpeg":      {},
		"image/png":       {},
		"image/gif":       {},
		"video/mp4":       {},
		"video/quicktime": {},
	}

	for _, file := range postData.Media {
		// Open the file to read the header
		file, err := file.Open()
		if err != nil {
			return &respPostData, err
		}
		defer file.Close()

		// Read the first 512 bytes to determine the content type
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			return &respPostData, err
		}

		// Reset the file position after reading
		_, err = file.Seek(0, 0)
		if err != nil {
			return &respPostData, err
		}

		// Get the content type based on the file content
		contentType := http.DetectContentType(buffer)

		// Check if the content type is allowed
		if _, ok := allowedTypes[contentType]; !ok {
			return &respPostData, errors.New("unsupported file type,should be a jpeg,png,gif,mp4 or quicktime")
		}
	}

	sess, errInit := aws.AWSSessionInitializer()
	if errInit != nil {
		fmt.Println(errInit)
		return &respPostData, errInit
	}

	for i, file := range postData.Media {
		mediaURL, err := aws.AWSS3MediaUploader(file, sess, &BucketFolder)
		if err != nil {
			fmt.Printf("Error uploading file %d: %v\n", i+1, err)
			return &respPostData, err
		}
		postData.MediaURLs = append(postData.MediaURLs, *mediaURL)
	}

	insertErr := r.PostRepo.AddNewPost(postData)
	if insertErr != nil {
		fmt.Println(insertErr)
		return &respPostData, insertErr
	}
	return &respPostData, nil

}

func (r *PostUseCase) GetAllPostByUser(userId *string) (*[]responsemodels.PostData, error) {
	postData, err := r.PostRepo.GetAllActivePostByUser(userId)
	if err != nil {
		return postData, err
	}
	for i, split := range *postData {
		postIdString := strconv.FormatUint(uint64(split.PostId), 10)
		postMedias, err := r.PostRepo.GetPostMediaById(&postIdString)
		if err != nil {
			return postData, err
		}
		(*postData)[i].MediaUrl = *postMedias

		currentTime := time.Now()
		duration := currentTime.Sub((*postData)[i].CreatedAt)

		minutes := int(duration.Minutes())
		hours := int(duration.Hours())
		days := int(duration.Hours() / 24)
		months := int(duration.Hours() / 24 / 7)

		if minutes < 60 {
			(*postData)[i].PostAge = fmt.Sprintf("%d mins ago", minutes)
		} else if hours < 24 {
			(*postData)[i].PostAge = fmt.Sprintf("%d hrs ago", hours)
		} else if days < 30 {
			(*postData)[i].PostAge = fmt.Sprintf("%d dy ago", days)
		} else {
			(*postData)[i].PostAge = fmt.Sprintf("%d weks ago", months)
		}
	}

	return postData, nil
}

func (r *PostUseCase) DeletePost(postId *requestmodels.PostId, userId *string) error {

	err := r.PostRepo.DeletePostById(&postId.PostId, userId)
	if err != nil {
		return err
	}
	err2 := r.PostRepo.DeletePostMedias(&postId.PostId)
	if err2 != nil {
		return err2
	}

	return nil
}

func (r *PostUseCase) LikePost(inputData *requestmodels.LikeRequest) error {
	err := r.PostRepo.LikePost(inputData)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostUseCase) UnLikePost(inputData *requestmodels.LikeRequest) error {
	err := r.PostRepo.UnLikePost(inputData)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostUseCase) GetAllRelatedPostsForHomeScreen(userId *string) (*[]responsemodels.PostData, error) {
	postData, err := r.PostRepo.GetAllActiveRelatedPostsForHomeScreen(userId)
	if err != nil {
		return postData, err
	}
	for i, split := range *postData {
		postIdString := strconv.FormatUint(uint64(split.PostId), 10)
		postMedias, err := r.PostRepo.GetPostMediaById(&postIdString)
		if err != nil {
			return postData, err
		}
		(*postData)[i].MediaUrl = *postMedias

		currentTime := time.Now()
		duration := currentTime.Sub((*postData)[i].CreatedAt)

		minutes := int(duration.Minutes())
		hours := int(duration.Hours())
		days := int(duration.Hours() / 24)
		months := int(duration.Hours() / 24 / 7)

		if minutes < 60 {
			(*postData)[i].PostAge = fmt.Sprintf("%d mins ago", minutes)
		} else if hours < 24 {
			(*postData)[i].PostAge = fmt.Sprintf("%d hrs ago", hours)
		} else if days < 30 {
			(*postData)[i].PostAge = fmt.Sprintf("%d dy ago", days)
		} else {
			(*postData)[i].PostAge = fmt.Sprintf("%d weks ago", months)
		}
	}

	return postData, nil
}
