package dtos

type RegisterRequest struct { //guarda la info de REGISTRO q llefo del frontend del usuario
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct { //guarda la info de LOGIN q llego del frontend
	Email    string `json:"email"`
	Password string `json:"password"`
}
