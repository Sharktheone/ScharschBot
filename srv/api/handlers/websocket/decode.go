package websocket

import (
	"Scharsch-Bot/discord/bot"
	"Scharsch-Bot/srv/playersrv"
	"Scharsch-Bot/types"
	"Scharsch-Bot/whitelist/whitelist"
	"fmt"
	"strings"
)

var (
	s = bot.Session
)

func (h *Handler) DecodePlayer(e *types.WebsocketEvent) (*PSRVEvent, error) {
	var (
		errMsg error
		pSrv   = &PSRVEvent{
			h:       h,
			e:       e,
			session: s,
		}
	)
	if e.Data.Player == "" {
		return pSrv, nil
	}
	if userID, onWhitelist := whitelist.GetOwner(e.Data.Player); onWhitelist {
		pSrv.onWhitelist = &onWhitelist
		pSrv.userID = &userID
		if member, err := s.GetUserProfile(userID); err != nil {
			errMsg = fmt.Errorf("failed to get user profile: %v", err)
			pSrv.footerIcon = &config.Discord.EmbedErrorIcon
			pSrv.username = &e.Data.Player
			pSrv.member = member
		} else {
			a := member.User.AvatarURL("40")
			u := member.User.String()
			pSrv.footerIcon = &a
			pSrv.username = &u
			pSrv.member = member
		}
	} else {
		pSrv.onWhitelist = &onWhitelist
		pSrv.footerIcon = &config.Discord.EmbedErrorIcon
	}
	playersrv.CheckAccount(strings.ToLower(e.Data.Player))

	return pSrv, errMsg
}
