package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * FROM rooms"

	id := r.URL.Query()["id"]
	roomName := r.URL.Query()["room_name"]
	if id != nil {
		fmt.Println(id[0])
		query += " WHERE id ='" + id[0] + "'"
	}
	if roomName != nil {
		if id[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " room_name='" + roomName[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var room Rooms
	var rooms []Rooms
	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.RoomName, &room.IdGame); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if len(rooms) < 5 {
		var response RoomsResponse
		response.Status = 200
		response.Message = "Success"
		response.Data = rooms
		json.NewEncoder(w).Encode(response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Error Array Size Not Correct"
		json.NewEncoder(w).Encode(response)
	}
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT participants.id, participants.id_room, participants.id_account FROM Participants INNER JOIN rooms ON rooms.id = participants.id_room"

	id := r.URL.Query()["id"]

	if id != nil {
		query += " WHERE rooms.id='" + id[0] + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var detailedRoom DetailedRooms
	var detailedRooms []DetailedRooms

	for rows.Next() {
		if err := rows.Scan(&detailedRoom.Participant.Id, &detailedRoom.Participant.IdRoom, &detailedRoom.Participant.IdAccount); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again")
			return
		} else {
			detailedRooms = append(detailedRooms, detailedRoom)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if len(detailedRooms) < 5 {
		var response DetailedRoomsResponse
		response.Data = detailedRooms
		json.NewEncoder(w).Encode(response)
	} else {
		var response ErrorResponse
		response.Status = 400
		response.Message = "Error Array Size Not Correct"
		json.NewEncoder(w).Encode(response)
	}

}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		sendErrorResponse(w, "failed")
		return
	}
	id := r.Form.Get("id")
	roomName := r.Form.Get("room_name")
	idGame, _ := strconv.Atoi(r.Form.Get("id_game"))

	_, errQuery := db.Exec("INSERT INTO rooms(id, room_name, id_game)values (?,?,?)",
		id,
		roomName,
		idGame,
	)
	var response RoomResponse
	response.Data.RoomName = roomName

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Insert Failed"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {

	db := Connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	Id := vars["id"]
	fmt.Println(Id)

	_, errQuery := db.Exec("DELETE FROM rooms WHERE id=?",
		Id,
	)
	var response RoomResponse
	if errQuery == nil {
		response.Status = 200
		response.Message = "Berhasil keluar room"
	} else {
		fmt.Println(errQuery)
		response.Status = 400
		response.Message = "Gagal keluar room"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response ErrorResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
