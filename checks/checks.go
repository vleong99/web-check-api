package checks

import (
	"net/http"
	"time"

	"github.com/xray-web/web-check-api/checks/store/legacyrank"
)

type Checks struct {
	Carbon      *Carbon
	LegacyRank  *LegacyRank
	LinkedPages *LinkedPages
	Rank        *Rank
	SocialTags  *SocialTags
	Tls         *Tls
}

func NewChecks() *Checks {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	return &Checks{
		Carbon:      NewCarbon(client),
		LegacyRank:  NewLegacyRank(legacyrank.NewInMemoryStore()),
		LinkedPages: NewLinkedPages(client),
		Rank:        NewRank(client),
		SocialTags:  NewSocialTags(client),
		Tls:         NewTls(client),
	}
}
