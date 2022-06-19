package crons

import (
	"github.com/MahmoudMekki/ChatSystem/pkg/repo/chatDAL"
	"github.com/rs/zerolog/log"
)

func UpdateChatMessagesCount() {
	log.Info().Msg("updating chat messages count has been started ....")
	err := chatDAL.UpdateChatMessagesCount()
	if err != nil {
		log.Err(err).Msg(err.Error())
		return
	}
	log.Info().Msg("updating chat messages count has finished successfully....")
}
