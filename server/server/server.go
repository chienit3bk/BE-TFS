package server

import (
	"fmt"
	"net/http"
	userAuthen "project/packages/authentication/user"
	"project/packages/handlers/bill"
	"project/packages/handlers/order"
	"project/packages/handlers/product"
	userHandlers "project/packages/handlers/user"
	"project/packages/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func RunServer() {
	///////
	fmt.Println("Starting server. Please open http://localhost:8080")
	defer func() {
		fmt.Println("Server is stopped")
	}()
	///////

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.ContentTypeCheckingMiddleware)
	router.Methods(http.MethodGet).Path("/search_db").HandlerFunc(product.SearchByNameDB)
	router.Methods(http.MethodGet).Path("/search_es").HandlerFunc(product.SearchByNameES)
	router.Methods(http.MethodGet).Path("/products/{id:[0-9]+}").HandlerFunc(product.GetByID)
	router.Methods(http.MethodGet).Path("/products").HandlerFunc(product.GetAll)
	router.Methods(http.MethodGet).Path("/products/{category}").HandlerFunc(product.Filter)
	router.Methods(http.MethodGet).Path("/products/filter/advanced").HandlerFunc(product.FilterAdvanced)
	router.Methods(http.MethodPost).Path("/login").HandlerFunc(userAuthen.Login)
	router.Methods(http.MethodPost).Path("/signup").HandlerFunc(userAuthen.Register)
	router.Methods(http.MethodPost).Path("/forgotpassword").HandlerFunc(userHandlers.ForgotPasswordStep1)
	router.Methods(http.MethodPost).Path("/forgotpassword/authen").HandlerFunc(userHandlers.ForgotPasswordStep2)
	router.Methods(http.MethodPost).Path("/forgotpassword/newpass").HandlerFunc(userHandlers.ForgotPasswordStep3)

	//Router for user
	userSubRouter := router.NewRoute().Subrouter()
	userSubRouter.Use(middleware.UserAuthorize)
	userSubRouter.Use(middleware.SpecificUserAuthorize)

	//Router for admin
	adminSubRouter := router.NewRoute().Subrouter()
	adminSubRouter.Use(middleware.AdminAuthorize)

	//user
	userGet := userSubRouter.Methods(http.MethodGet).Subrouter()
	userGet.Path("/orders/{userId:[0-9]+}").HandlerFunc(order.GetAll)
	userGet.Path("/orders/{userId:[0-9]+}/{orderId:[0-9]+}").HandlerFunc(order.GetByID)
	userGet.Path("/account/{userId:[0-9]+}").HandlerFunc(userHandlers.GetInfoByID)

	userPost := userSubRouter.Methods(http.MethodPost).Subrouter()
	userPost.Path("/orders/{userId:[0-9]+}").HandlerFunc(order.Create)
	userPost.Path("/orders/bill/{userId:[0-9]+}").HandlerFunc(bill.Create)
	userPost.Path("/orders/updatebill/{userId:[0-9]+}").HandlerFunc(bill.Update)
	userPost.Path("/orders/cancelbill/{userId:[0-9]+}").HandlerFunc(bill.Cancel)
	userPost.Path("/account/password/{userId:[0-9]+}").HandlerFunc(userHandlers.ChangePassword)

	userPut := userSubRouter.Methods(http.MethodPut).Subrouter()
	userPut.Path("/users/{userId:[0-9]+}").HandlerFunc(userHandlers.Update)

	userDelete := userSubRouter.Methods(http.MethodDelete).Subrouter()
	userDelete.Path("/orders/{userId:[0-9]+}/{orderId:[0-9]+}").HandlerFunc(order.Delete)

	//admin
	adminGet := adminSubRouter.Methods(http.MethodGet).Subrouter()
	adminGet.Path("/users").HandlerFunc(userHandlers.GetAllUser)
	adminGet.Path("/statistic").HandlerFunc(order.StatisticsHourly)

	adminPost := adminSubRouter.Methods(http.MethodPost).Subrouter()
	adminPost.Path("/products").HandlerFunc(product.Create)
	adminPost.Path("/products/importfile").HandlerFunc(product.CreateMulti)

	adminPut := adminSubRouter.Methods(http.MethodPut).Subrouter()
	adminPut.Path("/products").HandlerFunc(product.Update)

	adminDelete := adminSubRouter.Methods(http.MethodDelete).Subrouter()
	adminDelete.Path("/products/{productId:[0-9]+}").HandlerFunc(product.Delete)

	_ = cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
	}).Handler(router)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})
	handler := c.Handler(router)
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}

	// err := http.ListenAndServe(":8080", router)
	// if err != nil {
	// 	panic(err)
	// }
}
