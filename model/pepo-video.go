package model

type PepoVideo struct {
	ID                      int64    `json:"id"`
	LastModified            int64    `json:"modified_at"`
	Creator                 Creator  `json:"created_by"`
	URL                     string   `json:"url"`
	VideoURL                string   `json:"video_url"`
	TotalContributors       int64    `json:"total_contributors"`
	TotalContributionAmount string   `json:"total_contribution_amount"`
	Description             *string  `json:"description"`
	PosterImage             *string  `json:"poster_image"`
	Tags                    []string `json:"tags"`
}

type Creator struct {
	ID                 string  `json:"id"`
	Name               string  `json:"name"`
	ProfileImage       *string `json:"profile_image"`
	TokenholderAddress *string `json:"tokenholder_address"`
	TwitterHandle      *string `json:"twitter_handle"`
	GithubHandle       *string `json:"github_handle"`
}
