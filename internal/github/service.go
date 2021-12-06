package github

import (
	"context"

	"github.com/shurcooL/githubv4"
)

type GithubServicer interface {
	GetContributionsByUsername(ctx context.Context, options GetContributionsByUsernameOptions) (*Contributions, error)
	GetCurrentContributionStreakByUsername(ctx context.Context, username string) (*CurrentContributionStreak, error)
	GetLongestContributionStreakByUsername(ctx context.Context, username string) (*LongestContributionStreak, error)
	GetTotalContributionsByUsername(ctx context.Context, username string) (*TotalContribution, error)
}

type GithubService struct {
	githubClient *githubv4.Client
}

var _ GithubServicer = (*GithubService)(nil)

func NewGithubService(githubClient *githubv4.Client) *GithubService {
	return &GithubService{
		githubClient: githubClient,
	}
}
