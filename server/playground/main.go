package main

import "fmt"

type Comment struct {
	ID       string
	ParentID *string // nil = top-level
	Content  string
	Author   string
	Likes    int
}

type CommentResponse struct {
	ID       string
	ParentID *string // nil = top-level
	Content  string
	Author   string
	Likes    int
	Replies  []*CommentResponse
}

func ptrStr(s string) *string { return &s }

var comments = []Comment{
	{ID: "1", ParentID: nil, Content: "Top-level comment", Author: "alice", Likes: 5},
	{ID: "2", ParentID: ptrStr("1"), Content: "Reply to 1", Author: "bob", Likes: 2},
	{ID: "3", ParentID: ptrStr("1"), Content: "Another reply to 1", Author: "charlie", Likes: 1},
	{ID: "4", ParentID: ptrStr("2"), Content: "Reply to 2", Author: "david", Likes: 0},
	{ID: "5", ParentID: ptrStr("4"), Content: "Reply to 4", Author: "eve", Likes: 3},
	{ID: "6", ParentID: nil, Content: "Second top-level", Author: "frank", Likes: 4},
	{ID: "7", ParentID: ptrStr("6"), Content: "Reply to 6", Author: "grace", Likes: 2},
	{ID: "8", ParentID: ptrStr("7"), Content: "Reply to 7", Author: "henry", Likes: 1},
}

func NestedComments(comments []Comment) []*CommentResponse {
	commentMap := make(map[string]*CommentResponse)

	for _, c := range comments {
		commentMap[c.ID] = &CommentResponse{
			ID:       c.ID,
			ParentID: c.ParentID,
			Content:  c.Content,
			Author:   c.Author,
			Likes:    c.Likes,
			Replies:  []*CommentResponse{},
		}
	}

	topLevel := make([]*CommentResponse, 0)

	for _, c := range comments {
		current := commentMap[c.ID]

		if current.ParentID == nil {
			topLevel = append(topLevel, current)
			continue
		}

		if parent, ok := commentMap[*current.ParentID]; ok {
			parent.Replies = append(parent.Replies, current)
		}
	}

	return topLevel
}

func main() {
	comments := NestedComments(comments)
	for _, c := range comments {
		fmt.Println(*c)
	}
}
