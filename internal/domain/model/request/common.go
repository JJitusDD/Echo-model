package request

type Pagination struct {
	PerPage int64  `json:"per_page" mapstructure:"per_page" gorm:"per_page"`
	Page    int64  `json:"page" mapstructure:"page" gorm:"page"`
	Offset  int64  `json:"offset" mapstructure:"offset" gorm:"offset"`
	Order   string `json:"order" mapstructure:"order" gorm:"order"`
}
