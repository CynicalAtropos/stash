package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/stashapp/stash/internal/api/urlbuilders"
	"github.com/stashapp/stash/internal/manager"
	"github.com/stashapp/stash/pkg/models"
)

func (r *queryResolver) SceneStreams(ctx context.Context, id *string) ([]*manager.SceneStreamEndpoint, error) {
	sceneID, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}

	// find the scene
	var scene *models.Scene
	if err := r.withReadTxn(ctx, func(ctx context.Context) error {
		var err error
		scene, err = r.repository.Scene.Find(ctx, sceneID)

		if scene != nil {
			err = scene.LoadFiles(ctx, r.repository.Scene)
		}

		return err
	}); err != nil {
		return nil, err
	}

	if scene == nil {
		return nil, fmt.Errorf("scene with id %d not found", sceneID)
	}

	config := manager.GetInstance().Config

	baseURL, _ := ctx.Value(BaseURLCtxKey).(string)
	builder := urlbuilders.NewSceneURLBuilder(baseURL, scene)

	userID, err := getSignedURLUserID(ctx, config)
	if err != nil {
		return nil, err
	}

	streamURL := builder.GetStreamURL("")
	setStreamURLAuthParams(config, userID, streamURL)

	return manager.GetSceneStreamPaths(
		scene,
		streamURL,
		func(file *models.VideoFile) *url.URL {
			u := builder.GetFileStreamURL(file.ID, "")
			setStreamURLAuthParams(config, userID, u)
			return u
		},
		config.GetMaxStreamingTranscodeSize(),
	)
}
