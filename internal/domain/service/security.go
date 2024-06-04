package service

import (
	"bytes"
	configs "echo-model/config"
	"echo-model/pkg/helper/crypt"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"sort"
	"time"
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

func  (s *Service) ValidateAccesstoken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			user, err := crypt.DecryptAccessToken(c)

			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}
			//Parse Data token aes256-cbc
			userinfo, err := crypt.Decrypt(user.Payload, us.Config.HashAccessToken)

			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}

			var userData crypt.UserToken

			err = json.Unmarshal([]byte(userinfo), &userData)

			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}

			// Checking User
			userdetail, err := s.GetProfile(helpers.ContextWithTimeOut(), userData.UserId)

			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}

			timeLayout := "2006-01-02T15:04:05Z07:00"
			timeLockPin, err := time.Parse(timeLayout, userdetail.UserEntityV2.LoginInfo.PinLockExpDate)
			if err == nil {
				if timeLockPin.Unix() > helpers.GetCurrentTime().Unix() {
					return echo.NewHTTPError(http.StatusBadGateway, fmt.Errorf(configs.MsgAccountLock))
				}
			}

			token := c.Request().Header.Get("Authorization")[7:]



			s.Logger.WithFields(logrus.Fields{map[string]{"_id": userData.UserId}}).
				WithFields(logrus.Fields{map[string]{"_id": userData.UserId}}).Info("userDetailInfo")

			c.Set("user-detail", userdetail)
			c.Set("user-id", userData.UserId)
			c.Set("device-id", userData.DeviceId)
			c.Set("user-id-new", userdetail.User.UserId)
			c.Set("user-role", userdetail.Role)
			c.Set("phone_number", userdetail.PhoneNumber)
			c.Set("identify_number", userdetail.IdentityNumber)
			c.Set("secret", userData.Secret)
			c.Set("access_token", token)
			c.Set("customer_ref_id", userdetail.CustomerRefId)
			c.Set("user-name", userdetail.Name)
			c.Set("user_ref_id", userdetail.UserRefID)

			bUserV2, err := json.Marshal(userdetail.UserEntityV2)
			if err == nil {
				c.Set("user-entity", string(bUserV2))
			}

			return next(c)
		}
	}
}