package query

import (
	"context"
	"github.com/CATISNOTSODIUM/taggy-backend/prisma/db"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/utils"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
)


func GetUsers(currentDB * database.Database) ([]*models.User, error) {
	ctx := context.Background()
	userObjects, err := currentDB.Client.User.FindMany().Exec(ctx)

	if err != nil {
		return nil, err
	}

	users := []*models.User{}
	for _, userObject := range userObjects {
		users = append(users, &models.User {
			ID: userObject.ID,
			Name: userObject.Name,
		})
	}
	return users, nil
}

func GetUserByID(currentDB * database.Database, id string) (* models.User, error) {
	ctx := context.Background()
	userObject, err := currentDB.Client.User.FindUnique(db.User.ID.Equals(id)).Exec(ctx)

	if err != nil {
		return nil, err
	}

	user := &models.User {
		ID: userObject.ID,
		Name: userObject.Name,
	}
	return user, nil
}

func VerifyUser(currentDB * database.Database, name string, password string) (* models.User, bool) {
	ctx := context.Background()
	userObject, err := currentDB.Client.User.FindUnique(db.User.Name.Equals(name)).Exec(ctx)
	if err != nil {
		return nil,  false
	}
	if (password == string(utils.Decode(userObject.Password))) {
		return &models.User {
			ID: userObject.ID,
			Name: userObject.Name,
		}, true
	}
	return nil, false
}
