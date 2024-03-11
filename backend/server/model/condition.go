package model

import (
	"log"

	"github.com/Mire0726/Genkiyoho/backend/server/db"
)

type UserCondition struct {
	UserID    string    `json:"user_id"`
	ConditionID int `json:"condition_id"`
	ConditionName string `json:"condition_name"` // この行を追加して、条件名を格納するためのフィールドを追加します
	StartDate string `json:"start_date"`
	EndDate   string`json:"end_date"`
	DamagePoint int `json:"damage_point"`
}

// conditionの登録
func InsertUserCondition(record *UserCondition) error {
	_, err := db.Conn.Exec(
		"INSERT INTO user_condition (user_id, condition_id, start_date, end_date, damage_points) VALUES (?, ?, ?, ?, ?)",
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

// func GetUserCondition(userID string) ([]UserCondition, error) {
//     rows, err := db.Conn.Query("SELECT user_id, condition_id, start_date, end_date, damage_points FROM user_condition WHERE user_id=?", userID)
//     if err != nil {
//         log.Printf("Error querying user_condition from database: %v", err)
//         return nil, err
//     }
//     defer rows.Close()

//     var userConditions []UserCondition

//     for rows.Next() {
//         var uc UserCondition
//         if err := rows.Scan(&uc.UserID, &uc.ConditionID, &uc.StartDate, &uc.EndDate, &uc.DamagePoint); err != nil {
//             log.Printf("Error scanning user_condition from database: %v", err)
//             return nil, err
//         }

//         userConditions = append(userConditions, uc)
//     }
//     if err = rows.Err(); err != nil {
//         log.Printf("Error during rows iteration: %v", err)
//         return nil, err
//     }

//     return userConditions, nil
// }
func GetUserCondition(userID string) ([]UserCondition, error) {
    query := `
    SELECT 
        uc.user_id, 
        uc.condition_id, 
        c.name as condition_name, 
        uc.start_date, 
        uc.end_date, 
        uc.damage_points 
    FROM 
        user_condition uc 
    JOIN 
        conditions c ON uc.condition_id = c.id 
    WHERE 
        uc.user_id=?
    `

    rows, err := db.Conn.Query(query, userID)
    if err != nil {
        log.Printf("Error querying user_condition and condition from database: %v", err)
        return nil, err
    }
    defer rows.Close()
    log.Println("user_condition and condition successfully fetched.")
    var userConditions []UserCondition
    for rows.Next() {
        var uc UserCondition
        var conditionName string
        if err := rows.Scan(&uc.UserID, &uc.ConditionID, &conditionName, &uc.StartDate, &uc.EndDate, &uc.DamagePoint); err != nil {
            log.Printf("Error scanning user_condition and condition from database: %v", err)
            return nil, err
        }
        uc.ConditionName = conditionName // この行を追加して、取得した条件名をUserCondition構造体に格納します
        userConditions = append(userConditions, uc)
    }
    if err = rows.Err(); err != nil {
        log.Printf("Error during rows iteration: %v", err)
        return nil, err
    }
    log.Println("user_condition and condition successfully fetched!!.")
    return userConditions, nil
}


//conditionの削除

//conditionの更新(ダメージなど)

//condition一覧取得
