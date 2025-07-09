package controller

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/samber/lo"

	"github.com/qww83728/gsam.git/domain/entity"
	repo_entity "github.com/qww83728/gsam.git/domain/entity/repo"
	repo "github.com/qww83728/gsam.git/domain/repository"
	userSvc "github.com/qww83728/gsam.git/domain/service/user"
)

type UserController interface {
	InitSqlLite()
	GetVideo() ([]repo_entity.Video, error)

	GetUser(id string) (string, error)
	TestMap(input entity.TestMap) error
	TestRoutine() error
}

type UserControllerImpl struct {
	userService userSvc.UserService
	videoRepo   repo.VideoRepo
}

func NewUserController(
	userService userSvc.UserService,
	videoRepo repo.VideoRepo,
) UserController {
	return &UserControllerImpl{
		userService: userService,
		videoRepo:   videoRepo,
	}
}

func (c *UserControllerImpl) GetUser(
	id string,
) (string, error) {
	if err := c.userService.CheckID(id); err != nil {
		return "", err
	}

	return fmt.Sprintf("Hello %s", id), nil
}

func (c *UserControllerImpl) TestMap(
	input entity.TestMap,
) error {
	testMap := make(map[string]entity.UserPost, 0)
	for _, user := range input.Maps {
		testMap[user.ID] = user
	}
	fmt.Println("#### testMap:", testMap)

	// array to map
	loMap := lo.KeyBy(input.Maps, func(u entity.UserPost) string {
		return u.ID
	})
	fmt.Println("#### loMap:", loMap)

	// array to array
	loIDArray := lo.Map(input.Maps, func(u entity.UserPost, _ int) string {
		return u.ID
	})
	fmt.Println("#### loIDArray:", loIDArray)

	// map to array
	loUserArray := lo.Values(loMap)
	fmt.Println("#### loUserArray:", loUserArray)

	return nil
}

func (c *UserControllerImpl) TestRoutine() error {
	maxRoutines := 30
	totalJob := 1000

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results := make(chan int, totalJob)
	buffer := make(chan struct{}, maxRoutines) // 控制 goroutine 上限
	var wg sync.WaitGroup

	for i := 1; i <= totalJob; i++ {
		buffer <- struct{}{} // 佔一個名額，如果滿了就阻塞
		wg.Add(1)

		jobID := i // 避免 goroutine capture 問題
		go func(results chan int, jobID int) {
			defer wg.Done()
			defer func() { <-buffer }() // 釋放名額
			fmt.Println("#### Job ID :", jobID)
			routineSum(ctx, i, results)

			// 模擬延遲
			time.Sleep(10 * time.Millisecond)
		}(results, jobID)
	}

	// 正確等待 goroutine 全部完成後才關閉 channel
	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for r := range results {
		total += r
	}

	fmt.Printf("#### 總和為: %d\n", total)
	return nil
}

func routineSum(ctx context.Context, id int, results chan<- int) {
	// defer wg.Done()

	sum := 0
	maxValue := 1000
	for i := 0; i < maxValue; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("routine %d 被取消\n", id)
			return
		default:
			sum += 1
		}
	}
	// fmt.Println("#### routine sum:", sum)
	results <- sum
}

func (c *UserControllerImpl) TestSqlLite(
	input entity.TestMap,
) error {
	testMap := make(map[string]entity.UserPost, 0)
	for _, user := range input.Maps {
		testMap[user.ID] = user
	}
	fmt.Println("#### testMap:", testMap)

	// array to map
	loMap := lo.KeyBy(input.Maps, func(u entity.UserPost) string {
		return u.ID
	})
	fmt.Println("#### loMap:", loMap)

	// array to array
	loIDArray := lo.Map(input.Maps, func(u entity.UserPost, _ int) string {
		return u.ID
	})
	fmt.Println("#### loIDArray:", loIDArray)

	// map to array
	loUserArray := lo.Values(loMap)
	fmt.Println("#### loUserArray:", loUserArray)

	return nil
}

func (c *UserControllerImpl) InitSqlLite() {
	c.videoRepo.SetupDatabase()
}

func (c *UserControllerImpl) GetVideo() ([]repo_entity.Video, error) {
	videos, err := c.videoRepo.GetVideo()
	if err != nil {
		return nil, err
	}

	return videos, nil
}
