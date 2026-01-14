package services

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
	"github.com/joho/godotenv"
)

func CreateVNPAYPaymentURL(orderID uint, amount float64, ipAddr string) string {
	_ = godotenv.Load()
	vnp_TmnCode := os.Getenv("VNPAY_TMN_CODE")
	vnp_HashSecret := os.Getenv("VNPAY_HASH_SECRET")
	vnp_URL := os.Getenv("VNPAY_URL")
	vnp_ReturnUrl := os.Getenv("VNPAY_RETURN_URL")

	vnp_Params := make(map[string]string)
	vnp_Params["vnp_Version"] = "2.1.0"
	vnp_Params["vnp_Command"] = "pay"
	vnp_Params["vnp_TmnCode"] = vnp_TmnCode
	vnp_Params["vnp_Amount"] = fmt.Sprintf("%d", int(amount*100))
	vnp_Params["vnp_CurrCode"] = "VND"
	vnp_Params["vnp_TxnRef"] = fmt.Sprintf("%d", orderID)
	vnp_Params["vnp_OrderInfo"] = fmt.Sprintf("Thanh toan don hang %d", orderID)
	vnp_Params["vnp_OrderType"] = "other"
	vnp_Params["vnp_Locale"] = "vn"
	vnp_Params["vnp_ReturnUrl"] = vnp_ReturnUrl
	vnp_Params["vnp_IpAddr"] = ipAddr
	vnp_Params["vnp_CreateDate"] = time.Now().Format("20060102150405")

	// 1. Sắp xếp các key theo thứ tự bảng chữ cái
	keys := make([]string, 0, len(vnp_Params))
	for k := range vnp_Params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 2. Xây dựng chuỗi signData (để băm) và query (để tạo URL)
	var signData strings.Builder
	var query strings.Builder

	for i, k := range keys {
		val := vnp_Params[k]
		// Quan trọng: VNPAY 2.1.0 yêu cầu encode cả giá trị trong chuỗi băm
		// Sử dụng url.QueryEscape và thay thế các ký tự không khớp chuẩn VNPAY
		encodedKey := url.QueryEscape(k)
		encodedVal := url.QueryEscape(val)

		if i > 0 {
			signData.WriteString("&")
			query.WriteString("&")
		}

		signData.WriteString(encodedKey + "=" + encodedVal)
		query.WriteString(encodedKey + "=" + encodedVal)
	}

	// 3. Tạo chữ ký HMAC-SHA512
	h := hmac.New(sha512.New, []byte(vnp_HashSecret))
	h.Write([]byte(signData.String()))
	vnpSecureHash := hex.EncodeToString(h.Sum(nil))

	// 4. Trả về URL cuối cùng kèm mã băm
	return vnp_URL + "?" + query.String() + "&vnp_SecureHash=" + vnpSecureHash
}