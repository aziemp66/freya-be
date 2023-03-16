package post

import (
	"context"

	postDomain "github.com/aziemp66/freya-be/internal/domain/post"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepositoryImplementation struct {
	db *mongo.Database
}

func NewPostRepositoryImplementation(db *mongo.Database) *PostRepositoryImplementation {
	return &PostRepositoryImplementation{db}
}

func (r *PostRepositoryImplementation) Insert(ctx context.Context, post postDomain.Post) (err error) {
	_, err = r.db.Collection("posts").InsertOne(ctx, post)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostRepositoryImplementation) FindByID(ctx context.Context, id string) (post postDomain.Post, err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return postDomain.Post{}, err
	}

	err = r.db.Collection("posts").FindOne(ctx, postDomain.Post{ID: objId}).Decode(&post)

	if err != nil {
		return postDomain.Post{}, err
	}

	return post, nil
}

func (r *PostRepositoryImplementation) FindAll(ctx context.Context) (posts []postDomain.Post, err error) {
	cursor, err := r.db.Collection("posts").Find(ctx, postDomain.Post{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post postDomain.Post

		err = cursor.Decode(&post)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepositoryImplementation) Delete(ctx context.Context, id string) (err error) {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = r.db.Collection("posts").DeleteOne(ctx, postDomain.Post{ID: objId})

	if err != nil {
		return err
	}

	return nil
}
