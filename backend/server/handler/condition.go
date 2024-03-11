package handler

import (
	"net/http"
    // "time"

	"github.com/Mire0726/Genkiyoho/backend/server/context/auth"
	"github.com/Mire0726/Genkiyoho/backend/server/model"
	"github.com/labstack/echo/v4"
    "log"
)

type conditionCreateRequest struct {
    Condition_id int `json:"condition_id"`
    ConditionName string `json:"condition_name"`
    Start_date string `json:"start_date"`
    End_date string `json:"end_date"`
    Damage_points int `json:"damage_points"`
}

// conditionの登録
func HandleConditionCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &conditionCreateRequest{}
		if err := c.Bind(req); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body format")
        }

        // Contextから認証済みのユーザIDを取得
        ctx := c.Request().Context()
        userID := auth.GetUserIDFromContext(ctx)
        if userID == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
        }
        log.Println(userID)
        log.Println("41")

		// 対象のユーザデータを取得（存在チェック）
        userData, err := model.SelectUserByPrimaryKey(userID)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user data")
        }
        if userData == nil {
            return echo.NewHTTPError(http.StatusNotFound, "User not found")
        }
        log.Println("52")

		if err := model.InsertUserCondition(&model.UserCondition{
            UserID: userID,
            ConditionID: req.Condition_id,
            StartDate: req.Start_date,
            EndDate: req.End_date,
            DamagePoint: req.Damage_points,
            
        }); err != nil {
            return err
        
        }
		return c.NoContent(http.StatusOK)
	}
}
// 特定のユーザーのuser_conditionの取得	
func HandleuserConditionGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        ctx := c.Request().Context()
        userID := auth.GetUserIDFromContext(ctx)
        if userID == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
        }
        log.Println("userIDisoK")
        // 対象のユーザデータを取得（存在チェック）
        userData, err := model.SelectUserByPrimaryKey(userID)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user data")
        }
        if userData == nil {
            return echo.NewHTTPError(http.StatusNotFound, "User not found")
        }
        log.Println("userDataisoK")
        userConditions, err := model.GetUserCondition(userID)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch user conditions")
        }
        return c.JSON(http.StatusOK, userConditions)
    }
}

//conditionの削除

//conditionの更新(ダメージなど)


