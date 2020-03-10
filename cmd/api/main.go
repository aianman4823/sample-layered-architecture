package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/julienschmidt/httprouter"
	"local.packages/interfaces/handler"
	"local.packages/infrastructure/persistence"
	"local.packages/usecase"
)

func main() {
	// 依存関係を注入 (DIとはまたちょいと違う)
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/v2/get", userHandler.HandleUserGet)
	router.POST("/api/v2/signup", userHandler.HandleUserSignup)

	// サーバー起動
	fmt.Println("Server Running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}