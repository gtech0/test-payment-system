package controller

import (
	"net/http"
	"test-payment-system/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WalletController struct {
	walletService *service.WalletService
}

func NewWalletController() *WalletController {
	return &WalletController{
		walletService: service.NewWalletService(),
	}
}

// GetBalance godoc
// @Tags         Wallet
// @Summary      Get Balance
// @Description  Get wallet balance by id
// @Produce      json
// @Param   	 address path string true "Wallet id"
// @Success      200 {object} dto.WalletBalanceDto
// @Failure      400 {object} errs.CustomError
// @Router       /wallet/{address}/balance [get]
func (c *WalletController) GetBalance(ctx *gin.Context) {
	addressStr := ctx.Param("address")
	address, err := uuid.Parse(addressStr)
	if err != nil {
		ctx.Error(err)
		return
	}

	balance, err := c.walletService.GetBalance(address)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, balance)
}
