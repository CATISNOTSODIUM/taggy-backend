package mutation

import (
	"context"
	"errors"
	"time"

	"github.com/CATISNOTSODIUM/taggy-backend/internal/database"
	"github.com/CATISNOTSODIUM/taggy-backend/internal/models"
	"github.com/CATISNOTSODIUM/taggy-backend/prisma/db"
)

func CreateThread(currentDB * database.Database, user * models.User, title string, content string) (* models.Thread, error) {
	ctx := context.Background()
	threadObject, err := currentDB.Client.Thread.CreateOne(
		db.Thread.Title.Set(title),
		db.Thread.Content.Set(content),
		db.Thread.User.Link(
			db.User.ID.Equals(user.ID),
		),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	thread := models.Thread {
		ID: threadObject.ID,
		Title: threadObject.Title,
		Content: threadObject.Content,
		Likes: threadObject.Likes,
		Views: threadObject.Views,
		User: *user,
		CreatedAt: threadObject.CreatedAt,
    	UpdatedAt: threadObject.UpdatedAt,
	}
	return &thread, nil
}

// Edit thread
func UpdateThread(currentDB * database.Database, id string, user * models.User, title string, content string) (* models.Thread, error) {
	ctx := context.Background()
	// find the thread
	threadObjectDB := currentDB.Client.Thread.FindUnique(db.Thread.ID.Equals(id))
	threadObject, err := threadObjectDB.Exec(ctx)
	if err != nil {
		return nil, err
	}
	if (threadObject.UserID != user.ID) {
		return nil, errors.New("User ID does not have permission to update the thread")
	}
	threadObject, err = threadObjectDB.Update(
		db.Thread.Title.Set(title),
		db.Thread.Content.Set(content),
		db.Thread.UpdatedAt.Set(time.Now()),
	).Exec(ctx)
	
	if err != nil {
		return nil, err
	}

	thread := models.Thread {
		ID: threadObject.ID,
		Title: threadObject.Title,
		Content: threadObject.Content,
		Likes: threadObject.Likes,
		Views: threadObject.Views,
		User: *user,
		CreatedAt: threadObject.CreatedAt,
    	UpdatedAt: threadObject.UpdatedAt,
	}
	return &thread, nil
}
// Check if you are the owner of the thread

func ViewThread(currentDB * database.Database, id string) (int, error) {
	ctx := context.Background()
	threadObject, err := 
		currentDB.Client.Prisma.ExecuteRaw(`UPDATE "Thread" SET views = views + 1 WHERE id = $1`, id).Exec(ctx)
	
		if err != nil {
		return 0, err
	}

	return threadObject.Count, nil
}

func LikeThread(currentDB * database.Database, id string) (int, error) {
	ctx := context.Background()
	threadObject, err := 
		currentDB.Client.Prisma.ExecuteRaw(`UPDATE "Thread" SET likes = likes + 1 WHERE id = $1`, id).Exec(ctx)
	
		if err != nil {
		return 0, err
	}

	return threadObject.Count, nil
}

func UnlikeThread(currentDB * database.Database, id string) (int, error) {
	ctx := context.Background()
	threadObject, err := 
		currentDB.Client.Prisma.ExecuteRaw(`UPDATE "Thread" SET likes = likes - 1 WHERE id = $1 AND likes > 0`, id).Exec(ctx)
		if err != nil {
		return 0, err
	}

	return threadObject.Count, nil
}