package controllers

type Accounts struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Games struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	MaxPlayer int    `json:"maxPlayer"`
}

type Rooms struct {
	Id       int    `json:"id"`
	RoomName string `json:"roomName"`
	IdGame   Games  `json:"idGame"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Rooms  `json:"data"`
}

type RoomsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}

type DetailedRoomsResponse struct {
	Data []DetailedRooms `json:"Rooms"`
}

type DetailedRooms struct {
	Participant Participants `json:"participants"`
	Room        Rooms        `json:"rooms"`
}

type Participants struct {
	Id        int      `json:"id"`
	IdRoom    Rooms    `json:"idRRoom"`
	IdAccount Accounts `json:"idAccount"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
