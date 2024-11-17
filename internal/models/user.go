package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	RoleID    int    `json:"role_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ClusterID int    `json:"cluster_id"`
}
