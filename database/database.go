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
	SELECT name, us.nim, title AS class, SCHEDULE AS time, no_hp
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

func EditEventToDB(data model.EventEditRequest, jobEvery string) error {
	query := "UPDATE tbl_event SET title=?, description=?, schedule=?, job_every=?, id_event_type=? WHERE id=?"
	_, err := helper.Db.Exec(query, data.Title, data.Description, data.Schedule, data.JobEvery, data.IdEventType, data.Id)
	if err != nil {
		return err
	}

	return nil
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

func GetCourse() ([]map[string]interface{}, error) {
	query := "SELECT * from tbl_event WHERE id_event_type=1"
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func GetCourseById(id int64) (map[string]interface{}, error) {
	query := "SELECT * from tbl_event WHERE id=?"
	data, err := GeneralSelect(query, id)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func RegisterStudent(name string, nim string, no_hp string, major string, class []int, tuition_fee int, va_account string, last_payment_date string) error {
	//insert to tbl_user_student
	query := "INSERT INTO tbl_user_student (username, name, nim, major) VALUES (?, ?, ?, ?)"
	_, err := helper.Db.Exec(query, nim, name, nim, major)
	if err != nil {
		return err
	}

	//insert to tbl_user_event
	for i := 0; i < len(class); i++ {
		query = "INSERT INTO tbl_user_event (username, nim, id_event) VALUES (?,?,?)"
		_, err = helper.Db.Exec(query, nim, nim, class[i])
		if err != nil {
			return err
		}

	}

	// insert to tbl_user_notif
	query = "INSERT INTO tbl_user_notif (username, nim, no_hp, last_update, is_allowed) VALUES (?, ?, ?, ?,?)"
	_, err = helper.Db.Exec(query, nim, nim, no_hp, library.CurrTimestamp(), 1)
	if err != nil {
		return err
	}

	// insert to tbl_user_payment
	query = "INSERT INTO tbl_user_payment (nim, bill, va_account, last_payment_date) VALUES (?, ?, ?, ?)"
	_, err = helper.Db.Exec(query, nim, tuition_fee, va_account, last_payment_date)
	if err != nil {
		return err
	}
	return nil
}

func EditStudent(name string, nim string, no_hp string, major string, class []int, tuition_fee int, va_account string, last_payment_date string) error {
	//update to tbl_user_student
	query := "UPDATE tbl_user_student SET username=?, name=?, nim=?, major=? WHERE nim=?"
	_, err := helper.Db.Exec(query, nim, name, nim, major, nim)
	if err != nil {
		return err
	}

	// delete tbl_user_event
	query = "DELETE FROM tbl_user_event WHERE nim=?"
	_, err = helper.Db.Exec(query, nim)
	if err != nil {
		return err
	}

	//insert to tbl_user_event
	for i := 0; i < len(class); i++ {
		query = "INSERT INTO tbl_user_event (username, nim, id_event) VALUES (?,?,?)"
		_, err = helper.Db.Exec(query, nim, nim, class[i])
		if err != nil {
			return err
		}

	}

	// insert to tbl_user_notif
	query = "UPDATE tbl_user_notif SET username=?, nim=?, no_hp=?, last_update=?, is_allowed=? WHERE nim=?"
	_, err = helper.Db.Exec(query, nim, nim, no_hp, library.CurrTimestamp(), 1, nim)
	if err != nil {
		return err
	}

	// insert to tbl_user_payment
	query = "UPDATE tbl_user_payment SET bill=?, va_account=?, last_payment_date=? WHERE nim=?"
	_, err = helper.Db.Exec(query, tuition_fee, va_account, last_payment_date, nim)
	if err != nil {
		return err
	}

	return nil
}

func DeleteStudent(nim string) error {
	//update to tbl_user_student
	query := "DELETE FROM tbl_user_student WHERE nim=?"
	_, err := helper.Db.Exec(query, nim)
	if err != nil {
		return err
	}

	// delete tbl_user_event
	query = "DELETE FROM tbl_user_event WHERE nim=?"
	_, err = helper.Db.Exec(query, nim)
	if err != nil {
		return err
	}

	// insert to tbl_user_notif
	query = "DELETE FROM tbl_user_notif WHERE nim=?"
	_, err = helper.Db.Exec(query, nim)
	if err != nil {
		return err
	}

	// delete to tbl_user_payment
	query = "DELETE FROM tbl_user_payment WHERE nim=?"
	_, err = helper.Db.Exec(query, nim)
	if err != nil {
		return err
	}

	return nil
}

func GetStudentInfo() ([]map[string]interface{}, error) {
	query := `SELECT us.nim, us.name, un.no_hp, us.major, ue.id_event as class_id, e.title AS class_title FROM tbl_user_student us 
		JOIN tbl_user_notif un ON 
		us.nim = un.nim
		JOIN tbl_user_event ue ON
		us.nim = ue.nim
		JOIN tbl_event e ON
		ue.id_event = e.id`
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func GetDetailStudentInfo(nim string) ([]map[string]interface{}, error) {
	query := `SELECT us.nim, us.name, un.no_hp, us.major, ue.id_event as class_id, e.title AS class_title,
		up.bill, up.va_account, up.last_payment_date
	 FROM tbl_user_student us 
		JOIN tbl_user_notif un ON 
		us.nim = un.nim
		JOIN tbl_user_event ue ON
		us.nim = ue.nim
		JOIN tbl_event e ON
		ue.id_event = e.id 
		JOIN tbl_user_payment up on
		up.nim=us.nim
		WHERE us.nim= ?
		`
	data, err := GeneralSelectRows(query, nim)
	fmt.Println("data dari db", data)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func InsertCourseToDB(title string, description string, schedule string, day string, job_day string) (id int64, err error) {
	//insert to tbl_user_student
	query := "INSERT INTO tbl_event (title, description, schedule, job_every, event_day, id_event_type) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := helper.Db.Exec(query, title, description, schedule, job_day, day, 1)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func InsertToTableJob(job_name string, job_id string, id_event int64) error {
	//insert to tbl_job
	query := "INSERT INTO tbl_job (job_name, job_id, id_event) VALUES (?, ?, ?)"
	_, err := helper.Db.Exec(query, job_name, job_id, id_event)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCourseToDB(id int64, title string, description string, schedule string, day string, job_day string) error {
	//insert to tbl_user_student
	query := "UPDATE tbl_event SET title=?, description=?, schedule=?, job_every=?, event_day=?, id_event_type=? WHERE id=?"
	_, err := helper.Db.Exec(query, title, description, schedule, job_day, day, 1, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateJob(job_name string, job_id string, idEvent int64) error {
	//insert to tbl_user_student
	query := "UPDATE tbl_job SET job_name=?, job_id=? WHERE id_event=?"
	_, err := helper.Db.Exec(query, job_name, job_id, idEvent)
	if err != nil {
		return err
	}

	return nil
}

func GetJob(idEvent int64) (map[string]interface{}, error) {
	query := "SELECT * FROM tbl_job WHERE id_event = ?"
	data, err := GeneralSelect(query, idEvent)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func DeleteCourse(idEvent int64) error {
	// delete data from tbl_event, tbl_job, tbl_user_event
	// delete tbl_event
	query := "DELETE FROM tbl_event WHERE id=?"
	_, err := helper.Db.Exec(query, idEvent)
	if err != nil {
		return err
	}

	// delete tbl_job
	query = "DELETE FROM tbl_job WHERE id_event=?"
	_, err = helper.Db.Exec(query, idEvent)
	if err != nil {
		return err
	}

	// delete tbl_user_event
	query = "DELETE FROM tbl_user_event WHERE id_event=?"
	_, err = helper.Db.Exec(query, idEvent)
	if err != nil {
		return err
	}

	return nil
}

func InsertToTrxLog(idEvent int, count int, trxType string) error {
	//insert to tbl_job
	query := "INSERT INTO tbl_trx_log (id_event, user_success, trx_type, trx_date) VALUES (?, ?, ?, ?)"
	_, err := helper.Db.Exec(query, idEvent, count, trxType, library.CurrTimestamp())
	if err != nil {
		return err
	}
	return nil
}

func GetTrxLog() ([]map[string]interface{}, error) {
	query := "SELECT tl.*,e.title from tbl_trx_log tl JOIN tbl_event e ON e.id=tl.id_event WHERE trx_type='course' ORDER BY id DESC"
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func GetUserPayment() ([]map[string]interface{}, error) {
	query := `SELECT up.id, us.name, up.nim, up.bill, up.va_account, no_hp,
	DATE_FORMAT(up.last_payment_date, '%d/%m/%Y') AS last_payment_date,
	DATE_FORMAT(LAST_DAY(CURDATE()), '%d/%m/%Y') AS due_date
			FROM tbl_user_payment up
			JOIN tbl_user_student us ON up.nim=us.nim 
			JOIN tbl_user_notif un ON up.nim=un.nim
			WHERE MONTH(last_payment_date) < MONTH(CURDATE())`
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func GetUserPaymentByNim(nim string) (map[string]interface{}, error) {
	query := `SELECT up.id, us.name, up.nim, up.bill, up.va_account, no_hp,
	DATE_FORMAT(up.last_payment_date, '%d/%m/%Y') AS last_payment_date,
	DATE_FORMAT(LAST_DAY(CURDATE()), '%d/%m/%Y') AS due_date
			FROM tbl_user_payment up
			JOIN tbl_user_student us ON up.nim=us.nim 
			JOIN tbl_user_notif un ON up.nim=un.nim
			WHERE MONTH(last_payment_date) < MONTH(CURDATE()) AND us.nim=?`
	data, err := GeneralSelect(query, nim)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func GetTransactionInfo() ([]map[string]interface{}, error) {
	query := `SELECT et.name AS title, COUNT(*) as total, CONCAT(COUNT(*), ' hit/year')total_str FROM tbl_trx_log tl 
	JOIN tbl_event_type et
	ON et.trx_type = tl.trx_type
	WHERE YEAR(trx_date) = YEAR(CURDATE())
	GROUP BY tl.trx_type`
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func JobDetail() (map[string]interface{}, error) {
	query := `SELECT e.*,j.job_name,j.job_id FROM tbl_event e JOIN tbl_job j ON e.id=j.id_event WHERE id_event_type=3`
	data, err := GeneralSelect(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func GetGraduation() ([]map[string]interface{}, error) {
	query := `
	SELECT g.*, name, major FROM tbl_graduation g JOIN tbl_user_student us ON g.nim=us.nim
	`
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func GetNotRegisGraduation() ([]map[string]interface{}, error) {
	query := `
	SELECT us.name, us.major, un.no_hp FROM tbl_graduation g JOIN tbl_user_student us ON g.nim=us.nim 
	JOIN tbl_user_notif un ON g.nim=un.nim
	WHERE is_registered=0
	`
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func GetUserGraduation(nim string) (map[string]interface{}, error) {
	query := `
	SELECT us.name, us.major, un.no_hp FROM tbl_graduation g JOIN tbl_user_student us ON g.nim=us.nim 
	JOIN tbl_user_notif un ON g.nim=un.nim
	WHERE g.nim=?
	`
	data, err := GeneralSelect(query, nim)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil // Handle the case where no data was found for the given idEvent.
	}
	return data, nil
}

func GetThesis() ([]map[string]interface{}, error) {
	query := `SELECT t.*,name  FROM tbl_thesis t JOIN tbl_user_student us on us.nim=t.nim`
	data, err := GeneralSelectRows(query)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	return data, nil
}
