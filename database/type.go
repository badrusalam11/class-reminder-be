package database

type Event struct {
	Id             int
	Title          string
	Description    string
	Schedule       string
	IdEventType    int
	IsSpecificUser int
}

type EventUser struct {
	Username   string
	Notif_id   string
	Is_allowed int
}
