package model

import (
	"time"
)

type CreateGistRequest struct {
	Description string
	Public      bool
	Files       []*CreateFile
}

type CreateFile struct {
	FileName string
	Content  string
}

type Gist struct {
	ID          string
	Description string
	Public      bool
	Owner       *User
	Files       []*File
	Comments    int
	HTMLURL     string
	GitPullURL  string
	GitPushURL  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	NodeID      string
}

type User struct {
	Name      string
	Email     string
	Login     string
	AvatarURL string
	URL       string
}

type File struct {
	FileName string
	Type     string
	Language string
	RawURL   string
	Size     int
	Content  string
}

type Comment struct {
	ID        int64
	URL       string
	Body      string
	User      *User
	CreatedAt time.Time
}

type Commit struct {
	URL          string
	Version      string
	User         *User
	ChangeStatus *CommitStats
	CommittedAt  time.Time
	NodeID       string
}

type CommitStats struct {
	Additions int
	Deletions int
	Total     int
}
