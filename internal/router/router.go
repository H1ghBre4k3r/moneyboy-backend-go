package router

import "github.com/gofiber/fiber/v2"

type RouterParams struct {
	UserService UserService
	AuthService AuthService
}

type Router struct {
	controller []interface{}
}

func New(base fiber.Router, params *RouterParams) *Router {

	userController := userController(base.Group("/user"), params.UserService)
	authController := authController(base.Group("/auth"), params.AuthService)

	router := &Router{}
	router.controller = append(router.controller, userController, authController)

	return router
}
