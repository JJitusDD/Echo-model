package helpers

import (
	"encoding/base64"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/dongri/emv-qrcode/emv/mpm"
	"github.com/skip2/go-qrcode"
)

func JSONMarshalString(v interface{}) string {
	result := "{}"
	b, err := json.Marshal(v)
	if err != nil {
		return result
	}
	result = string(b)
	return result
}

func GenerateQRImage(code string) string {
	png, err := qrcode.Encode(code, qrcode.Medium, 572)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(png)
}

func GenerateQRString(accNo, msg, bankCode string, amount int64) string {
	emvqr := new(mpm.EMVQR)
	qrType := "11" // 11 is static qrcode
	if amount > 0 {
		qrType = "12" // 12 is dynamic qrcode
	}
	emvqr.SetPayloadFormatIndicator("01")
	emvqr.SetPointOfInitiationMethod(qrType) // 11 is static qrcode

	merchantAccountInformationMaster := new(mpm.MerchantAccountInformation)
	merchantAccountInformationMaster.SetGloballyUniqueIdentifier("A000000727")
	lenAccountNumber := strconv.Itoa(len(accNo))
	paymentNetworkSpecificTag1 := "0006" + bankCode + "01" + lenAccountNumber + accNo
	merchantAccountInformationMaster.AddPaymentNetworkSpecific("01", paymentNetworkSpecificTag1)
	merchantAccountInformationMaster.AddPaymentNetworkSpecific("02", "QRIBFTTA")
	emvqr.AddMerchantAccountInformation(mpm.ID("38"), merchantAccountInformationMaster)
	emvqr.SetTransactionCurrency("704")
	emvqr.SetCountryCode("VN")
	if amount > 0 {
		amountStr := strconv.FormatInt(amount, 10)
		emvqr.SetTransactionAmount(amountStr)
	}

	additionalTemplate := new(mpm.AdditionalDataFieldTemplate)
	if strings.TrimSpace(msg) != "" {
		msg = ProcessFullName(msg)
		if len(msg) > 20 {
			msg = msg[0:20]
		}
		additionalTemplate.SetPurposeTransaction(msg)
	}
	emvqr.SetAdditionalDataFieldTemplate(additionalTemplate)
	code, err := mpm.Encode(emvqr)
	if err != nil {
		return ""
	}
	return code
}

func ProcessFullName(s string) string {
	reg, _ := regexp.Compile("[^a-zA-Z ]+")
	processedString := strings.TrimSpace(reg.ReplaceAllString(strings.ToUpper(RemoveAccentsVietnamese(s)), ""))
	space := regexp.MustCompile(`\s+`)
	processedString = space.ReplaceAllString(processedString, " ")
	return strings.TrimSpace(processedString)
}

func RemoveAccentsVietnamese(str string) string {
	var Regexp_a = `à|á|ạ|ã|ả|ă|ắ|ằ|ẳ|ẵ|ặ|â|ấ|ầ|ẩ|ẫ|ậ`
	var Regexp_A = `À|Á|Ạ|Ã|Ả|Ă|Ắ|Ằ|Ẳ|Ẵ|Ặ|Â|Ấ|Ầ|Ẩ|Ẫ|Ậ`
	var Regexp_e = `è|ẻ|ẽ|é|ẹ|ê|ề|ể|ễ|ế|ệ`
	var Regexp_E = `È|Ẻ|Ẽ|É|Ẹ|Ê|Ề|Ể|Ễ|Ế|Ệ`
	var Regexp_i = `ì|ỉ|ĩ|í|ị`
	var Regexp_I = `Ì|Ỉ|Ĩ|Í|Ị`
	var Regexp_u = `ù|ủ|ũ|ú|ụ|ư|ừ|ử|ữ|ứ|ự`
	var Regexp_U = `Ù|Ủ|Ũ|Ú|Ụ|Ư|Ừ|Ử|Ữ|Ứ|Ự`
	var Regexp_y = `ỳ|ỷ|ỹ|ý|ỵ`
	var Regexp_Y = `Ỳ|Ỷ|Ỹ|Ý|Ỵ`
	var Regexp_o = `ò|ỏ|õ|ó|ọ|ô|ồ|ổ|ỗ|ố|ộ|ơ|ờ|ở|ỡ|ớ|ợ`
	var Regexp_O = `Ò|Ỏ|Õ|Ó|Ọ|Ô|Ồ|Ổ|Ỗ|Ố|Ộ|Ơ|Ờ|Ở|Ỡ|Ớ|Ợ`
	var Regexp_d = `đ`
	var Regexp_D = `Đ`
	reg_a := regexp.MustCompile(Regexp_a)
	reg_A := regexp.MustCompile(Regexp_A)

	reg_e := regexp.MustCompile(Regexp_e)
	reg_E := regexp.MustCompile(Regexp_E)

	reg_i := regexp.MustCompile(Regexp_i)
	reg_I := regexp.MustCompile(Regexp_I)

	reg_o := regexp.MustCompile(Regexp_o)
	reg_O := regexp.MustCompile(Regexp_O)

	reg_u := regexp.MustCompile(Regexp_u)
	reg_U := regexp.MustCompile(Regexp_U)

	reg_y := regexp.MustCompile(Regexp_y)
	reg_Y := regexp.MustCompile(Regexp_Y)

	reg_d := regexp.MustCompile(Regexp_d)
	reg_D := regexp.MustCompile(Regexp_D)

	str = reg_a.ReplaceAllLiteralString(str, "a")
	str = reg_A.ReplaceAllLiteralString(str, "A")

	str = reg_e.ReplaceAllLiteralString(str, "e")
	str = reg_E.ReplaceAllLiteralString(str, "E")

	str = reg_i.ReplaceAllLiteralString(str, "i")
	str = reg_I.ReplaceAllLiteralString(str, "I")

	str = reg_o.ReplaceAllLiteralString(str, "o")
	str = reg_O.ReplaceAllLiteralString(str, "O")

	str = reg_u.ReplaceAllLiteralString(str, "u")
	str = reg_U.ReplaceAllLiteralString(str, "U")

	str = reg_y.ReplaceAllLiteralString(str, "y")
	str = reg_Y.ReplaceAllLiteralString(str, "Y")
	str = reg_d.ReplaceAllLiteralString(str, "d")
	str = reg_D.ReplaceAllLiteralString(str, "D")
	return str
}

// GetDuplicateStr return a slice with all common strings between s and s1
func GetDuplicateStr(s []string, s1 []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, v := range s {
		if !allKeys[v] {
			allKeys[v] = true
		}
	}

	for _, v := range s1 {
		if allKeys[v] {
			list = append(list, v)
		}
	}

	return list
}
