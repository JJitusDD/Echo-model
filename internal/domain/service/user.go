package service

import (
	"net/http"
	"strings"

	"echo-model/internal/domain/model/request"
	"echo-model/internal/domain/model/response"
	"echo-model/internal/domain/utilties"
	helpers "echo-model/pkg/helper"
	"github.com/labstack/echo/v4"
)

// @Param request body request.LoginReq true "query params"
// @Success 200 {object} response.Response
// @Router /user/login [post]
// @tags User
func (s *Service) UserLogin(c echo.Context) (err error) {
	//body := new(request.LoginReq)
	//if err := utilities.BindingBody(body, c); err != nil {
	//	return err
	//}

	return
}

// @Param request body request.LogoutReq true "query params"
// @Success 200 {object} response.Response
// @Router /user/logout [post]
// @tags User
func (s *Service) UserLogout(c echo.Context) (err error) {
	//body := new(request.LogoutReq)
	//if err := utilities.BindingBody(body, c); err != nil {
	//	return err
	//}

	return
}

// @Param request body request.RefreshTokenReq true "query params"
// @Success 200 {object} response.Response
// @Router /user/refresh-token [post]
// @tags User
func (s *Service) UserRefreshToken(c echo.Context) (err error) {
	//body := new(request.RefreshTokenReq)
	//if err := utilities.BindingBody(body, c); err != nil {
	//	return err
	//}

	return
}

// @Param request body request.UserProfileReq true "query params"
// @Success 200 {object} response.Response
// @Router /user/profile [post]
// @tags User
func (s *Service) UserProfile(c echo.Context) (err error) {
	//body := new(request.UserProfileReq)
	//if err := utilities.BindingBody(body, c); err != nil {
	//	return err
	//}

	return
}

// @Param request body request.UserSearchReq true "query params"
// @Success 200 {object} response.Response
// @Router /user/search [post]
// @tags User
func (s *Service) UserSearch(c echo.Context) (err error) {
	body := new(request.UserSearchReq)
	if err := utilties.BindingBody(body, c); err != nil {
		return err
	}

	var reqCollection map[string]interface{}
	if body.Customer != nil {
		reqCollection = helpers.ExtractFiltersWithPrefix(*body.Customer, "customers")
		ids := s.User.FindCustomerID(c.Request().Context(), reqCollection)
		if len(ids) == 0 {
			return response.NewResponseSuccess(c, nil)
		}
		query := strings.Split(ids, ",")
		if body.CustomerIdIN != nil {
			customerIn := helpers.GetDuplicateStr(*body.CustomerIdIN, query)
			body.CustomerIdIN = &customerIn
		} else {
			body.CustomerIdIN = &query
		}
	}
	body.Customer = nil
	pagging := helpers.GetPagingRequest(body.Pagination)
	filter := helpers.ExtractFiltersNew(*body)

	result, resPage, err := s.User.FindAll(c.Request().Context(), filter, pagging)
	if err != nil {
		s.Logger.WithError(err).WithField("filter", filter).Error("SearchBillMaster::IBillMaster.FindAll")
		return response.NewResponseError(c, err, http.StatusBadRequest, nil)
	}

	return response.NewResponseSuccess(c, response.UserRespPages{
		User:     result,
		Pageable: resPage,
	})
}
