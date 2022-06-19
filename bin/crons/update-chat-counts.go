package crons

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/appDAL"
	"github.com/rs/zerolog/log"
)

func UpdateAppChatCount() {
	log.Info().Msg("updating app chats count has been started ....")
	err := appDAL.UpdateAppChatsCount()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return
	}
	log.Info().Msg("updating app chats count has finished successfully....")
}
