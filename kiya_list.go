package kiyaui

import (
	"context"
	"log"
	"time"

	cloudstore "cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type KiyaKey struct {
	Name string
	CreatedAt time.Time
	Owner string
}

func (k *Kiya) ListDetails(target string) (keys []KiyaKey) {
	profile, ok := k.Profiles[target]
	if !ok {
		log.Fatal("the fuck?")
	}
	ctx := context.Background()
	bucket := k.gcpStorage.Bucket(profile.Bucket)
	query := new(cloudstore.Query)
	it := bucket.Objects(ctx, query)

	for {
		next, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		keys = append(keys, KiyaKey{
			Name: next.Name,
			CreatedAt: next.Created,
			Owner: next.Owner,
		})
	}
	return keys
}
