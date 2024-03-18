package handler

import (
	"fmt"
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
	fmt.Println(conditions)
	var todayConditions_e []model.UserCondition
	for _, condition := range conditions { 
		c:= condition.Region
		if condition.Name == "花粉" {
			if weather.CheckPollen(c)>0 {
				todayConditions_e = append(todayConditions_e, condition)
			}
		}
		if condition.Name == "気圧の不調" {
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
		startDate,err:= time.Parse(time.RFC3339, condition.StartDate.Format(time.RFC3339))
		if err == nil && isInCurrentCycle(startDate, condition.Duration, condition.CycleLength) {
			conditionDetail := map[string]interface{}{
				"condition_name": condition.Name,
				"damage_point":   condition.DamagePoint,
			}
			conditionDetails = append(conditionDetails, conditionDetail)
		} else{
		// 環境条件に一致するかチェック
		if condition.Name == "花粉"  && (weather.CheckPollen(condition.Region)>0) {
			if weather.CheckPollen(condition.Region)>0 && weather.CheckPollen(condition.Region)<30{
				conditionDetail := map[string]interface{}{
					"condition_name": condition.Name,
					"damage_point":   condition.DamagePoint,
				}
				conditionDetails = append(conditionDetails, conditionDetail)
			}
			if weather.CheckPollen(condition.Region)>=30 && weather.CheckPollen(condition.Region)<60{
				conditionDetail := map[string]interface{}{
					"condition_name": condition.Name,
					"damage_point":   condition.DamagePoint*2,
				}
				conditionDetails = append(conditionDetails, conditionDetail)
			}
			if weather.CheckPollen(condition.Region)>=60 && weather.CheckPollen(condition.Region)<100{
				conditionDetail := map[string]interface{}{
					"condition_name": condition.Name,
					"damage_point":   condition.DamagePoint*3,
				}
				conditionDetails = append(conditionDetails, conditionDetail)
			}
		}

		if condition.Name=="気圧の不調" && weather.CheckPressure(condition.Region) {
			conditionDetail := map[string]interface{}{
				"condition_name": condition.Name,
				"damage_point":   condition.DamagePoint,
			}
			conditionDetails = append(conditionDetails, conditionDetail)
		}

		if condition.Name=="雨による不調" && weather.CheckWeather(condition.Region) {
			conditionDetail := map[string]interface{}{
				"condition_name": condition.Name,
				"damage_point":   condition.DamagePoint,
			}
			conditionDetails = append(conditionDetails, conditionDetail)
		}
	}

	}
	return c.JSON(http.StatusOK, conditionDetails)
}


func HandleUserTodayPointGet(c echo.Context) error {
    userID := auth.GetUserIDFromContext(c.Request().Context())
    if userID == "" {
        return echo.NewHTTPError(http.StatusUnauthorized, "User ID is empty")
    }
    
    conditions, err := model.GetUserConditions(userID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user conditions: "+err.Error())
    }

    var totalDamagePoints int
    for _, condition := range conditions {
        startDate, err := time.Parse(time.RFC3339, condition.StartDate.Format(time.RFC3339))
        if err != nil {
            continue // If start date parsing fails, skip this condition.
        }

        // Check if the condition is within the current cycle
        if isInCurrentCycle(startDate, condition.Duration, condition.CycleLength) {
            totalDamagePoints += condition.DamagePoint
        }


		
        // Additional checks for environmental conditions
        switch condition.Name {
        case "花粉":
            pollenCount := weather.CheckPollen(condition.Region)
            switch {
            case pollenCount > 0 && pollenCount < 30:
                totalDamagePoints += condition.DamagePoint
            case pollenCount >= 30 && pollenCount < 60:
                totalDamagePoints += condition.DamagePoint * 2
            case pollenCount >= 60:
                totalDamagePoints += condition.DamagePoint * 3
            }

        case "気圧の不調":
            if weather.CheckPressure(condition.Region) {
                totalDamagePoints += condition.DamagePoint
            }

        case "雨による不調":
            if weather.CheckWeather(condition.Region) {
                totalDamagePoints += condition.DamagePoint
            }
        }
    }

    genkiHP := 100 - totalDamagePoints
    if genkiHP < 0 {
        genkiHP = 0 // Ensure that Genki HP does not go below 0
    }

    return c.JSON(http.StatusOK, genkiHP)
}

