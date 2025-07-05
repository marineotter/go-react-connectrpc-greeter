package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"connectrpc.com/connect"
	"github.com/rs/cors"

	greeterv1 "backend/gen/greeter/v1"
	"backend/gen/greeter/v1/greeterv1connect"
)

type GreeterServer struct{}

func (s *GreeterServer) Greet(
	ctx context.Context,
	req *connect.Request[greeterv1.GreetRequest],
) (*connect.Response[greeterv1.GreetResponse], error) {
	name := req.Msg.Name
	if name == "" {
		name = "World"
	}

	message := fmt.Sprintf("Hello, %s!", name)

	return connect.NewResponse(&greeterv1.GreetResponse{
		Message: message,
	}), nil
}

func main() {
	greeterServer := &GreeterServer{}

	mux := http.NewServeMux()
	path, handler := greeterv1connect.NewGreeterServiceHandler(greeterServer)

	// APIパスをプレフィックス付きで設定
	mux.Handle("/api"+path, handler)
	// 開発環境互換用（直接アクセス）
	mux.Handle(path, handler)

	// 環境変数から設定を取得
	debug := os.Getenv("DEBUG") == "true"
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	if debug {
		// デバッグモードの場合は全てのオリジンを許可
		allowedOrigins = "*"
		fmt.Println("DEBUG mode enabled - allowing all origins")
	}

	// CORS設定
	corsOptions := cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}

	if debug {
		corsOptions.AllowedOrigins = []string{"*"}
		corsOptions.AllowCredentials = false // "*"を使用する場合はCredentialsをfalseにする必要がある
	} else {
		corsOptions.AllowedOrigins = strings.Split(allowedOrigins, ",")
	}

	c := cors.New(corsOptions)

	port := os.Getenv("PORT")
	if port == "" {
		port = "18080"
	}

	fmt.Printf("Server starting on :%s\n", port)
	fmt.Printf("Allowed origins: %s\n", allowedOrigins)
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(mux)))
}
