package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewArticle(t *testing.T) {
	testCases := []struct {
		title   string
		wantErr bool
	}{
		{title: "第一篇", wantErr: false},
		{title: "", wantErr: true},
	}

	for _, testCase := range testCases {
		article, err := newArticle(testCase.title)
		if (err != nil) != testCase.wantErr {
			t.Fatalf("newArticle(%q) 的错误 = %v，预期是否报错 = %t", testCase.title, err, testCase.wantErr)
		}
		if err == nil && (article.Title != testCase.title || article.Views != 0) {
			t.Fatalf("newArticle(%q) 得到 %+v，标题或阅读数不对", testCase.title, article)
		}
	}
}

func TestSaveAndLoadArticle(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "article.json")
	want := Article{Title: "测试文章", Views: 2}
	if err := saveArticle(want, filePath); err != nil {
		t.Fatalf("保存文章失败: %v", err)
	}

	got, err := loadArticle(filePath)
	if err != nil {
		t.Fatalf("读取文章失败: %v", err)
	}
	if got.Title != want.Title || got.Views != want.Views {
		t.Fatalf("读回的文章 = %+v，预期 = %+v", got, want)
	}
}

func TestPublishArticle(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "article.json")
	fileSaver := FileArticleSaver{filePath: filePath}
	if err := publishArticle("测试", fileSaver); err != nil {
		t.Fatalf("发布文章失败: %v", err)
	}

	article, err := loadArticle(filePath)
	if err != nil {
		t.Fatalf("读取发布后的文章失败: %v", err)
	}
	if article.Title != "测试" || article.Views != 0 {
		t.Fatalf("发布后的文章 = %+v，预期标题为测试、阅读数为 0", article)
	}
}

func TestPublishArticleEmptyTitle(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "article.json")
	fileSaver := FileArticleSaver{filePath: filePath}
	if err := publishArticle("", fileSaver); err == nil {
		t.Fatal("空标题发布时应返回错误")
	}
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		t.Fatalf("空标题不应创建文件，检查文件时得到的错误: %v", err)
	}
}
