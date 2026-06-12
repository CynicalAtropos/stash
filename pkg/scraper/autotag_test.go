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
		Path:                    "performer name.mp4",
		PerformerAssignmentLock: true,
	})

	assert.Nil(t, err)
	assert.Nil(t, ret)
	db.AssertExpectations(t)
}

func TestAutotagScraperViaSceneSkipsLockedStudio(t *testing.T) {
	db := mocks.NewDatabase()

	db.Performer.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Performer{}, nil).Once()
	db.Performer.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Performer{}, 0, nil).Once()
	db.Tag.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Tag{}, nil).Once()
	db.Tag.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Tag{}, 0, nil).Once()

	s := autotagScraper{
		txnManager:      db,
		performerReader: db.Performer,
		studioReader:    db.Studio,
		tagReader:       db.Tag,
	}

	ret, err := s.viaScene(context.Background(), nil, &models.Scene{
		Path:                 "studio name.mp4",
		StudioAssignmentLock: true,
	})

	assert.Nil(t, err)
	assert.Nil(t, ret)
	db.AssertExpectations(t)
}

func TestAutotagScraperViaSceneSkipsLockedTags(t *testing.T) {
	db := mocks.NewDatabase()

	db.Performer.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Performer{}, nil).Once()
	db.Performer.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Performer{}, 0, nil).Once()
	db.Studio.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Studio{}, nil).Once()
	db.Studio.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Studio{}, 0, nil).Once()

	s := autotagScraper{
		txnManager:      db,
		performerReader: db.Performer,
		studioReader:    db.Studio,
		tagReader:       db.Tag,
	}

	ret, err := s.viaScene(context.Background(), nil, &models.Scene{
		Path:              "tag name.mp4",
		TagAssignmentLock: true,
	})

	assert.Nil(t, err)
	assert.Nil(t, ret)
	db.AssertExpectations(t)
}

func TestAutotagScraperViaImageSkipsLockedPerformers(t *testing.T) {
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

	ret, err := s.viaImage(context.Background(), nil, &models.Image{
		Path:                    "performer name.jpg",
		PerformerAssignmentLock: true,
	})

	assert.Nil(t, err)
	assert.Nil(t, ret)
	db.AssertExpectations(t)
}

func TestAutotagScraperViaImageSkipsLockedStudio(t *testing.T) {
	db := mocks.NewDatabase()

	db.Performer.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Performer{}, nil).Once()
	db.Performer.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Performer{}, 0, nil).Once()
	db.Tag.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Tag{}, nil).Once()
	db.Tag.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Tag{}, 0, nil).Once()

	s := autotagScraper{
		txnManager:      db,
		performerReader: db.Performer,
		studioReader:    db.Studio,
		tagReader:       db.Tag,
	}

	ret, err := s.viaImage(context.Background(), nil, &models.Image{
		Path:                 "studio name.jpg",
		StudioAssignmentLock: true,
	})

	assert.Nil(t, err)
	assert.Nil(t, ret)
	db.AssertExpectations(t)
}

func TestAutotagScraperViaImageSkipsLockedTags(t *testing.T) {
	db := mocks.NewDatabase()

	db.Performer.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Performer{}, nil).Once()
	db.Performer.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Performer{}, 0, nil).Once()
	db.Studio.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Studio{}, nil).Once()
	db.Studio.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Studio{}, 0, nil).Once()

	s := autotagScraper{
		txnManager:      db,
		performerReader: db.Performer,
		studioReader:    db.Studio,
		tagReader:       db.Tag,
	}

	ret, err := s.viaImage(context.Background(), nil, &models.Image{
		Path:              "tag name.jpg",
		TagAssignmentLock: true,
	})

	assert.Nil(t, err)
	assert.Nil(t, ret)
	db.AssertExpectations(t)
}

func TestAutotagScraperViaGallerySkipsLockedPerformers(t *testing.T) {
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

	ret, err := s.viaGallery(context.Background(), nil, &models.Gallery{
		Path:                    "performer name",
		PerformerAssignmentLock: true,
	})

	assert.Nil(t, err)
	assert.Nil(t, ret)
	db.AssertExpectations(t)
}

func TestAutotagScraperViaGallerySkipsLockedStudio(t *testing.T) {
	db := mocks.NewDatabase()

	db.Performer.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Performer{}, nil).Once()
	db.Performer.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Performer{}, 0, nil).Once()
	db.Tag.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Tag{}, nil).Once()
	db.Tag.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Tag{}, 0, nil).Once()

	s := autotagScraper{
		txnManager:      db,
		performerReader: db.Performer,
		studioReader:    db.Studio,
		tagReader:       db.Tag,
	}

	ret, err := s.viaGallery(context.Background(), nil, &models.Gallery{
		Path:                 "studio name",
		StudioAssignmentLock: true,
	})

	assert.Nil(t, err)
	assert.Nil(t, ret)
	db.AssertExpectations(t)
}

func TestAutotagScraperViaGallerySkipsLockedTags(t *testing.T) {
	db := mocks.NewDatabase()

	db.Performer.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Performer{}, nil).Once()
	db.Performer.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Performer{}, 0, nil).Once()
	db.Studio.On("QueryForAutoTag", mock.Anything, mock.Anything).Return([]*models.Studio{}, nil).Once()
	db.Studio.On("Query", mock.Anything, mock.Anything, mock.Anything).Return([]*models.Studio{}, 0, nil).Once()

	s := autotagScraper{
		txnManager:      db,
		performerReader: db.Performer,
		studioReader:    db.Studio,
		tagReader:       db.Tag,
	}

	ret, err := s.viaGallery(context.Background(), nil, &models.Gallery{
		Path:              "tag name",
		TagAssignmentLock: true,
	})

	assert.Nil(t, err)
	assert.Nil(t, ret)
	db.AssertExpectations(t)
}
