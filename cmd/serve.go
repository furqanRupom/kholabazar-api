package cmd

import (
	"fmt"
	"kholabazar/config"
	"kholabazar/infra/db"
	"kholabazar/repo"
	"kholabazar/rest"
	 productHandler "kholabazar/rest/handlers/product"
	"kholabazar/rest/handlers/review"
	userHandler "kholabazar/rest/handlers/user"
	middleware "kholabazar/rest/middlewares"
	"kholabazar/user"
	"kholabazar/product"
	"os"
)

func Serve() {
	conf := config.GetConfig()
	dbCon, err := db.NewConnection(conf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if dbCon == nil {
		fmt.Println(err)
	}
	if dbCon != nil {
		fmt.Println("Database connected successfully!")
	}
	err = db.MigrateDB(dbCon,"./migrations")
	if err != nil {
		fmt.Println("Database migration failed!")
	}
	middleware := middleware.NewMiddlewares(conf)

	// repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	userSvc := user.NewService(userRepo)
	productSvc := product.NewService(productRepo)

	// handlers
	prdHandler := productHandler.NewHandler(middleware, productSvc)
	usrHandler := userHandler.NewHandler(conf,userSvc)
	reviewHandler := review.NewHandler()
	server := rest.NewServer(
		conf,
		usrHandler,
		prdHandler,
		reviewHandler,
	)
	server.Start()
}
