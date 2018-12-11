package sites

import (
	"testing"

	"github.com/Girbons/comics-downloader/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestComicExtraSetup(t *testing.T) {
	comic := new(core.Comic)
	comic.URLSource = "https://www.comicextra.com/batman-2016/chapter-58/full"

	SetupComicExtra(comic)

	assert.Equal(t, "batman-2016", comic.Name)
	assert.Equal(t, "chapter-58", comic.IssueNumber)
	assert.Equal(t, 26, len(comic.Links))
}
