package controller

import (
	"net/http"
	"strconv"
	"test-payment-system/dto"
	"test-payment-system/service"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	transactionService *service.TransactionService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		transactionService: service.NewTransactionService(),
	}
}

// Send godoc
// @Tags         Transaction
// @Summary      Send
// @Description  Transfer money from one wallet to the other
// @Produce      json
// @Param   	 payload body dto.SendDto false "Form data"
// @Success      200
// @Failure      400 {object} errs.CustomError
// @Router       /send [post]
func (t *TransactionController) Send(ctx *gin.Context) {
	var sendDto dto.SendDto
	if err := ctx.Bind(&sendDto); err != nil {
		ctx.Error(err)
		return
	}

	err := t.transactionService.Send(sendDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

// GetLast godoc
// @Tags         Transaction
// @Summary      Get Last
// @Description  Get last N transactions
// @Produce      json
// @Param   	 count query string true "Transaction count"
// @Success      200 {object} []dto.TransactionGetDto
// @Failure      400 {object} errs.CustomError
// @Router       /transactions [get]
func (t *TransactionController) GetLast(ctx *gin.Context) {
	countStr := ctx.Query("count")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	transactions, err := t.transactionService.GetLast(count)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}
