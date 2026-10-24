package main

import "fmt"

func age() {
	age_article, err := newArticle("hello")
	if err != nil {
		fmt.Println(err)
		return
	}
	age_article.Read()

	fmt.Print(age_article.Views)
}
