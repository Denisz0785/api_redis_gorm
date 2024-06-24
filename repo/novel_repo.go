package repo

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"redis_gorm_fiber/domain"
	"redis_gorm_fiber/model"
)

type novelRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

// GetNovelById retrieves a novel by its ID from the repository.
// It first checks if the novel is available in Redis. If it is, it retrieves it from Redis.
// If it is not, it retrieves it from MySQL and saves it in Redis.
func (n *novelRepo) GetNovelById(id int) (model.Novel, error) {
	var novels model.Novel
	var ctx = context.Background()

	// Check redis for novel
	result, err := n.rdb.Get(ctx, "novel"+strconv.Itoa(id)).Result()

	// If error is not redis.Nil, return the error
	if err != nil && err != redis.Nil {
		return novels, errors.New("failed to get novel from redis: " + err.Error())
	}

	// If data is available in redis, unmarshal it and return
	if len(result) > 0 {
		if err := json.Unmarshal([]byte(result), &novels); err != nil {
			return novels, errors.New("failed to get novel from redis: " + err.Error())
		}
		return novels, nil
	}

	// If data not available in redis, retrieve it from mysql and save it in redis
	err = n.db.Model(model.Novel{}).Select("id", "name", "author", "description").Where("id = ?", id).Find(&novels).Error
	if err != nil || novels.ID == 0 {
		return novels, errors.New("failed to get novel from mysql: ")
	}
	log.Println("check mysql", novels)
	// Marshal novel to json
	jsonBytes, err := json.Marshal(novels)
	if err != nil {
		return novels, errors.New("failed to get novel from mysql: ")
	}
	jsonString := string(jsonBytes)

	// Save novel to redis
	if err := n.rdb.Set(ctx, "novel"+strconv.Itoa(id), jsonString, 24*time.Hour).Err(); err != nil {
		return novels, errors.New("failed to set novel to redis: ")
	}

	return novels, nil
}

// CreateNovel implements domain.NovelRepo.
func (n *novelRepo) CreateNovel(createNovel model.Novel) error {

	if err := n.db.Create(&createNovel).Error; err != nil {

		return errors.New("failed to create novel")
	}

	return nil
}

func (n *novelRepo) DeleteNovel(id int) error {

	// delete from redis
	if err := n.rdb.Del(context.Background(), "novel"+strconv.Itoa(id)).Err(); err != nil {
		log.Println("failed to delete novel from redis")
	}

	if err := n.db.Where("id=?", id).Delete(&model.Novel{}).Error; err != nil {
		return errors.New("failed to delete novel")
	}
	return nil
}

func (n *novelRepo) UpdateNovel(id int, updateNovel model.Novel) error {

	var novel model.Novel

	if err := n.db.First(&novel, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("novel not found")
		}
		return errors.New("failed to update novel: " + err.Error())
	}
	updateNovel.ID = novel.ID
	// Marshal novel to json
	jsonBytes, err := json.Marshal(updateNovel)
	if err != nil {
		return errors.New("failed to update novel: " + err.Error())
	}
	jsonString := string(jsonBytes)

	// update in redis
	if err := n.rdb.Set(context.Background(), "novel"+strconv.Itoa(id), jsonString, 24*time.Hour).Err(); err != nil {
		return errors.New("failed to update novel in redis")
	}

	// update in mysql
	if err := n.db.Model(&model.Novel{}).Where("id = ?", id).Updates(&updateNovel).Error; err != nil {
		return errors.New("failed to update novel")
	}
	return nil
}

// NewNovelRepo creates a new NovelRepo instance with the given database and Redis connection.
func NewNovelRepo(db *gorm.DB, rdb *redis.Client) domain.NovelRepo {

	return &novelRepo{
		db:  db,  // Set the GORM database connection.
		rdb: rdb, // Set the Redis client connection.
	}
}
