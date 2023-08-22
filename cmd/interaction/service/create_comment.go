package service

import (
	"strconv"
	"sync"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

// CreateComment create comment
func (s *InteractionService) CreateComment(req *interaction.CommentActionRequest, userId int64) (*db.Comment, error) {

	var wg sync.WaitGroup
	commentModel := &db.Comment{
		VideoId: req.VideoId,
		UserId:  userId,
		Content: *req.CommentText,
	}

	errs := make([]error, 3)
	comment := new(db.Comment)
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		comment, err = db.CreateComment(s.ctx, commentModel)
		errs[0] = err
	}()

	key := strconv.FormatInt(req.VideoId, 10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		exist, err := cache.IsExistComment(s.ctx, key)
		errs[1] = err
		if err != nil {
			return
		}
		if exist == 1 {
			err = cache.DeleteComments(s.ctx, key)
			errs[2] = err
		}
	}()
	wg.Wait()

	for _, err = range errs {
		if err != nil {
			return nil, err
		}
	}

	return comment, nil
}
