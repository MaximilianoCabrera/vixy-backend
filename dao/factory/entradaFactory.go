package factory

import (
	"../../dao/interfaces"
	"../../dao/mysql"
)

func EntradaFactoryDAO(e string) interfaces.EntradaDAO {
	var i interfaces.EntradaDAO
	i = mysql.EntradaImplMysql{}

	return i
}
