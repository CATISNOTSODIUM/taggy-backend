package query

import (
	"context"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
	"github.com/CATISNOTSODIUM/taggy-backend/prisma/db"
)


func GetThreads(currentDB * database.Database) ([]*models.Thread, error) {
	ctx := context.Background()

	// fix add pagination system

	threadObjects, err := currentDB.Client.Thread.FindMany().Take(7).OrderBy(
		db.Thread.CreatedAt.Order(db.SortOrderDesc),
	).With(
		db.Thread.Tags.Fetch().With( 
			db.TagsOnThreads.Tag.Fetch(), // get tag name
		),
	).With(
		db.Thread.User.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}


	threads := []*models.Thread{}
	for _, threadObject := range threadObjects {
		userObject := threadObject.User()
		user := models.User {
			Name: userObject.Name,
			ID: userObject.ID,
		}

		tags := [] models.Tag{}

		tagObjects := threadObject.Tags()
		for _, tagObject := range tagObjects { 
			tag := models.Tag {
				ID: tagObject.TagID,
				Name: tagObject.Tag().Name,
			}	
			tags = append(tags, tag)
		} 

		thread := models.Thread {
			ID: threadObject.ID,
			Title: threadObject.Title,
			Content: threadObject.Content,
			Likes: threadObject.Likes,
			Views: threadObject.Views,
			User: user,
			Tags: tags,
			CreatedAt: threadObject.CreatedAt,
			UpdatedAt: threadObject.UpdatedAt,
		}	
		threads = append(threads, &thread)
	}
	
	return threads, nil
}

func GetThreadByID(currentDB * database.Database, id string) (* models.Thread, error) {
	ctx := context.Background()
	threadObject, err := currentDB.Client.Thread.FindUnique(db.Thread.ID.Equals(id)).With(
		db.Thread.Tags.Fetch().With( 
			db.TagsOnThreads.Tag.Fetch(), // get tag name
		),
	).With(
		db.Thread.User.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	userObject := threadObject.User()
	user := models.User {
		Name: userObject.Name,
		ID: userObject.ID,
	}

	tags := [] models.Tag{}
	tagObjects := threadObject.Tags()
	for _, tagObject := range tagObjects { 
		tag := models.Tag {
			ID: tagObject.TagID,
			Name: tagObject.Tag().Name,
		}	
		tags = append(tags, tag)
	} 

	thread := &models.Thread {
		ID: threadObject.ID,
		Title: threadObject.Title,
		Content: threadObject.Content,
		Likes: threadObject.Likes,
		Views: threadObject.Views,
		User: user,
		Tags: tags,
		CreatedAt: threadObject.CreatedAt,
    	UpdatedAt: threadObject.UpdatedAt,
	}
	return thread, nil
}

func GetThreadTagsByID(currentDB * database.Database, id string) ([] models.Tag, error) {
	ctx := context.Background()
	tagOnThreadObjects, err := currentDB.Client.TagsOnThreads.FindMany(
		db.TagsOnThreads.ThreadID.Equals(id),
	).With(
		db.TagsOnThreads.Tag.Fetch(),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	tags := [] models.Tag{}
	
	for _, tagOnThreadObject := range tagOnThreadObjects {

		tag := models.Tag {
			ID: tagOnThreadObject.Tag().ID,
			Name: tagOnThreadObject.Tag().Name,
		}	
		tags = append(tags, tag)
	}

	return tags, nil
}