package graph

import (
	"sync"

	"GRPC/graphql-1/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	books       []*model.Book
	authors     []*model.Author
	mu          sync.Mutex
	counter     int
	subscribers []chan *model.Book
}

func (r *Resolver) SeedData() {
	author1 := &model.Author{ID: "author-1", Name: "Alan Donovan"}
	author2 := &model.Author{ID: "author-2", Name: "Katherine Cox-Buday"}

	book1 := &model.Book{ID: "book-1", Title: "The Go Programming Language", Year: 2015, Author: author1}
	book2 := &model.Book{ID: "book-2", Title: "The Practice of Programming", Year: 1999, Author: author1}
	book3 := &model.Book{ID: "book-3", Title: "Concurrency in Go", Year: 2017, Author: author2}

	author1.Books = []*model.Book{book1, book2}
	author2.Books = []*model.Book{book3}

	r.authors = []*model.Author{author1, author2}
	r.books = []*model.Book{book1, book2, book3}
	r.counter = 3
}
