package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"

	"echo-model/pkg/helper/crypt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// VerifyHash check if the request was edited or not
func (s *Service) VerifyHash() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			c := make(map[string]interface{})
			var bodyBytes []byte
			if context.Request().Body != nil {
				bodyBytes, _ = io.ReadAll(context.Request().Body)
				e := json.Unmarshal(bodyBytes, &c)
				if e != nil {
					return e
				}

				if s.Config.VerifyHash {
					err := verifyhash(c, s.Config.ApiSecretKey)
					if err != nil {
						return echo.NewHTTPError(http.StatusBadRequest, err.Error())
					}
				}
			}

			// Restore the io.ReadCloser to its original state
			context.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			return next(context)
		}
	}
}

func verifyhash(c map[string]interface{}, apiSecretKey string) (e error) {
	stringHash := ""
	hashValue := ""
	var keys []string
	for s := range c {
		if s != "hash" {
			keys = append(keys, s)
		}

	}
	sort.Strings(keys)
	for _, v := range keys {
		switch c[v].(type) {
		case float64:
			if c[v].(float64) > 1000 {
				convert := c[v].(float64)
				c[v] = int64(convert)
			}
		default:
			c[v] = c[v]
		}
		str := fmt.Sprint(c[v])

		if stringHash != "" {
			stringHash = stringHash + "|" + str
		} else {
			stringHash = stringHash + str
		}

	}

	if c["hash"] != nil {
		hashValue = fmt.Sprint(c["hash"])
	}

	sha1Verify, _ := crypt.VerifyHash(stringHash+"|"+apiSecretKey, hashValue)

	if !sha1Verify {
		return
	}
	return
}

func (s *Service) ValidateAccesstoken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := crypt.DecryptAccessToken(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}
			//Parse Data token
			userinfo, err := crypt.Decrypt(user.Payload, s.Config.HashAccessToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}

			var userData crypt.UserToken
			err = json.Unmarshal([]byte(userinfo), &userData)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}

			// Checking User
			ctx, _ := context.WithTimeout(context.Background(), time.Minute*3)
			userdetail, err := s.EchoModelFacade.User.GetProfile(ctx, userData.UserId)
			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}

			token := c.Request().Header.Get("Authorization")[7:]

			s.Logger.WithFields(logrus.Fields{"_id": userData.UserId}).
				WithFields(logrus.Fields{"_user_info": userData}).Info("userDetailInfo")

			c.Set("user-detail", userdetail)
			c.Set("user-id", userData.UserId)
			c.Set("device-id", userData.DeviceId)
			c.Set("user-id-new", userData.UserId)
			c.Set("phone_number", userdetail.PhoneNumber)
			c.Set("identify_number", userdetail.IdentityNumber)
			c.Set("secret", userData.Secret)
			c.Set("access_token", token)
			c.Set("user-name", userdetail.Name)

			return next(c)
		}
	}
}
