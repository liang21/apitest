package mock

import (
	"apitest/apitest/mock/result"
	"github.com/labstack/echo"
)

type User struct {
	Person
}
type Person struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")
	u := &User{
		Person: Person{
			Username: username,
			Password: password,
		},
	}
	return result.Success(c, "添加成功新用户", u)
}
func UpdateUser(c echo.Context) error {
	name := c.Param("name")
	username := c.FormValue("username")
	password := c.FormValue("password")
	u := &User{
		Person: Person{
			Username: username,
			Password: password,
		},
	}
	if name == username {
		return result.Failed(c, "更新失败，该用户已存在")
	}
	return result.Success(c, "更新成功新用户", u)
}
func GetUser(c echo.Context) error {
	u := &User{Person: Person{Username: "Jon", Password: "jon@labstack.com"}}
	username := c.Param("name")
	if username != u.Person.Username {
		result.Failed(c, "查询失败，该用户不存在！")
	} else {
		return result.Success(c, "查询成功", u)
	}
	return nil

}

func DeleteUser(c echo.Context) error {
	u := &User{Person: Person{Username: "Jon", Password: "jon@labstack.com"}}
	username := c.Param("name")
	if username != u.Username {
		return result.Failed(c, "删除失败，没有该用户")
	}
	return result.Success(c, "查询成功", u)
}
