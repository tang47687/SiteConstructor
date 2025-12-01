package main

import (
	"fmt"
	"net/http"
)

type Site struct {
	Name   string
	Pages  []string
	Domain string
}

func (s *Site) CreateSite(name, domain string) {
	s.Name = name
	s.Domain = domain
	s.Pages = []string{}
	fmt.Printf("Сайт '%s' успешно создан на домене '%s'\n", s.Name, s.Domain)
}

func (s *Site) AddPage(page string) {
	s.Pages = append(s.Pages, page)
	fmt.Printf("Страница '%s' добавлена к сайту '%s'\n", page, s.Name)
}

func main() {
	site := &Site{}
	site.CreateSite("Мой Первый Сайт", "example.com")
	site.AddPage("Главная")
	site.AddPage("Контакты")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Информация о сайте: %s на домене %s. Страницы: %v", site.Name, site.Domain, site.Pages)
	})

	http.ListenAndServe(":8080", nil)
}
