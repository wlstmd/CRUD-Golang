package network

import (
	"CRUD_SERVER/service"
	"CRUD_SERVER/types"
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
	var req types.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewAPIResponse("fail", -1, err.Error()),
		})
	} else if err = u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewAPIResponse("fail", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewAPIResponse("success", 1, nil),
		})
	}
}

func (u *userRouter) get(c *gin.Context) {

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewAPIResponse("success", 1, nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewAPIResponse("fail", -1, err.Error()),
		})
	} else if err = u.userService.Update(req.Name, req.UpdatedAge); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewAPIResponse("fail", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewAPIResponse("success", 1, nil),
		})
	}
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewAPIResponse("fail", -1, err.Error()),
		})
	} else if err = u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewAPIResponse("fail", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewAPIResponse("success", 1, nil),
		})
	}
}
