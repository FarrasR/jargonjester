package repository

import (
	"context"
	"fmt"
	"jargonjester/domain"
	"log"
	"time"

	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/memorystore"
)

type configRepository struct {
	store limiter.Store
}

func NewconfigRepository(token uint64, interval int) domain.ConfigRepository {
	store, err := memorystore.New(&memorystore.Config{
		Tokens:   token,
		Interval: time.Duration(interval) * time.Minute,
	})

	if err != nil {
		log.Fatal(err)
	}

	return &configRepository{
		store: store,
	}
}

func (r *configRepository) IsLimited(key string) error {

	ctx := context.Background()

	_, _, reset, ok, err := r.store.Take(ctx, key)

	if err != nil {
		return err
	}

	if !ok {
		return fmt.Errorf("rate limited: retry at %s", time.Unix(0, int64(reset)).Format(time.RubyDate))
	}

	return nil
}
