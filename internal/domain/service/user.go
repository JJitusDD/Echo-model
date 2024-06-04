package service

import "github.com/labstack/echo/v4"

// @Param request body request.LoginReq true "query params"
// @Success 200 {object} response.Response
// @Router /user/login [post]
// @tags User
func (s *Service) UserLogin(c echo.Context) (err error) {
	//body := new(request.LoginReq)
	//if err := utils.BindingBody(body, c); err != nil {
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
	//if err := utils.BindingBody(body, c); err != nil {
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
	//if err := utils.BindingBody(body, c); err != nil {
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
	//if err := utils.BindingBody(body, c); err != nil {
	//	return err
	//}

	return
}
