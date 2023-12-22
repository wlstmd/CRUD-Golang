package network

import (
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
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
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
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("GET")
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("UPDATE")
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("DELETE")
}
