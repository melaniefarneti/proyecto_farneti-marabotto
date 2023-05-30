// Función para obtener un usuario por su nombre de usuario
func GetUserByUsername(username string) (*User, error) {
	// Realizar la consulta a la base de datos para obtener los datos del usuario
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := db.QueryRow(query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// El usuario no existe
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// Función para crear un usuario
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Leer los datos del usuario del cuerpo de la solicitud
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verificar si el usuario ya existe en la base de datos
	existingUser, err := GetUserByUsername(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]interface{}{
			"message": "El usuario ya existe",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generar el hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Almacenar los datos del usuario en la base de datos
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err = db.Exec(query, user.Username, hashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Usuario creado exitosamente",
	}
	json.NewEncoder(w).Encode(response)
}
