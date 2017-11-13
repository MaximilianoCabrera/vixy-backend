package factory

import (
	"github.com/MaximilianoCabrera/Vixy/dao/interfaces"
	//"github.com/MaximilianoCabrera/Vixy/dao/psql"
	"github.com/MaximilianoCabrera/Vixy/dao/mysql"
	"log"
)

func FactoryDAO(e string) interfaces.UserDAO {
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
