package factory

import (
	"../../dao/interfaces"
	"../../dao/mysql"
)

func UsuarioFactoryDAO(e string) interfaces.UsuarioDAO {
	var i interfaces.UsuarioDAO
	i = mysql.UsuarioImplMysql{}
	return i
}
