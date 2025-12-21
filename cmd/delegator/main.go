package main

import (
	"delegator/internal/adapter/handler"
	"delegator/internal/application/usecase"
	"delegator/internal/infrastructure/external"
	"delegator/internal/infrastructure/persistence"
	"delegator/web"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetEnvPrefix("delegator")
	viper.AutomaticEnv()
}

func main() {
	// Echoインスタンスの初期化
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	api := e.Group("/api")

	// データベースに接続
	db, err := gorm.Open(sqlite.Open("delegator.db?_pragma=foreign_keys(1)"))
	if err != nil {
		panic(err)
	}

	// DBのマイグレーションを実行
	err = persistence.Migrate(db)
	if err != nil {
		panic(err)
	}
	// 初期データの登録
	err = persistence.Seed(db)
	if err != nil {
		panic(err)
	}

	// 静的ファイルのホスト
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "dist",
		Filesystem: http.FS(web.DistFS),
		HTML5:      true,
	}))

	// アプリケーション設定用のエンドポイントの設定
	settingsStore := persistence.NewSettingsStore(db)
	settingsUsecase := usecase.NewSettingsUsecase(settingsStore)
	settingsHandler := handler.NewSettingsHandler(settingsUsecase)
	api.GET("/settings", settingsHandler.Load)
	api.PUT("/settings", settingsHandler.Save)

	// Delegation用エンドポイントの設定
	delegationStore := persistence.NewDelegationStore(db)
	keycloakNotifier := external.NewKeycloakNotifier(http.DefaultClient, settingsStore)
	delegationUsecase := usecase.NewDelegationUsecase(delegationStore, keycloakNotifier)
	delegationHandler := handler.NewDelegationHandler(delegationUsecase)
	api.GET("/delegations", delegationHandler.List)
	api.POST("/delegations", delegationHandler.Create)
	api.GET("/delegations/:delegationId", delegationHandler.GetById)
	api.PUT("/delegations/:delegationId/approve", delegationHandler.ApproveById)
	api.PUT("/delegations/:delegationId/cancel", delegationHandler.CancelById)
	api.PUT("/delegations/:delegationId/unauthorize", delegationHandler.UnauthorizeById)
	api.DELETE("/delegations/:delegationId", delegationHandler.DeleteById)

	// 待受開始
	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}
