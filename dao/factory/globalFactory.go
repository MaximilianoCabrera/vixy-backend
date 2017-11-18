package factory

import (
"../../dao/interfaces"
"../../dao/mysql"
)

func GlobalFactoryDAO(e string) interfaces.GlobalDAO {
	var i interfaces.GlobalDAO
	i = mysql.GlobalImplMysql{}
	return i
}

