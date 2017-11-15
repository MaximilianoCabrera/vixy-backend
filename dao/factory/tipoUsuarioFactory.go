package factory

import (
	"../../dao/interfaces"
	"../../dao/mysql"
	"log"
)

func TipoUsuarioFactoryDAO(e string) interfaces.TipoUsuarioDAO {
	var i interfaces.TipoUsuarioDAO
	switch e {
	//case "postgres":
	//	i = psql.UserImplPsql{}
	case "mysql":
		i = mysql.TipoUsuarioImplMysql{}
	default:
		log.Fatalf("El motor %s no está implementado", e)
		return nil
	}
	return i
}