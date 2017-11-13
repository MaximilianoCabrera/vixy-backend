package factory

import (
	"../../dao/interfaces"
	"../../dao/mysql"
	"log"
	//"github.com/MaximilianoCabrera/Vixy/dao/psql"
)

func EntradaFactoryDAO(e string) interfaces.EntradasDAO {
	var i interfaces.EntradasDAO
	switch e {
	//case "postgres":
	//	i = psql.UserImplPsql{}
	case "mysql":
		i = mysql.EntradaImplMysql{}
	default:
		log.Fatalf("El motor %s no está implementado", e)
		return nil
	}
	return i
}
