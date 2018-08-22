package story

import (
	"strings"

	log "github.com/sirupsen/logrus"

	apistory "github.com/arkhaix/lit-reader/api/story"
)

func (s *Server) queryStoryByURL(url string) (string, error) {
	url = fixURL(url)

	row := s.DB.QueryRow("SELECT Id FROM story WHERE Url = $1", url)

	var id string
	err := row.Scan(&id)
	if err != nil {
		log.Errorf("Error querying database: %s", err.Error())
	}

	return id, err
}

func (s *Server) fetchStoryFromDb(id string) (*apistory.Story, error) {
	row := s.DB.QueryRow("SELECT Url,Title,Author,NumChapters FROM story WHERE Id = $1", id)

	var url, title, author string
	var numChapters int
	err := row.Scan(&url, &title, &author, &numChapters)

	return &apistory.Story{
		Id:          id,
		Url:         url,
		Title:       title,
		Author:      author,
		NumChapters: int32(numChapters),
	}, err
}

func (s *Server) saveStoryToDb(story *apistory.Story) (string, error) {
	story.Url = fixURL(story.Url)

	id := story.GetId()
	var err error

	if len(story.GetId()) > 0 {
		_, err = s.DB.Exec("UPSERT INTO story (Id,Url,Title,Author,NumChapters) VALUES ($1,$2,$3,$4,$5)",
			id, story.GetUrl(), story.GetTitle(), story.GetAuthor(), int(story.GetNumChapters()))
	} else {
		row := s.DB.QueryRow("UPSERT INTO story (Url,Title,Author,NumChapters) VALUES ($1,$2,$3,$4) RETURNING Id",
			story.GetUrl(), story.GetTitle(), story.GetAuthor(), int(story.GetNumChapters()))
		err = row.Scan(&id)
	}

	if err != nil {
		log.Errorf("Error inserting into db: %s", err.Error())
	}

	return id, err
}

func fixURL(url string) string {
	protocolIndex := strings.Index(url, "://")
	if protocolIndex < 0 {
		return url
	}
	return url[protocolIndex+3:]
}
