package activity

import (
	"time"
)

type CountCertification struct {
	Level_1 int64 `json:"level_1"`
	Level_2 int64 `json:"level_2"`
	Level_3 int64 `json:"level_3"`
	Level_4 int64 `json:"level_4"`
	Level_5 int64 `json:"level_5"`
	Level_6 int64 `json:"level_6"`
	Level_7 int64 `json:"level_7"`
}

type LatestCertfication struct {
	UserId             int64     `json:"userId" bson:"user_id"`
	OverallPassStatus  string    `json:"overall_pass_status" bson:"overall_pass_status"`
	CertificationLevel int       `json:"certification_level" bson:"certification_level"`
	ExamStartDate      time.Time `json:"exam_start_date" bson:"exam_start_date"`
	SSOUser            string    `json:"SSOUser" bson:"is_sso_user"`
}

type CertificationUsersResp struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	CertificationLevel int    `json:"certification_level"`
	ExamStartDate      string `json:"exam_start_date"`
	OverallPassStatus  string `json:"overall_pass_status"`
	EntityName         string `json:"entityName"`
}

type UserDetails struct {
	UserId  int32 `bson:"userId"`
	SSOUser bool  `bson:"SSOUser"`
}

type MostActiveUsersResp struct {
	UserName   string `json:"userName"`
	UserEmail  string `json:"userEmail"`
	DocCount   int64  `json:"document_count"`
	EntityName string `json:"entityName"`
}

type LoggedInUsers struct {
	UserName   string `json:"userName"`
	UserEmail  string `json:"userEmail"`
	LastLogin  string `json:"last_login"`
	EntityName string `json:"entityName"`
}

type EntityToIp struct {
	EntityId  int64    `json:"entityId"`
	Ip        []string `json:"ip"`
	IPEnabled string   `json:"ipEnabled"`
	IsRange   bool     `json:"isRange"`
}

type UserActivityMetrics struct {
	RegUsers     int64 `json:"reg_users"`
	RetailUsers  int64 `json:"retail_users"`
	DocsAnalyzed int64 `json:"docs_analyzed"`
}

type PaymentsResp struct {
	Email                string    `json:"email"`
	OrderID              string    `json:"order_id"`
	TransactionStartDate time.Time `json:"txn_start_datetime"`
	PaymentType          string    `json:"title"`
	Amount               string    `json:"amount"`
	PaymentStatus        string    `json:"status"`
}

type RetailUsers struct {
	UserName    string    `json:"userName"`
	UserEmail   string    `json:"userEmail"`
	DateCreated time.Time `json:"dateCreated"`
	LastLogin   string    `json:"lastLogin"`
	Amount      float64   `json:"amount"`
}

type OrderIdReq struct {
	OrderId string `json:"orderId"`
}
type UsersResp struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Login        string    `json:"login"`
	MobileNumber string    `json:"userMobileNumber"`
	RenewalDate  time.Time `json:"renewalDate"`
	UserType     string    `json:"userType"`
	EntityName   string    `json:"entityName"`
}

type EntityDetailsTemp struct {
	EntityId   int64  `json:"entityId"`
	EntityName string `json:"entityName"`
}

type TxnDetails struct {
	Email                string    `json:"email"`
	OrderID              string    `json:"orderId"`
	TransactionStartDate time.Time `json:"txnStartDate"`
	TransactionEndDate   time.Time `json:"txn_end_date"`
	PaymentType          string    `json:"paymentType"`
	Amount               string    `json:"amount"`
	PaymentStatus        string    `json:"status"`
	Description          string    `json:"description"`
	CallBackResp         string    `json:"call_back_response"`
	StatusResp           string    `json:"status_response"`
	OrderDetails         string    `json:"order_details"`
	GST                  float64   `json:"gst"`
	InvoiceNumber        string    `json:"invoice_number"`
	CountryCode          string    `json:"country_code"`
}

type OrderDetails struct {
	Email            string    `json:"email"`
	OrderID          string    `json:"orderId"`
	InvoiceNumber    string    `json:"invoice_number"`
	BankName         string    `json:"bankName"`
	TransactionDate  time.Time `json:"txnDate"`
	PrimaryAddress   Address   `json:"info1"`
	SecondaryAddress Address   `json:"info2"`
	GSTAmount        float64   `json:"gst_amount"`
	Orders           []Order   `json:"orders"`
	TaxableAmount    float64   `json:"taxable_amount"`
	Total            float64   `json:"total"`
	TotalInWords     string    `json:"total_in_words"`
	CurrencySym      string    `json:"currency_symbol"`
}

type Order struct {
	Discount      float64 `json:"discount"`
	DiscountTotal float64 `json:"discount_total"`
	Item          string  `json:"item"`
	PromoCode     string  `json:"promo_code"`
	Quanity       int64   `json:"quantity"`
	UnitPrice     float64 `json:"unit_price"`
}

type Address struct {
	GST     string `json:"gst,omitempty"`
	City    string `json:"city,omitempty"`
	Name    string `json:"name,omitempty"`
	State   string `json:"state,omitempty"`
	Address string `json:"address,omitempty"`
	Country string `json:"country,omitempty"`
	ZipCode string `json:"zip_code,omitempty"`
}

type StatusResp struct {
	Data struct {
		PaymentInstrument struct {
			Type string `json:"type"`
		} `json:"paymentInstrument"`
	} `json:"data"`
}

type EntityIpReq struct {
	EipID     int64  `json:"eipId"`
	EntityID  int64  `json:"entityId"`
	IPAddress string `json:"ipAddress"`
	IPEnabled string `json:"ipEnabled"`
	IsRange   bool   `json:"isRange"`
}

type GetIPAddress struct {
	EipID           int64   `json:"eipId"`
	EntityID        int64   `json:"entityId"`
	EntityName      string  `json:"entityName"`
	IPAddress       string  `json:"ipAddress"`
	EntityIPAddress *string `json:"entityIPAddress"`
	EnableIPLogin   string  `json:"enableIPLogin"`
	IsRange         bool    `json:"isRange"`
}
