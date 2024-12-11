package services

import (
	"context"
	"errors"
	"fmt"
	"news-app/internal/db"
)

// CreateNews creates a new news article
func CreateNews(news *db.News) error {
	_, err := db.DB.NewInsert().Model(news).Exec(context.Background())
	if err != nil {
		return errors.New("failed to create news: " + err.Error())
	}
	return nil
}

// GetAllNews retrieves all news articles
func GetAllNews() ([]db.News, error) {
	var newsList []db.News
	err := db.DB.NewSelect().Model(&newsList).Scan(context.Background())
	if err != nil {
		return nil, errors.New("failed to fetch news: " + err.Error())
	}
	return newsList, nil
}

// GetNewsByID retrieves a single news article by ID
func GetNewsByID(id int64) (*db.News, error) {
	news := &db.News{}
	err := db.DB.NewSelect().Model(news).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return nil, errors.New("news not found: " + err.Error())
	}
	return news, nil
}

func GetNewsByCat(cat string) ([]db.News, error) {
	fmt.Println("Fetching news for category:", cat)

	var newsList []db.News
	err := db.DB.NewSelect().
		Model(&newsList).
		Where("category = ?", cat).
		Scan(context.Background())
	if err != nil {
		return nil, errors.New("news not found: " + err.Error())
	}

	return newsList, nil
}

func GetCategories() ([]string, error) {
	var categories []string

	err := db.DB.NewSelect().
		Model((*db.News)(nil)).
		ColumnExpr("DISTINCT category").
		Scan(context.Background(), &categories)
	if err != nil {
		return nil, errors.New("failed to fetch categories: " + err.Error())
	}

	return categories, nil
}

// UpdateNews updates an existing news article
func UpdateNews(id int64, updatedNews *db.News) error {
	_, err := db.DB.NewUpdate().
		Model(updatedNews).
		OmitZero().
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		return errors.New("failed to update news: " + err.Error())
	}
	return nil
}

// DeleteNews deletes a news article by ID
func DeleteNews(id int64) error {
	_, err := db.DB.NewDelete().
		Model((*db.News)(nil)).
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		return errors.New("failed to delete news: " + err.Error())
	}
	return nil
}
