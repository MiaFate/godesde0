package users

import (
	"fmt"
	"time"

	"github.com/miafate/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Mia", time.Now(), true)
	fmt.Println(u)
}
