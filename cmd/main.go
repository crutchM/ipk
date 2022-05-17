package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"ipk"
	"ipk/pkg/handler"
	"ipk/pkg/repository"
	"ipk/pkg/service"
)

// @Title Ipk VSOKO app
// @Version 1.0

// @host localhost:8081
// @BasePath /

// @SecurityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
//происходит инициализация всех компонентов, редактировать ничего не надо, если надо поправить порт/строку подключения к бд- идем в config.yml
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("error in config")
	}
	//db, err := repository.NewPostgresDb(repository.Config{
	//	Host:     "localhost",
	//	Port:     "5432",
	//	Username: "postgres",
	//	Password: "postgres",
	//	DBName:   "postgres",
	//	SSLMode:  "disable",
	//})
	db, err := repository.NewPostgresDb(repository.Config{
		ConnectionRow: viper.GetString("connectionRow"),
	})
	if err != nil {
		logrus.Fatal(err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	hands := handler.NewHandler(services)
	srv := new(ipk.Server)
	if err := srv.Run(viper.GetString("port"), hands.InitRoutes()); err != nil {
		logrus.Fatal(err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
