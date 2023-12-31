// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ChangeStatus struct {
	Total     *int `json:"total,omitempty"`
	Additions *int `json:"additions,omitempty"`
	Deletions *int `json:"deletions,omitempty"`
}

type File struct {
	FileName *string `json:"fileName,omitempty"`
	Type     *string `json:"type,omitempty"`
	Language *string `json:"language,omitempty"`
	RawURL   *string `json:"rawURL,omitempty"`
	Size     *int    `json:"size,omitempty"`
}

type FileInput struct {
	FileName string `json:"fileName"`
	Content  string `json:"content"`
}

type ForkOf struct {
	ID    *string `json:"id,omitempty"`
	URL   *string `json:"url,omitempty"`
	Files []*File `json:"files"`
}

type Gist struct {
	ID          *string `json:"id,omitempty"`
	Description *string `json:"description,omitempty"`
	Public      *bool   `json:"public,omitempty"`
	Owner       *User   `json:"Owner,omitempty"`
	Files       []*File `json:"files"`
	Comments    *int    `json:"Comments,omitempty"`
	HTMLURL     *string `json:"htmlURL,omitempty"`
	GitPullURL  *string `json:"gitPullURL,omitempty"`
	GitPushURL  *string `json:"gitPushURL,omitempty"`
	CreatedAt   *string `json:"createdAt,omitempty"`
	UpdatedAt   *string `json:"updatedAt,omitempty"`
	NodeID      *string `json:"nodeID,omitempty"`
}

type GistComment struct {
	ID        *int    `json:"id,omitempty"`
	URL       *string `json:"url,omitempty"`
	Body      *string `json:"body,omitempty"`
	User      *User   `json:"user,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
}

type GistCommit struct {
	URL          *string       `json:"url,omitempty"`
	Version      *string       `json:"version,omitempty"`
	User         *User         `json:"user,omitempty"`
	ChangeStatus *ChangeStatus `json:"change_status,omitempty"`
	CommittedAt  *string       `json:"committed_at,omitempty"`
	NodeID       *string       `json:"nodeID,omitempty"`
}

type GistInput struct {
	Description string       `json:"description"`
	Public      bool         `json:"public"`
	Files       []*FileInput `json:"files"`
}

type History struct {
	User         *User         `json:"user,omitempty"`
	Version      *string       `json:"version,omitempty"`
	CommittedAt  *string       `json:"committedAt,omitempty"`
	ChangeStatus *ChangeStatus `json:"changeStatus,omitempty"`
	URL          *string       `json:"URL,omitempty"`
}

type User struct {
	Name      *string `json:"name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Login     *string `json:"login,omitempty"`
	AvatarURL *string `json:"avatarURL,omitempty"`
	URL       *string `json:"URL,omitempty"`
}
