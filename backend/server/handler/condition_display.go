package handler

import (
	"net/http"
	"time"

	"github.com/Mire0726/Genkiyoho/backend/server/context/auth"
	"github.com/Mire0726/Genkiyoho/backend/server/model"
	"github.com/Mire0726/Genkiyoho/backend/server/weather"
	"github.com/labstack/echo/v4"
)
func isInCurrentCycle(startDate time.Time, duration, cycleLength int) bool {
    today := time.Now()

    // EndDateが定義されていない場合、サイクルに基づいて終了日を計算
    expectedEndDate := startDate.AddDate(0, 0, duration+cycleLength-1)

    // 今日の日付が開始日と計算された終了日の間にあるかを判定
    return !today.Before(startDate) && !today.After(expectedEndDate)
}



func HandleUserTodayCycleConditionGet(c echo.Context) error {
	userID := auth.GetUserIDFromContext(c.Request().Context())
	if userID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
	}
	conditions, err := model.GetUserConditions(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user conditions: "+err.Error())
	}

	var todayConditions_c []model.UserCondition
	for _, condition := range conditions {
		startDate, err := time.Parse(time.RFC3339, condition.StartDate.Format(time.RFC3339)) // Convert condition.StartDate to string
		if err != nil {
			continue // 日付のパースに失敗した場合は、このコンディションをスキップします。
		}
		if isInCurrentCycle(startDate, condition.Duration, condition.CycleLength) {
			todayConditions_c = append(todayConditions_c, condition)
		}
	}

	return c.JSON(http.StatusOK, todayConditions_c)
}

func HandleUserTodayEnvironmentConditionGet(c echo.Context) error {
	userID := auth.GetUserIDFromContext(c.Request().Context())
	if userID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
	}
	conditions, err := model.GetUserConditions(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user conditions: "+err.Error())
	}

	var todayConditions_e []model.UserCondition
	for _, condition := range conditions { 
		if condition.ConditionID == 1003 {
			c:= condition.Region
			if weather.CheckPollen(c){
			todayConditions_e = append(todayConditions_e, condition)
			}
		}
		if condition.ConditionID == 1004 {
			c:= condition.Region
			if weather.CheckPressure(c){
			todayConditions_e = append(todayConditions_e, condition)
			}
		}
		
	}

	return c.JSON(http.StatusOK, todayConditions_e)
}
