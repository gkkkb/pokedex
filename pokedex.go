package pokedex

import (
	"os"
	"sync"

	"github.com/gkkkb/pokedex/pkg/mysql"
	"github.com/gkkkb/pokedex/pkg/storage"
	"github.com/gkkkb/pokedex/pkg/telolet"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

type Pokedex struct {
	DB      *sqlx.DB
	Storage storage.StorageInterface
	Logger  *logrus.Logger
}

var pokedex *Pokedex
var once sync.Once

func GetInstance() *Pokedex {
	once.Do(func() {
		gotenv.Load(os.Getenv("GOPATH") + "/src/github.com/bukalapak/pokedex/.env")

		environment := os.Getenv("ENV")
		db := mysql.Init()

		logger := initLogger()

		var store storage.StorageInterface
		var err error
		if environment == "development" {
			store, err = storage.InitLocal()
		} else {
			store, err = storage.InitAWS2()
		}

		if err != nil {
			panic(err)
		}

		pokedex = &Pokedex{DB: db, Storage: store, Telolet: teloletClient, Logger: logger}
	})

	return pokedex
}

func initLogger() *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(os.Getenv("LOGGING_LEVEL"))
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.Level = level

	return logger
}
