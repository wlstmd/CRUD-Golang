package network

import (
	"CRUD_SERVER/service"
	"CRUD_SERVER/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		router.registerPOST("/", userRouterInstance.create)
		router.registerGET("/", userRouterInstance.get)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)
	})

	return userRouterInstance
}

// CRUD

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("CREATE")

	u.router.okResponse(c, &types.CreateUserResponse{
		UserResponse: &types.UserResponse{
			APIResponse: types.NewAPIResponse("success", 1),
			User:        nil,
		},
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("GET")

	u.router.okResponse(c, &types.GetUserResponse{
		UserResponse: &types.UserResponse{
			APIResponse: types.NewAPIResponse("success", 1),
			User:        nil,
		},
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("UPDATE")

	u.router.okResponse(c, &types.UpdateUserResponse{
		UserResponse: &types.UserResponse{
			APIResponse: types.NewAPIResponse("success", 1),
			User:        nil,
		},
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("DELETE")

	u.router.okResponse(c, &types.DeleteUserResponse{
		UserResponse: &types.UserResponse{
			APIResponse: types.NewAPIResponse("success", 1),
			User:        nil,
		},
	})
}
