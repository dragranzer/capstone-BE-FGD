package presentation

import (
	"fmt"
	"net/http"

	"github.com/dragranzer/capstone-BE-FGD/features/threads"
	"github.com/dragranzer/capstone-BE-FGD/middleware"
	"github.com/labstack/echo/v4"
)

type ThreadsHandler struct {
	threadBussiness threads.Bussiness
}

func NewThreadHandler(ub threads.Bussiness) *ThreadsHandler {
	return &ThreadsHandler{
		threadBussiness: ub,
	}
}

func (uh *ThreadsHandler) GetThreadHome(c echo.Context) error {
	// userID := request.UserID{}
	// c.Bind(&userID)
	temp := middleware.ExtractClaim(c)
	ownerID := temp["user_id"].(float64)
	fmt.Println(ownerID)
	data := threads.Core{
		OwnerID: int(ownerID),
	}
	threads, err := uh.threadBussiness.GetThreadHome(data)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    threads,
		"message": "data success di masukkan",
	})
}
