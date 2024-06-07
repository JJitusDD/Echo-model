package request

type UserSearchReq struct {
	ID            *string     `json:"id,omitempty"`
	CustomerIdIN  *[]string   `json:"customer_id_in,omitempty"`
	PackageId     *string     `json:"package_id_in,omitempty"`
	ExpiredDate   *int64      `json:"expired_date,omitempty"`
	StartDate     *string     `json:"start_date,omitempty"`
	Status        *[]int      `json:"status_in,omitempty"`
	FromDate      *int64      `json:"from_date,omitempty" `
	ToDate        *int64      `json:"to_date,omitempty"`
	PaymentStatus *int        `json:"payment_status,omitempty"`
	Pagination    *Pagination `json:"pagination,omitempty"`
	Customer      *Customer   `json:"customer"`
}

type Customer struct {
	MerchantID *string `json:"merchant_id"`
}
