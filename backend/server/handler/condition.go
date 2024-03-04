package handler

import (
	"net/http"
	"github.com/google/uuid"
	"time"
)

// conditionの登録
func HandleConditionCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &conditionCreateRequest{}
		if err := c.Bind(req); err !=
		
//conditionの削除

//conditionの更新(ダメージなど)

//user_conditionの取得

//condition一覧取得