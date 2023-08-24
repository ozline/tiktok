package service

import "github.com/ozline/tiktok/kitex_gen/interaction"

func (s *InteractionService) GetUserTotalFavorited(req *interaction.UserTotalFavoritedRequest) (int64, error) {
	// TODO: video rpc call
	return 0, nil
}
