package model

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/Mire0726/Genkiyoho/backend/server/db"
	"time"
	"log"
)

type UserCondition struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ConditionID string `json:"condition_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	DamagePoint int `json:"damage_point"`
}

// conditionの登録
func InsertUserCondition(record *UserCondition) error {
	_, err := db.Conn.Exec(
		"INSERT INTO user_conditions (id, user_id, condition_id, start_date, end_date, damage_point) VALUES (?, ?, ?, ?, ?, ?)",
		record.ID,
		record.UserID,
		record.ConditionID,
		record.StartDate,
		record.EndDate,
		record.DamagePoint,
	)
	if err != nil {
		log.Printf("Error inserting user_condition into database: %v", err) // ログ追加
		return err
	}
	log.Println("User_condition successfully registered.") // 成功メッセージもログに記録
	return nil
}

//conditionの削除

//conditionの更新(ダメージなど)

//user_conditionの取得

//condition一覧取得