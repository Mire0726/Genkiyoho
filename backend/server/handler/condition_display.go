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
		// if condition.ConditionID == 1003 {
		// 	c:= condition.Region
		// 	if weather.CheckPollen(c){
		// 	todayConditions_e = append(todayConditions_e, condition)
		// 	}
		// }
		if condition.ConditionID == 1004 {
			c:= condition.Region
			if weather.CheckPressure(c){
			todayConditions_e = append(todayConditions_e, condition)
			}
		}
		
	}

	return c.JSON(http.StatusOK, todayConditions_e)
}

func HandleUserTodayConditionGet(c echo.Context) error {
	userID := auth.GetUserIDFromContext(c.Request().Context())
	if userID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
	}
	conditions, err := model.GetUserConditions(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user conditions: "+err.Error())
	}

	// 各条件のconditionnameとdamagepointを格納するためのスライス
	var conditionDetails []map[string]interface{}

	for _, condition := range conditions {
		// 現在のサイクルに一致するかチェック
		startDate, err := time.Parse(time.RFC3339, condition.StartDate.Format(time.RFC3339))
		if err == nil && isInCurrentCycle(startDate, condition.Duration, condition.CycleLength) {
			conditionDetail := map[string]interface{}{
				"condition_name": condition.Name,
				"damage_point":   condition.DamagePoint,
			}
			conditionDetails = append(conditionDetails, conditionDetail)
		}

		// 環境条件に一致するかチェック
		if condition.ConditionID == 1004 && weather.CheckPressure(condition.Region) {
			conditionDetail := map[string]interface{}{
				"condition_name": condition.Name,
				"damage_point":   condition.DamagePoint,
			}
			conditionDetails = append(conditionDetails, conditionDetail)
		}

		if condition.ConditionID == 1006 && weather.CheckWeather(condition.Region) {
			conditionDetail := map[string]interface{}{
				"condition_name": condition.Name,
				"damage_point":   condition.DamagePoint,
			}
			conditionDetails = append(conditionDetails, conditionDetail)
	}
}
	return c.JSON(http.StatusOK, conditionDetails)
}

func HandleUserTodayPointGet(c echo.Context) error {
	userID := auth.GetUserIDFromContext(c.Request().Context())
	if userID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
	}
	conditions, err := model.GetUserConditions(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user conditions: "+err.Error())
	}

	// ユーザの元気ポイントを計算
	var genkipoint int
	for _, condition := range conditions {
		// 現在のサイクルに一致するかチェック
		startDate, err := time.Parse(time.RFC3339, condition.StartDate.Format(time.RFC3339))
		if err == nil && isInCurrentCycle(startDate, condition.Duration, condition.CycleLength) {
			genkipoint += condition.DamagePoint
		}

		// 環境条件に一致するかチェック
		if condition.ConditionID == 1004 && weather.CheckPressure(condition.Region) {
			genkipoint += condition.DamagePoint
		}
		// ここに他の環境条件のチェックを追加することができます
	}

	var genkiHP int
	genkiHP = 100 - genkipoint
	return c.JSON(http.StatusOK, genkiHP)
}