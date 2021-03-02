package main
import(
	"Go/filestore/first/db/mysql"
)
func main() {
	mysql.Init("Tb_User", true)
}
