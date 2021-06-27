package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/opaku-dummy-backend/config"
	"github.com/ydhnwb/opaku-dummy-backend/handler"
	"github.com/ydhnwb/opaku-dummy-backend/middleware"
	"github.com/ydhnwb/opaku-dummy-backend/repo"
	"github.com/ydhnwb/opaku-dummy-backend/service"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                   = config.SetupDatabaseConnection()
	userRepo           repo.UserRepository        = repo.NewUserRepo(db)
	productRepo        repo.ProductRepository     = repo.NewProductRepo(db)
	transactionRepo    repo.TransactionRepository = repo.NewTransactionRepo(db)
	cartRepo           repo.CartRepository        = repo.NewCartRepo(db)
	transactionService service.TransactionService = service.NewTransactionService(transactionRepo, productRepo)
	authService        service.AuthService        = service.NewAuthService(userRepo)
	cartService        service.CartService        = service.NewCartService(cartRepo)
	jwtService         service.JWTService         = service.NewJWTService()
	userService        service.UserService        = service.NewUserService(userRepo)
	productService     service.ProductService     = service.NewProductService(productRepo)
	transactionHandler handler.TransactionHandler = handler.NewTransactionHandler(transactionService, jwtService)
	authHandler        handler.AuthHandler        = handler.NewAuthHandler(authService, jwtService, userService)
	userHandler        handler.UserHandler        = handler.NewUserHandler(userService, jwtService)
	productHandler     handler.ProductHandler     = handler.NewProductHandler(productService, jwtService)
	cartHandler        handler.CartHandler        = handler.NewCartHandler(cartService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	productRoutes := server.Group("api/product")
	{
		productRoutes.GET("/", productHandler.All)
		productRoutes.GET("/:id", productHandler.FindOneProductByID)
		productRoutes.GET("/search", productHandler.FindProductsByName)
	}

	cartRoutes := server.Group("api/cart", middleware.AuthorizeJWT(jwtService))
	{
		cartRoutes.POST("/", cartHandler.AddToCart)
		cartRoutes.GET("/", cartHandler.FindAllCart)
		cartRoutes.DELETE("/:id", cartHandler.DeleteCart)
	}

	transactionRoutes := server.Group("api/transaction", middleware.AuthorizeJWT(jwtService))
	{
		transactionRoutes.GET("/", transactionHandler.FindAllMyTransaction)
		transactionRoutes.POST("/", transactionHandler.CreateTransaction)
	}

	server.Run()
}
