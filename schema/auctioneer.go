package schema

import (
	"bidding/pkg/errors"
	"strings"
)

type Auctioneer struct {
}

type AuctionReq struct {
	AuctionID string `json:"auction_id"`
}

func (b *AuctionReq) Ok() error {
	switch {
	case strings.TrimSpace(b.AuctionID) == "":
		return errors.IsRequiredErr("auction id")
	}

	return nil
}
