package middleware

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/Mire0726/Genkiyoho/backend/server/context/auth"
	"github.com/Mire0726/Genkiyoho/backend/server/model"
	// "github.com/Mire0726/Genkiyoho/backend/server/handler"
	// _ "github.com/go-sql-driver/mysql" // MySQLドライバーをインポート
)

// AuthenticateMiddleware ユーザ認証を行ってContextへユーザID情報を保存する
func AuthenticateMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		fmt.Println("middleware,line40")
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			fmt.Println("middleware,line43")
			// リクエストヘッダからx-token(認証トークン)を取得
			token := c.Request().Header.Get("x-token")
			if token == "" {
				return errors.New("x-token is empty")
			}
			fmt.Println("middleware,line49")
			// データベースから認証トークンに紐づくユーザの情報を取得
			user, err := model.SelectUserByAuthToken(token)
			if err != nil {
				return err
			}
			if user == nil {
				return fmt.Errorf("user not found. token=%s", token)
			}
			fmt.Println("middleware,line57")
			// ユーザIDをContextへ保存して以降の処理に利用する
			c.SetRequest(c.Request().WithContext(auth.SetUserID(ctx, user.ID)))
			fmt.Println("middleware,line61")
			// 次の処理
			return next(c)
		}
	}
}
