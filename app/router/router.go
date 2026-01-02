package router

import (
	"github.com/core-wallet/app/config"
	"github.com/core-wallet/app/controller"
	"github.com/core-wallet/app/handler"
	"github.com/core-wallet/app/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := config.DB
	checkHealthHandler := handler.NewCheckHealthHandler()
	checkHealthController := controller.NewCheckHealthController(checkHealthHandler)

	transactionRepo := repository.NewTransactionRepository(db)

	walletRepo := repository.NewWalletRepository(db)
	walletHandler := handler.NewWalletHandler(walletRepo)
	topUpHandler := handler.NewTopUpHandler(walletRepo, transactionRepo)

	walletController := controller.NewWalletController(walletHandler, topUpHandler)

	RegisterWalletRoutes(app, walletController)
	RegisterCheckHealthRoutes(app, checkHealthController)
}
