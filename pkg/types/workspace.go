package types

type GitUserData struct {
	Name  string
	Email string
} // @name GitUserData

type Repository struct {
	Url         string       `json:"url"`
	Branch      string       `json:"branch,omitempty"`
	Sha         string       `json:"sha"`
	Owner       string       `json:"owner"`
	PrNumber    uint32       `json:"prNumber,omitempty"`
	Source      string       `json:"source"`
	Path        string       `json:"path,omitempty"`
	GitUserData *GitUserData `json:"-"`
} // @name Repository

type Project struct {
	Name        string      `json:"name"`
	Repository  *Repository `json:"repository"`
	WorkspaceId string      `json:"workspaceId"`
	ApiKey      string      `json:"-"`
} // @name Project

type WorkspaceProvider struct {
	Name    string `json:"name"`
	Profile string `json:"profile"`
} // @name WorkspaceProvider

type Workspace struct {
	Id       string             `json:"id"`
	Name     string             `json:"name"`
	Projects []*Project         `json:"projects"`
	Provider *WorkspaceProvider `json:"provider"`
} // @name Workspace

type ProjectInfo struct {
	Name             string `json:"name"`
	Created          string `json:"created"`
	Started          string `json:"started"`
	Finished         string `json:"finished"`
	IsRunning        bool   `json:"isRunning"`
	ProviderMetadata string `json:"providerMetadata,omitempty"`
	WorkspaceId      string `json:"workspaceId"`
} // @name ProjectInfo

type WorkspaceInfo struct {
	Name             string         `json:"name"`
	Projects         []*ProjectInfo `json:"projects"`
	ProviderMetadata string         `json:"providerMetadata,omitempty"`
} // @name WorkspaceInfo