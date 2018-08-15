package main

import (
	"flag"
	"fmt"

	epub "github.com/bmaupin/go-epub"

	rr "github.com/arkhaix/lit-reader/scraper/royalroad"
)

func main() {
	// Parse input
	url := flag.String("url", "https://www.royalroad.com/fiction/5701/savage-divinity", "Story URL")
	epubFile := flag.String("epub", "test.epub", "Output epub file")
	flag.Parse()

	// Validate
	if !rr.IsSupportedStoryURL(*url) {
		panic("Unsupported story URL")
	}
	if len(*epubFile) == 0 {
		panic("A valid output file is required")
	}

	// Fetch metadata
	fmt.Println("Fetching metadata for", *url)
	story, err := rr.FetchStoryMetadata(*url)
	if err != nil {
		panic(err)
	}
	fmt.Println("url:", story.URL)
	fmt.Println("title:", story.Title)
	fmt.Println("author:", story.Author)
	fmt.Println("chapters:", len(story.Chapters))

	// Fetch a chapter
	chapterIndex := len(story.Chapters) - 1
	// chapterIndex := 0
	fmt.Println("Fetching chapter", chapterIndex)
	err = rr.FetchChapter(&story, chapterIndex)
	if err != nil {
		panic(err)
	}
	chapter := &story.Chapters[chapterIndex]
	fmt.Println("url:", chapter.URL)
	fmt.Println("title:", chapter.Title)
	// fmt.Println("text:")
	// fmt.Println(chapter.Text)

	e := epub.NewEpub(story.Title)
	e.SetAuthor(story.Author)
	e.AddSection(chapter.HTML, chapter.Title, "", "")
	e.Write(*epubFile)
}
