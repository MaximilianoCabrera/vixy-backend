package factory

import (
	"../mysql"
	"../interfaces"
	"log"

)

func ImagenFactoryDAO(img string) interfaces.ImagenDAO {
	var i interfaces.ImagenDAO
	switch img {
	//case "postgres":
	//	i = psql.UserImplPsql{}
	case "mysql":
		i = mysql.ImagenImplMysql{}
	default:
		log.Fatalf("El motor %s no está implementado", img)
		return nil
	}
	return i
}

