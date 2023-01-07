package server

type customer struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// type reservation struct {
// 	TableId  int `json:"tableid"`
// 	Day      int `json:"day"`
// 	Hour     int `json:"hour"`
// 	Duration int `json:"duration"`
// 	Persons  int `json:"persons"`
// }
