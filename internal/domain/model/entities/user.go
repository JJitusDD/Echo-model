package entities

import "time"

type User struct {
	Id                string    `json:"id" bson:"_id,omitempty" gorm:"id"`
	UserId            string    `json:"user_id" bson:"user_id,omitempty" gorm:"user_id"`
	PhoneNumber       string    `json:"phone_number" bson:"phone_number,omitempty" gorm:"phone_number"`
	Email             string    `json:"email,omitempty" bson:"email,omitempty" gorm:"email"`
	Name              string    `json:"name,omitempty" bson:"name,omitempty" gorm:"name"`
	Avatar            string    `json:"avatar,omitempty" bson:"avatar,omitempty" gorm:"avatar"`
	Gender            string    `json:"gender,omitempty" bson:"gender,omitempty" gorm:"gender"`
	Status            string    `json:"status,omitempty" bson:"status,omitempty" gorm:"status"`
	Timezone          string    `json:"timezone,omitempty" gorm:"timezone"`
	Language          string    `json:"language,omitempty" gorm:"language"`
	Title             string    `json:"title,omitempty" gorm:"title"`
	DateTime          int64     `json:"date_time,omitempty" bson:"date_time,omitempty" gorm:"date_time"`
	Currency          string    `json:"currency,omitempty" bson:"currency" gorm:"currency"`
	IdentityNumber    string    `json:"identity_number,omitempty" bson:"identity_number" gorm:"identity_number"`
	BirthDay          string    `json:"birth_day,omitempty" bson:"birth_day" gorm:"birth_day"`
	Address           string    `json:"address,omitempty" bson:"address" gorm:"address"`
	DeviceIdLogin     string    `json:"device_id_login,omitempty" bson:"device_id_login" gorm:"device_id_login"`
	LastDeviceIdLogin string    `json:"last_device_id_login,omitempty" bson:"last_device_id_login" gorm:"last_device_id_login"`
	Password          string    `json:"-" bson:"password" gorm:"password"`
	CreatedAt         time.Time `json:"created_at,omitempty" bson:"created_at" gorm:"created_at"`
	UpdatedAt         time.Time `json:"updated_at,omitempty" bson:"updated_at" gorm:"updated_at"`
	PinCode           string    `json:"-" gorm:"pin_code"`
	PinCodeExpired    int64     `json:"-" gorm:"pin_code_expired"`
	LockTimeExpired   int64     `json:"-" gorm:"lock_time_expired"`
	WrongPinCode      int64     `json:"-" gorm:"wrong_pin_code"`
	UserIdentityID    string    `json:"user_identity_id,omitempty" bson:"user_identity_id,omitempty" gorm:"user_identity_id"`
	MerchantId        string    `json:"merchant_id"`
	LinkMcLockExpTime int64     `json:"link_mc_lock_exp_time"`
	TerminalId        string    `json:"terminal_id"`
}
