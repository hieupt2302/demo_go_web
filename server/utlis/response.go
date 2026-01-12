package utils

import (
	"github.com/gin-gonic/gin"
)

// ApiResponse định nghĩa cấu trúc chuẩn gửi về Client
type ApiResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success trả về phản hồi thành công với dữ liệu
func Success(c *gin.Context, statusCode int, data interface{}, message string) {
	c.JSON(statusCode, ApiResponse{
		Data:    data,
		Message: message,
	})
}

// Error trả về phản hồi lỗi chuẩn
func Error(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, ApiResponse{
		Error: errorMessage,
	})
}