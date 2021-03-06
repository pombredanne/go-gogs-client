// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Permission represents a API permission.
type Permission struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

// Repository represents a API repository.
type Repository struct {
	Id          int64      `json:"id"`
	Owner       User       `json:"owner"`
	FullName    string     `json:"full_name"`
	Private     bool       `json:"private"`
	Fork        bool       `json:"fork"`
	HtmlUrl     string     `json:"html_url"`
	CloneUrl    string     `json:"clone_url"`
	SshUrl      string     `json:"ssh_url"`
	Permissions Permission `json:"permissions"`
}

// ListMyRepos lists all repositories for the authenticated user that has access to.
func (c *Client) ListMyRepos() ([]*Repository, error) {
	repos := make([]*Repository, 0, 10)
	err := c.getParsedResponse("GET", "/user/repos", nil, nil, &repos)
	return repos, err
}

type CreateRepoOption struct {
	Name        string `json:"name" binding:"Required"`
	Description string `json:"description" binding:"MaxSize(255)"`
	Private     bool   `form:"private"`
	AutoInit    bool   `form:"auto_init"`
	Gitignore   string `form:"gitignore"`
	License     string `form:"license"`
}

// CreateRepo creates a repository for authenticated user.
func (c *Client) CreateRepo(opt CreateRepoOption) (*Repository, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	repo := new(Repository)
	return repo, c.getParsedResponse("POST", "/user/repos",
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), repo)
}

// CreateOrgRepo creates an organization repository for authenticated user.
func (c *Client) CreateOrgRepo(org string, opt CreateRepoOption) (*Repository, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	repo := new(Repository)
	return repo, c.getParsedResponse("POST", fmt.Sprintf("/org/%s/repos", org),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), repo)
}
