package backendgis

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

// func TestUpdateGetData(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "healhero_db")
// 	datagedung := GetAllBangunanLineString(mconn, "healhero_db")
// 	fmt.Println(datagedung)
// }

func TestGeneratePasswordHash(t *testing.T) {
	password := "farhaniki45"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("iki12345", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gischapter5")
	var userdata User
	userdata.Username = "iki12345"
	userdata.Password = "farhaniki45"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gischapter5")
	var userdata User
	userdata.Username = "iki12345"
	userdata.Password = "farhaniki45"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

func TestInsertUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gischapter5")
	var userdata User
	userdata.Username = "iki12345"
	userdata.Password = "farhaniki45"

	nama := InsertUser(mconn, "user", userdata)
	fmt.Println(nama)
}
