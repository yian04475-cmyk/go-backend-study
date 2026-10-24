package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Article struct {
	Title string
	Views int
}

type FileArticleSaver struct {
	filePath string
}

type ArticleSaver interface {
	SaveArticle(Article) error
}

func publishArticle(title string, saver ArticleSaver) error {
	article, err := newArticle(title)
	if err != nil {
		return err
	}
	return saver.SaveArticle(article)
}

func newArticle(title string) (Article, error) {
	if title == "" {
		return Article{}, errors.New("标题不能为空")
	}
	return Article{Title: title}, nil
}

func (a *Article) Read() {
	a.Views++
}

func (a Article) summary() string {
	return fmt.Sprintf("title: %s read: %d\n", a.Title, a.Views)
}

func (s FileArticleSaver) SaveArticle(article Article) error {
	return saveArticle(article, s.filePath)
}

func saveArticle(article Article, filePath string) error {
	data, err := json.Marshal(article)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func loadArticle(filePath string) (Article, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Article{}, err
	}
	article := Article{}
	if err := json.Unmarshal(data, &article); err != nil {
		return Article{}, err
	}
	return article, nil
}

func main() {
	filePath := "article.json"
	fileSaver := FileArticleSaver{filePath: filePath}
	if err := publishArticle("测试", fileSaver); err != nil {
		fmt.Println("发布文章失败:", err)
		return
	}

	savedArticle, err := loadArticle(filePath)
	if err != nil {
		fmt.Println("读取文章失败:", err)
		return
	}
	fmt.Print(savedArticle.summary())
}
