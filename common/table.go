package common

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"gorest/entity"
)

func PrintUsers(users *[]entity.User) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("ID", "Username", "Password", "Role", "AvatarUrl", "Description", "Valid", "Recipes")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, u := range *users {
		tbl.AddRow(u.ID, u.Username, u.Password, u.Role, u.AvatarUrl, u.Description, u.Valid, u.Recipes)
	}
	tbl.Print()
}
