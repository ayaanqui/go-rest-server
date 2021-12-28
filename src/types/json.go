package types

type Response struct {
	Message string `json:"message"`
}

type Person struct {
	Name string `json:"name"`
	Age uint `json:"age"`
}

type Result struct {
	Data interface{} `json:"data"`
}

type DbConnection struct {
	DbName string `json:"db_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreatePost struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Summary string `json:"summary"`
}

type CreateUser struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Tokens struct {
	JwtKey string `json:"jwt_key"`
}