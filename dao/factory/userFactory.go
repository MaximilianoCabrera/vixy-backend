package factory

import (
	"../../dao/interfaces"
	"../../dao/mysql"
	"log"
	//"github.com/MaximilianoCabrera/Vixy/dao/psql"
)

func UserFactoryDAO(e string) interfaces.UserDAO {
	var i interfaces.UserDAO
	switch e {
	//case "postgres":
	//	i = psql.UserImplPsql{}
	case "mysql":
		i = mysql.UserImplMysql{}
	default:
		log.Fatalf("El motor %s no está implementado", e)
		return nil
	}
	return i
}
