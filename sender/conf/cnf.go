package conf

import (
	"database/sql"
	"io/ioutil"

	"github.com/julienschmidt/httprouter"
	"github.com/olebedev/config"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Ip          string
	Port        string
	DbPath      string
	SQLInitPath string
}

type Store struct {
	Db     *sql.DB
	Log    *logrus.Logger
	Routes *httprouter.Router
}

func New() *Store {
	s := &Store{
		Log:    logrus.New(),
		Routes: httprouter.New(),
	}
	return s
}

func (s *Store) InitDB(c *Config) error {
	db, err := sql.Open("sqlite3", c.DbPath)
	if err != nil {
		return err
	}

	sql, err := ioutil.ReadFile(c.SQLInitPath)
	if err != nil {
		panic(err)
	}

	_, err = db.Query(string(sql))
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.Db = db
	return nil
}

func Cnf() (*Config, error) {
	file, err := ioutil.ReadFile("conf/cnf.yml")
	if err != nil {
		return nil, err
	}
	yamlString := string(file)

	cfg, err := config.ParseYaml(yamlString)
	if err != nil {
		return nil, err
	}
	SQLInitPath, err := cfg.String("sql_init_path")
	if err != nil {
		return nil, err
	}
	Ip, err := cfg.String("ip")
	if err != nil {
		return nil, err
	}
	Port, err := cfg.String("port")
	if err != nil {
		return nil, err
	}
	Db, err := cfg.String("db")
	if err != nil {
		return nil, err
	}
	return &Config{
		Ip:          Ip,
		Port:        Port,
		DbPath:      Db,
		SQLInitPath: SQLInitPath,
	}, nil
}
