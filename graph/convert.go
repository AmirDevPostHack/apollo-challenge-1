package graph

import (
	graph_model "github.com/amir/graphql-c1/graph/model"
	service_model "github.com/amir/graphql-c1/internal/model"
)

func gistsToGraph(gists []*service_model.Gist) []*graph_model.Gist {
	result := make([]*graph_model.Gist, 0, len(gists))

	for _, gist := range gists {
		result = append(result, gistToGraph(gist))
	}

	return result
}

func gistToGraph(gist *service_model.Gist) *graph_model.Gist {
	createdAtString := gist.CreatedAt.String()
	updatedAtString := gist.UpdatedAt.String()

	return &graph_model.Gist{
		ID:          &gist.ID,
		Description: &gist.Description,
		Public:      &gist.Public,
		Owner:       userToGraph(gist.Owner),
		Files:       filesToGraph(gist.Files),
		Comments:    &gist.Comments,
		HTMLURL:     &gist.HTMLURL,
		GitPullURL:  &gist.GitPullURL,
		GitPushURL:  &gist.GitPullURL,
		CreatedAt:   &createdAtString,
		UpdatedAt:   &updatedAtString,
		NodeID:      &gist.NodeID,
	}
}

func userToGraph(user *service_model.User) *graph_model.User {
	return &graph_model.User{
		Name:      &user.Name,
		Email:     &user.Email,
		Login:     &user.Login,
		AvatarURL: &user.AvatarURL,
		URL:       &user.URL,
	}
}

func filesToGraph(files []*service_model.File) []*graph_model.File {
	result := make([]*graph_model.File, 0, len(files))

	for _, file := range files {
		result = append(result, fileToGraph(file))
	}

	return result
}

func fileToGraph(file *service_model.File) *graph_model.File {
	return &graph_model.File{
		FileName: &file.FileName,
		Type:     &file.Type,
		Language: &file.Language,
		RawURL:   &file.RawURL,
		Size:     &file.Size,
	}
}

func commentsToGraph(comments []*service_model.Comment) []*graph_model.GistComment {
	result := make([]*graph_model.GistComment, 0, len(comments))

	for _, comment := range comments {
		result = append(result, commentToGraph(comment))
	}

	return result
}

func commentToGraph(comment *service_model.Comment) *graph_model.GistComment {
	id := int(comment.ID)
	createdAtString := comment.CreatedAt.String()

	return &graph_model.GistComment{
		ID:        &id,
		URL:       &comment.URL,
		Body:      &comment.Body,
		User:      userToGraph(comment.User),
		CreatedAt: &createdAtString,
	}
}

func commitsToGraph(commits []*service_model.Commit) []*graph_model.GistCommit {
	result := make([]*graph_model.GistCommit, 0, len(commits))

	for _, commit := range commits {
		result = append(result, commitToGraph(commit))
	}

	return result
}

func commitToGraph(commit *service_model.Commit) *graph_model.GistCommit {
	committedAtString := commit.CommittedAt.String()

	return &graph_model.GistCommit{
		URL:          &commit.URL,
		Version:      &commit.Version,
		User:         userToGraph(commit.User),
		ChangeStatus: changeStatusToGraph(commit.ChangeStatus),
		CommittedAt:  &committedAtString,
		NodeID:       &commit.NodeID,
	}
}

func changeStatusToGraph(changeStatus *service_model.CommitStats) *graph_model.ChangeStatus {
	return &graph_model.ChangeStatus{
		Total:     &changeStatus.Total,
		Additions: &changeStatus.Additions,
		Deletions: &changeStatus.Deletions,
	}
}
