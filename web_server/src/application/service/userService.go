package appservice

import (
	ctrldto "aprendendo/web_server/src/controller/dto"
	entity "aprendendo/web_server/src/domain/entities"
	repo "aprendendo/web_server/src/repository"
	postResp "aprendendo/web_server/src/repository/responses"
)

func UserPostService(dto *ctrldto.UserPostDTO) (*postResp.PostResponse, error) {
	post := &entity.Post{
		Title: dto.Title,
		Body:  dto.Body,
	}
	// pegar retorno aqui
	newPost, err := repo.CreatePost(post)

	if err != nil {
		return nil, err
	}

	return newPost, nil
}
