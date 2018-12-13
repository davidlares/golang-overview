package main
import "fmt"

type User interface{
  Permissions() int
  Name() string
}

type Admin struct {
  name string
}

func (this Admin) Permissions() int {
  return 5
}

func (this Admin) Name() string {
  return this.name
}

func auth(user User) string{
  if user.Permissions() > 4 {
    return user.Name() + " got admin permissions"
  }
  return ""
}

func main(){
  admin := Admin{"David"}
  fmt.Println(auth(admin))
}
