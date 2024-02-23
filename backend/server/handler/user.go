package handler

import (
	"net/http"
	"github.com/google/uuid"
	"time"
	"net/mail"

	"github.com/Mire0726/Genkiyoho/backend/server/db"
	"github.com/Mire0726/Genkiyoho/backend/server/model"
	"github.com/labstack/echo/v4"
	"database/sql"
	"log"
)

type userCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AppHandler struct {
	DB *sql.DB
}

// validateEmail はメールアドレスが有効な形式かどうかを検証します。
func validateEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

func (h *AppHandler) HandleUserCreate(c echo.Context) error {
    req := &userCreateRequest{}
    if err := c.Bind(req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
    }

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

    // データベースにユーザデータを登録
    user := &model.User{
        AuthToken: authToken.String(),
        Email:     req.Email,
        Password:  req.Password,
        Name:      req.Name,
        CreatedAt: now,
        UpdatedAt: now,
    }
    if err := model.InsertUser(h.DB, user); err != nil {
        log.Printf("Error registering user: %v", err) // ログ追加
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to register user")
    }

    // 生成した認証トークンを返却
    return c.JSON(http.StatusOK, &userCreateResponse{Token: authToken.String()})
}

type userCreateResponse struct {
	Token string `json:"token"`
}

// GetUser 全ユーザ情報を取得するエンドポイントのハンドラ
func (h *AppHandler) HandleGetUser(c echo.Context) error{
	dbConn, err := db.ConnectToDB()
	if err != nil {
		// echo.ContextのErrorメソッドを使用してエラーを返す
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer dbConn.Close()

	rows, err := dbConn.Query("SELECT id,auth_token, email, password, name, created_at, updated_at FROM users")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.AuthToken,&u.Email, &u.Password, &u.Name, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		users = append(users, u)
	}

	// echo.ContextのJSONメソッドを使用してユーザリストをJSON形式で返す
	return c.JSON(http.StatusOK, users)
}
