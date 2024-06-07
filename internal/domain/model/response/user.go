package response

import "echo-model/internal/domain/model/entities"

type UserRespPages struct {
	User     []*entities.User    `json:"customer_sub_pkg" mapstructure:"customer_sub_pkg" `
	Pageable *PaginationResponse `json:"pageable" mapstructure:"pageable"`
}
