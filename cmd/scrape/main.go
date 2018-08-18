package main

import (
	"flag"
	"fmt"
	"sync/atomic"
	"time"

	epub "github.com/bmaupin/go-epub"
	"github.com/davecgh/go-spew/spew"

	lit "github.com/arkhaix/lit-reader/common"
	"github.com/arkhaix/lit-reader/scraper"
)

func main() {
	// Parse input
	// url := flag.String("url", "https://www.royalroad.com/fiction/5701/savage-divinity", "Story URL")
	// url := flag.String("url", "https://www.fictionpress.com/s/2961893/1/Mother-of-Learning", "Story URL")
	url := flag.String("url", "", "Story URL")
	epubFile := flag.String("out", "out.epub", "Output epub file")
	firstChapter := flag.Int("begin", 1, "First chapter to be scraped (1-based)")
	lastChapter := flag.Int("end", -1, "Last chapter to be scraped (omit this or use -1 for all chapters)")
	metadataOnly := flag.Bool("metadata", false, "If true, only print metadata, do not fetch chapters")
	flag.Parse()

	// Validate
	if !scraper.CheckStoryURL(*url) {
		fmt.Println("Unsupported story URL")
		return
	}
	if len(*epubFile) == 0 && *metadataOnly == false {
		fmt.Println("A valid output file is required")
		return
	}

	// Fetch metadata
	fmt.Println("Fetching metadata for", *url)
	story, err := scraper.FetchStoryMetadata(*url)
	if err != nil {
		fmt.Println(err)
		return
	}
	numChapters := len(story.Chapters)
	fmt.Println("url:", story.URL)
	fmt.Println("title:", story.Title)
	fmt.Println("author:", story.Author)
	fmt.Println("chapters:", numChapters)

	if *metadataOnly {
		return
	}

	// Validate chapter range
	if *lastChapter < 1 || *lastChapter > numChapters {
		*lastChapter = numChapters
	}
	if *firstChapter < 1 || *firstChapter > numChapters || *firstChapter > *lastChapter {
		fmt.Println("Invalid chapter range")
		return
	}

	// Fetch chapters
	totalChapters := *lastChapter - (*firstChapter - 1)
	fmt.Printf("Fetching %d chapters...\n", totalChapters)

	var chaptersFetched int32
	var numErrors int32
	for i := *firstChapter - 1; i < *lastChapter; i++ {
		var fetchChapter func(int)
		fetchChapter = func(chapterIndex int) {
			// Fetch the chapter
			err = scraper.FetchChapter(&story, chapterIndex)

			// Handle errors
			if err != nil {
				if scraperError, ok := err.(lit.ScraperError); ok {
					if scraperError.CanRetry() {
						fetchChapter(chapterIndex)
					} else /* not retryable */ {
						atomic.AddInt32(&numErrors, 1)
						spew.Dump(scraperError, scraperError.Err)
					}
				} else /* not a ScraperError */ {
					atomic.AddInt32(&numErrors, 1)
					spew.Dump(err)
				}
			} else /* err == nil */ {
				atomic.AddInt32(&chaptersFetched, 1)
			}
		}

		ii := i // save i to prevent a race
		go fetchChapter(ii)
	}

	sleepDurationMillis := time.Duration(100)
	updateDelayMillis := time.Duration(3000)
	lastUpdateTime := time.Now()
	for {
		fetched := atomic.LoadInt32(&chaptersFetched) + atomic.LoadInt32(&numErrors)
		if fetched >= int32(totalChapters) {
			break
		}
		if time.Since(lastUpdateTime) > (time.Millisecond * updateDelayMillis) {
			lastUpdateTime = time.Now()
			fmt.Printf("[%d%%] Fetched %d of %d chapters...\n",
				int((float32(fetched)/float32(totalChapters))*100), chaptersFetched, totalChapters)
		}
		time.Sleep(sleepDurationMillis * time.Millisecond)
	}

	fmt.Printf("Finished fetching %d chapters (%d errors)\n", totalChapters, numErrors)

	// Write the epub
	fmt.Println("Writing epub...")
	e := epub.NewEpub(story.Title)
	e.SetAuthor(story.Author)
	for i := *firstChapter - 1; i < *lastChapter; i++ {
		chapter := &story.Chapters[i]
		e.AddSection(chapter.HTML, chapter.Title, "", "")
	}
	e.Write(*epubFile)

	fmt.Println("Done")
}
