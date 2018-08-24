package chapter

import (
	log "github.com/sirupsen/logrus"

	api "github.com/arkhaix/lit-reader/api/chapter"
)

func (s *Server) fetchChapterFromDb(storyID string, chapterID int) (*api.Chapter, error) {
	row := s.DB.QueryRow("SELECT Url,Title,Html FROM chapter WHERE Story = $1 AND Id = $2", storyID, chapterID)

	var url, title, html string
	err := row.Scan(&url, &title, &html)

	return &api.Chapter{
		StoryId:   storyID,
		ChapterId: int32(chapterID),
		Url:       url,
		Title:     title,
		Html:      html,
	}, err
}

func (s *Server) saveChapterToDb(chapter *api.Chapter) error {
	var err error
	_, err = s.DB.Exec("UPSERT INTO chapter (Story,Id,Url,Title,Html) VALUES ($1,$2,$3,$4,$5)",
		chapter.GetStoryId(), chapter.GetChapterId(), chapter.GetUrl(), chapter.GetTitle(), chapter.GetHtml())

	if err != nil {
		log.Errorf("Error inserting chapter into db: %s", err.Error())
	}

	return err
}
