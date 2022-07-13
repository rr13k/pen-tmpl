package user_servers

import (
	"test2/internal/app/models"
	"test2/internal/app/servers"
)

type User struct {
	servers.Pagination
	UserId     int `json:"user_id"`
	Day        *int
	ProjectIds []int `json:"project_ids"`

	Id   *int    `json:"id"`
	Name *string `json:"name"`
}

func (c *User) Get() (*models.User, error) {
	return models.GetUser(c.Id)
}

func (u *User) GetAll() ([]*models.User, error) {
	return models.GetAllUser(u.GetCurrentPage(), u.GetPageSize(), u.getMaps())
}

func (c *User) Count() (int, error) {
	return models.GetUserTotal(c.getMaps())
}

// 自定义过滤规则
func (c *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["yn"] = 1

	if c.Name != nil {
		maps["name like ?"] = "%" + *c.Name + "%"
	}

	if len(c.ProjectIds) > 0 {
		maps["project_id in (?)"] = c.ProjectIds
	}

	if c.Day != nil {
		maps["to_days( NOW( ) ) - to_days(cases.`created_date`) = ? "] = c.Day
	}
	return maps
}
