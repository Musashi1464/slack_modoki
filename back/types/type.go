package types

import "time"

type Workspace struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID          int
	WorkspaceID int
	Name        string
	Email       string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type WorkspaceForm struct {
	WorkspaceName string
	UserName      string
	Email         string
	Password      string
}

type LoginForm struct {
	WorkspaceName string
	Email         string
	Password      string
}

type LoginReturn struct {
	Success bool
	Message string
}
