package connect

import (
	"fmt"
	"time"

	"crm/config"

	dbConfig "crm/internal/interactor/pkg/connect/postgres"
	"crm/internal/interactor/pkg/util"
	"crm/internal/interactor/pkg/util/log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func PostgresSQL() (db *gorm.DB, err error) {
	const dsn string = "host=%s port=%d user=%s dbname=%s sslmode=%s password=%s"
	pgConfig := dbConfig.Config{}
	pgConfig.DSN = util.PointerString(
		fmt.Sprintf(dsn, config.SourceHost, config.SourcePort, config.SourceUser, config.SourceDataBase,
			config.SourceSSLMode, config.SourcePassword))
	pgConfig.PreferSimpleProtocol = util.PointerBool(true)
	pgConfig.NowFunc = func() time.Time { return time.Now().UTC() }
	if gin.Mode() == "debug" {
		pgConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err = pgConfig.Connect()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, nil
}
