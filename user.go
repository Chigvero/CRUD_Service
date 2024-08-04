package todo

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"` //Валидирует наличие данных полей в теле запроса
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
