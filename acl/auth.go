package acl

type Auth struct {
	Is     int               `json:"i"`
	Data   map[string]string `json:"d"`
	Access []Access          `json:"a"`
}

type Access struct {
	Method     string `json:"m"`
	Controller string `json:"c"`
	Action     string `json:"a"`
}
