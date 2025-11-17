package rest

import (
	"fmt"
	"kholabazar/config"
	"kholabazar/rest/handlers/product"
	"kholabazar/rest/handlers/review"
	"kholabazar/rest/handlers/user"
	"kholabazar/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	conf           *config.Config
	userHandler    *user.Handler
	productHandler *product.Handler
	reviewHandler  *review.Handler
}

func NewServer(
	conf *config.Config,
	userHandler *user.Handler,
	productHandler *product.Handler,
	reviewHandler *review.Handler,
) *Server {
	return &Server{
		conf:           conf,
		userHandler:    userHandler,
		productHandler: productHandler,
		reviewHandler:  reviewHandler,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.PreFlight, middleware.Cors, middleware.Logger)
	wrappedMux := manager.WrapMux(mux)
	server.userHandler.RegisterRoutes(mux, manager)
	server.productHandler.ProductRoutes(mux, manager)
	server.reviewHandler.ReviewRoutes(mux, manager)
	addr := "127.0.0.1:" + strconv.Itoa(server.conf.HttpPort)
	fmt.Println("Server is running on port :", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server :", err)
		os.Exit(1)
	}
}
