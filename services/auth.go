package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ItKarma/idocks/models"
	"github.com/ItKarma/idocks/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// função de registro de usuario
func RegisterUser(db *mongo.Collection, email, password, company, cnpj string) error {
	//verificando se o usuario existe
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.User
	err := db.FindOne(ctx, bson.M{"email": email}).Decode(&existingUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Usuário não encontrado, podemos continuar com o registro
		} else {
			// Outro tipo de erro
			return err
		}
	} else {
		// Se o usuário já existe
		return errors.New("usuário já registrado")
	}

	// hash da senha
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	//inserir o novo usuario

	newCompany := models.Company{
		Nome: company,
		CNPJ: cnpj,
	}
	newUser := models.User{
		Email:    email,
		Password: hashedPassword,
		Company:  newCompany,
	}

	_, err = db.InsertOne(ctx, newUser)
	return err

}

// função de autenticação do usuario

func LoginUser(db *mongo.Collection, email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// buscar usuario
	var user models.User
	err := db.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("usuário ou senha inválidos")
	}

	// verificar a senha
	if !utils.VerifyPassword(user.Password, password) {
		return "", errors.New("usuário ou senha inválidos")
	}

	return utils.GerenateToken(user.ID.Hex())
}
