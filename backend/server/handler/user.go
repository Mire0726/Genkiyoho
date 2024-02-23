package handler

import (
	"net/http"
	"github.com/google/uuid"
	"time"

	"github.com/Mire0726/Genkiyoho/backend/server/db"
	"github.com/Mire0726/Genkiyoho/backend/server/model"
	"github.com/labstack/echo/v4"
	"database/sql"
)

type userCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AppHandler struct {
	DB *sql.DB
}

func (h *AppHandler) HandleUserCreate(c echo.Context) error {
	req := &userCreateRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}
	// UUIDで認証トークンを生成する
	authToken, err := uuid.NewRandom()
	if err != nil {
		// UUIDの生成に失敗した場合は、500 Internal Server Error エラーを返す
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate authentication token")
	}

	now := time.Now() // 現在の時刻

	// データベースにユーザデータを登録する
	user := &model.User{
		AuthToken: authToken.String(),
		Email:     req.Email,
		Password:  req.Password, // パスワードはハッシュ化することを推奨
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := model.InsertUser(h.DB, user); err != nil {
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
