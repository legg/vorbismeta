package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mewkiz/flac"
	"github.com/mewkiz/flac/meta"
	"github.com/mewkiz/pkg/errutil"
)

func main() {
	var title = flag.String("title", "", "track title")
	var tracknumber = flag.Int("tracknumber", 0, "track number")
	var artist = flag.String("artist", "", "artist name")
	var albumartist = flag.String("albumartist", "", "album artist name")
	var album = flag.String("album", "", "album name")
	var year = flag.Int("year", 0, "release year")
	var src = flag.String("src", "", "source flac file")
	var dst = flag.String("dst", "", "destination flac file")

	flag.Parse()

	fmt.Println("Writing metadata: ", *title, *tracknumber, *artist, *albumartist, *album, *year, " to ", *dst)

	if err := enc(*src, *title, strconv.Itoa(*tracknumber), *artist, *albumartist, *album, strconv.Itoa(*year), *dst); err != nil {
		log.Fatal(err)
	}

}

// enc decodes the given source FLAC file and stores a re-encoded copy at the
// destination path.
func enc(src, title, tracknumber, artist, albumartist, album, year, dst string) error {
	// Decode FLAC file.
	stream, err := flac.ParseFile(src)
	if err != nil {
		return err
	}
	defer stream.Close()

	// Add custom vorbis comment.
	for _, block := range stream.Blocks {
		if comment, ok := block.Body.(*meta.VorbisComment); ok {
			comment.Tags = append(comment.Tags, [2]string{"TITLE", title})
			comment.Tags = append(comment.Tags, [2]string{"TRACKNUMBER", "1"})
			comment.Tags = append(comment.Tags, [2]string{"ARTIST", artist})
			comment.Tags = append(comment.Tags, [2]string{"ALBUMARTIST", albumartist})
			comment.Tags = append(comment.Tags, [2]string{"ALBUM", album})
			comment.Tags = append(comment.Tags, [2]string{"YEAR", year})
		}
	}

	// Encode FLAC file.
	f, err := os.Create(dst)
	if err != nil {
		return errutil.Err(err)
	}
	defer f.Close()
	if err := flac.Encode(f, stream); err != nil {
		return errutil.Err(err)
	}

	return nil
}
