package seed

import (
	"github.com/bxcodec/faker/v4"
	"log"
	"server-v2/config"
	"server-v2/models"
)

func main() {
	var users []models.UserInput

	for i := 0; i < 10; i++ {
		user := models.User{
			Username: faker.Username(),
			Password: faker.Password(),
			Email:    faker.Email(),
		}

		if err := config.DB.Create(&user).Error; err != nil {
			log.Println(err)
		}

		users = append(users, user)
	}

}
