package factory

import (
	"../../dao/interfaces"
	"../../dao/mysql"
	"log"
)

func UsuarioFactoryDAO(e string) interfaces.UsuarioDAO {
	var i interfaces.UsuarioDAO
	switch e {
	//case "postgres":
	//	i = psql.UserImplPsql{}
	case "mysql":
		i = mysql.UsuarioImplMysql{}
	default:
		log.Fatalf("El motor %s no está implementado", e)
		return nil
	}
	return i
}
