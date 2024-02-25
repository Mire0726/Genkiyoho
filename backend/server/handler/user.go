package handler

import (
	"net/http"
	"github.com/google/uuid"
	"time"
	"net/mail"

	"github.com/Mire0726/Genkiyoho/backend/server/model"
	"github.com/labstack/echo/v4"
	"log"
	"github.com/Mire0726/Genkiyoho/backend/server/context/auth"
	"errors"
	"strconv"
)

type userCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// validateEmail はメールアドレスが有効な形式かどうかを検証します。
func validateEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

func HandleUserCreate() echo.HandlerFunc {
    return func(c echo.Context) error {
		req := &userCreateRequest{}
		if err := c.Bind(req); err != nil {
			return err
		}
		log.Println("handler,line34")

    // メールアドレスの形式を検証
    if !validateEmail(req.Email) {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid email format")
    }

    // UUIDで認証トークンを生成
    authToken, err := uuid.NewRandom()
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate authentication token")
    }

    now := time.Now() // 現在の時刻
	// データベースにユーザデータを登録する
	if err := model.InsertUser(&model.User{
		AuthToken: authToken.String(),
        Email:     req.Email,
        Password:  req.Password,
        Name:      req.Name,
        CreatedAt: now,
        UpdatedAt: now,
	}); err != nil {
		return err
	}
	// 生成した認証トークンを返却
	return c.JSON(http.StatusOK, &userCreateResponse{Token: authToken.String()})
}
}

type userCreateResponse struct {
	Token string `json:"token"`
}

// GetUser 全ユーザ情報を取得するエンドポイントのハンドラ
func HandleGetUser() echo.HandlerFunc{
	return func(c echo.Context) error {

		users, err :=model.GetAllUsers()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// echo.ContextのJSONメソッドを使用してユーザリストをJSON形式で返す
		return c.JSON(http.StatusOK, users)
	}
}

// HandleUserGet ユーザ情報を取得するエンドポイントのハンドラ
func HandleUserGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        // Contextから認証済みのユーザIDを取得
        ctx := c.Request().Context()
        userIDStr := auth.GetUserIDFromContext(ctx)
        if userIDStr == "" {
            return errors.New("userID is empty")
        }

        userID, err := strconv.Atoi(userIDStr)
        if err != nil {
            return errors.New("invalid userID format")
        }

        // ユーザデータの取得処理を実装
        user, err := model.SelectUserByPrimaryKey(userID)
        if err != nil {
            return err
        }
        if user == nil {
            return errors.New("user not found")
        }

        // レスポンスに必要な情報を詰めて返却
        return c.JSON(http.StatusOK, user)
    }
}

type userUpdateRequest struct {
	Name string `json:"name"`
}
// HandleUserUpdate ユーザ情報を更新するエンドポイントのハンドラ
func HandleUserUpdate() echo.HandlerFunc {
    return func(c echo.Context) error {
        log.Println("99")

        // リクエストBodyから更新後情報を取得
        req := &userUpdateRequest{}
        if err := c.Bind(req); err != nil {
            return err
        }
        
        // Contextから認証済みのユーザIDを取得
        ctx := c.Request().Context()
        userIDStr := auth.GetUserIDFromContext(ctx)
        if userIDStr == "" {
            return errors.New("userID is empty")
        }

        userID, err := strconv.Atoi(userIDStr)
        if err != nil {
            return errors.New("invalid userID format")
        }

		userData,err:=model.SelectUserByPrimaryKey(userID)
		if err !=nil {
			return errors.New("userData is empty")
		}
        userData.Name = req.Name
        // ユーザデータの更新処理を実装
        err = model.UpdateUserByPrimaryKey(userData)
        if err != nil {
            return err
        }
        
        return c.NoContent(http.StatusOK)
    }
}



