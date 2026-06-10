package scraper

import (
	"context"
	"testing"

	"github.com/stashapp/stash/pkg/models"
	"github.com/stashapp/stash/pkg/models/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAutotagScraperViaSceneSkipsLockedPerformers(t *testing.T) {
	db := mocks.NewDatabase()

	db.Studio.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Studio{}, nil).Once()
	db.Studio.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Studio{}, 0, nil).Once()
	db.Tag.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Tag{}, nil).Once()
	db.Tag.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Tag{}, 0, nil).Once()

	s := autotagScraper{
		txnManager:      db,
		performerReader: db.Performer,
		studioReader:    db.Studio,
		tagReader:       db.Tag,
	}

	ret, err := s.viaScene(context.Background(), nil, &models.Scene{
		Path:                 "performer name.mp4",
		PerformerAutotagLock: true,
	})

	assert.Nil(t, err)
	assert.Nil(t, ret)
	db.AssertExpectations(t)
}
