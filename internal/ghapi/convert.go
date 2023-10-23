package ghapi

import (
	"github.com/amir/graphql-c1/internal/model"
	"github.com/google/go-github/v56/github"
)

func convertAPIToGists(gists []*github.Gist) []*model.Gist {
	result := make([]*model.Gist, 0, len(gists))

	for _, gist := range gists {
		result = append(result, convertAPIToGist(gist))
	}

	return result
}

func convertAPIToGist(g *github.Gist) *model.Gist {
	return &model.Gist{
		ID:          g.GetID(),
		Description: g.GetDescription(),
		Public:      g.GetPublic(),
		Owner:       convertAPIToUser(g.Owner),
		Files:       convertAPIToFiles(g.Files),
		Comments:    g.GetComments(),
		HTMLURL:     g.GetHTMLURL(),
		GitPullURL:  g.GetGitPullURL(),
		GitPushURL:  g.GetGitPushURL(),
		CreatedAt:   g.GetCreatedAt().Time,
		UpdatedAt:   g.GetUpdatedAt().Time,
		NodeID:      g.GetNodeID(),
	}
}

func convertAPIToUser(user *github.User) *model.User {
	return &model.User{
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		Login:     user.GetLogin(),
		AvatarURL: user.GetAvatarURL(),
		URL:       user.GetURL(),
	}
}

func convertAPIToFiles(files map[github.GistFilename]github.GistFile) []*model.File {
	result := make([]*model.File, 0, len(files))

	for _, file := range files {
		result = append(result, convertAPIToFile(file))
	}

	return result
}

func convertAPIToFile(file github.GistFile) *model.File {
	return &model.File{
		FileName: file.GetFilename(),
		Type:     file.GetType(),
		Language: file.GetLanguage(),
		RawURL:   file.GetRawURL(),
		Size:     file.GetSize(),
		Content:  file.GetContent(),
	}
}

func convertAPIToComments(comments []*github.GistComment) []*model.Comment {
	result := make([]*model.Comment, 0, len(comments))

	for _, comment := range comments {
		result = append(result, convertAPIToComment(comment))
	}

	return result
}

func convertAPIToComment(comment *github.GistComment) *model.Comment {
	user := convertAPIToUser(comment.GetUser())

	return &model.Comment{
		ID:        comment.GetID(),
		URL:       comment.GetURL(),
		Body:      comment.GetBody(),
		User:      user,
		CreatedAt: comment.GetCreatedAt().Time,
	}
}

func convertAPIToCommits(commits []*github.GistCommit) []*model.Commit {
	result := make([]*model.Commit, 0, len(commits))

	for _, commit := range commits {
		result = append(result, convertAPIToCommit(commit))
	}

	return result
}

func convertAPIToCommit(commit *github.GistCommit) *model.Commit {
	user := convertAPIToUser(commit.GetUser())

	return &model.Commit{
		URL:          commit.GetURL(),
		Version:      commit.GetVersion(),
		User:         user,
		ChangeStatus: convertAPIToChangeStatus(commit.GetChangeStatus()),
		CommittedAt:  commit.GetCommittedAt().Time,
		NodeID:       commit.GetNodeID(),
	}
}

func convertAPIToChangeStatus(changeStats *github.CommitStats) *model.CommitStats {
	return &model.CommitStats{
		Additions: changeStats.GetAdditions(),
		Deletions: changeStats.GetDeletions(),
		Total:     changeStats.GetTotal(),
	}
}
