package mapper

import (
	"github.com/savi2w/wesley-chan/model"
	"github.com/savi2w/wesley-chan/presenter/res"
)

func BoardModelToRes(boards []model.Board) (resp []res.Board) {
	for _, board := range boards {
		resp = append(resp, res.Board{
			ID:          board.ID,
			Name:        board.Name,
			Slug:        board.Slug,
			Description: board.Description,
			CreatedAt:   board.CreatedAt,
			UpdatedAt:   board.UpdatedAt,
			DeletedAt:   board.DeletedAt,
		})
	}

	return resp
}
