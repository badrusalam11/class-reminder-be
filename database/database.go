package database

import (
	"class-reminder-be/database/helper"
	library "class-reminder-be/library"
	"class-reminder-be/model"
	"database/sql"
	"fmt"
)

func GeneralSelect(query string, args ...interface{}) (map[string]interface{}, error) {
	rows, err := helper.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))

	for i := range values {
		valuePointers[i] = &values[i]
	}

	if rows.Next() { // Move to the first row
		err = rows.Scan(valuePointers...)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("record not found")
			}
			return nil, err
		}

		data := make(map[string]interface{})
		for i, column := range columns {
			val := values[i]
			data[column] = val
		}

		return data, nil
	}
	return nil, fmt.Errorf("no rows found")
}

func GeneralSelectRows(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := helper.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))

		for i := range values {
			valuePointers[i] = &values[i]
		}

		err := rows.Scan(valuePointers...)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("record not found")
			}
			return nil, err
		}

		data := make(map[string]interface{})
		for i, column := range columns {
			val := values[i]
			data[column] = val
		}

		result = append(result, data)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no rows found")
	}

	return result, nil
}

func GeneralQuery(query string, args ...interface{}) error {
	rows, err := helper.Db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func GetEventFromDB(idEvent int) (map[string]interface{}, error) {
	query := "SELECT e.*, is_specific_user FROM tbl_event e JOIN tbl_event_type et ON e.id_event_type = et.id WHERE e.id = ?"
	data, err := GeneralSelect(query, idEvent)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetUserEventFromDB(idEvent int) ([]map[string]interface{}, error) {
	// query := "SELECT ue.username, un.notif_id, un.is_allowed, no_hp FROM tbl_user_event ue JOIN tbl_user_notif un ON ue.username = un.username WHERE is_allowed = 1 and ue.id_event = ?"
	query := `
	SELECT name, nim, title AS class, SCHEDULE AS time, no_hp
	FROM tbl_user_event ue 
	JOIN tbl_user_notif un ON ue.username = un.username
	JOIN tbl_user_student us ON ue.username = us.username
	JOIN tbl_event e ON ue.id_event=e.id
	WHERE is_allowed = 1 and ue.id_event = ?
	`
	data, err := GeneralSelectRows(query, idEvent)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}

	// eventUser := &EventUser{
	// 	Username:   data["username"].(string),
	// 	Notif_id:   string(data["notif_id"].([]uint8)), // Convert []uint8 to string
	// 	Is_allowed: data["is_allowed"].(int),
	// }

	return data, nil
}

func GetContentFromDB(trxType string) (map[string]interface{}, error) {
	query := "SELECT * FROM tbl_content_notif WHERE trx_type=?"
	data, err := GeneralSelect(query, trxType)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func GetTrxTypeFromDB(idEvent int) (map[string]interface{}, error) {
	query := "SELECT trx_type FROM tbl_event e JOIN tbl_event_type et ON e.id_event_type=et.id WHERE e.id = ?"
	data, err := GeneralSelect(query, idEvent)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetUserNotifFromDB() ([]map[string]interface{}, error) {
	query := "SELECT username, notif_id, no_hp, is_allowed FROM tbl_user_notif where is_allowed = 1"
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func InsertEventToDB(data model.EventCreateRequest, jobEvery string) (int64, error) {
	query := "INSERT INTO tbl_event (title, description, schedule, job_every, id_event_type) VALUES (?, ?, ?, ?,?)"
	result, err := helper.Db.Exec(query, data.Title, data.Description, data.Schedule, jobEvery, data.IdEventType)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func GetUserFromDB(username string) (map[string]interface{}, error) {
	query := "SELECT * FROM tbl_user WHERE username = ?"
	result, err := GeneralSelect(query, username)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateJwtToDB(username string, jwt string) error {
	query := "UPDATE tbl_user SET token_key = ?, last_login=? WHERE username = ?"
	_, err := helper.Db.Exec(query, jwt, library.CurrTimestamp(), username)
	if err != nil {
		return err
	}
	return nil
}

func UpdateNotifId(username string, notifId string) (string, error) {
	query := "UPDATE tbl_user_notif SET notif_id = ?, last_update=?, is_allowed=1 WHERE username = ?"
	result, err := helper.Db.Exec(query, notifId, library.CurrTimestamp(), username)
	if err != nil {
		return "01", err
	}
	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "01", err
	}
	if rowsAffected == 0 {
		return "02", err
	}
	return "", nil
}

func GetNumberForBlast() ([]map[string]interface{}, error) {
	query := "SELECT no_hp from tbl_user_notif where is_allowed = 1"
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func InsertBlastHistory(message string, user_success int) error {
	query := "INSERT INTO tbl_blast_history (message, user_success, created_at) VALUES (?, ?, ?)"
	_, err := helper.Db.Exec(query, message, user_success, library.CurrTimestamp())
	if err != nil {
		return err
	}

	return nil
}

func GetBlastHistory() ([]map[string]interface{}, error) {
	query := "SELECT * from tbl_blast_history ORDER BY id DESC"
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}
