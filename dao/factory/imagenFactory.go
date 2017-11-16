package factory

import (
	"../../dao/interfaces"
	"../../dao/mysql"
)

func ImagenFactoryDAO(e string) interfaces.ImagenDAO {
	var i interfaces.ImagenDAO
	i = mysql.ImagenImplMysql{}

	return i
}

