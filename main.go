package main

import "fmt"

type Article struct {
	Title string
	Views int
}

func newArticle(title string) Article {
	a := Article{}
	a.Title = title
	a.Views = 0
	return a
}

func (a *Article) Read() {
	a.Views++
}

func (a Article) summary() string {
	res := fmt.Sprintf("title: %s read: %d\n", a.Title, a.Views)
	return res
}

func main() {
	article := newArticle("第一篇")
	article.Read()
	article.Read()
	fmt.Printf("%+v\n", article)
	ans := article.summary()
	fmt.Print(ans)
}
