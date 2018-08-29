package db

import (
	"fmt"
	"sync"
	"wyatt/config"
	"wyatt/util"

	"f.in/v/utils"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

const (
	REDIS = "redis"
	MYSQL = "mysql"
	MGO   = "mgo"
)

var (
	once     sync.Once
	instance *singleton
	conf     = config.GetConfig()
)

type (
	Driver string
)
type singleton struct {
	services *sync.Map
}

//获取mysql
func GetMysqlDB() *gorm.DB {
	return newInstance(MYSQL).(*gorm.DB)
}

//获取redis
func GetRedisDB() *redis.Client {
	return newInstance(REDIS).(*redis.Client)
}

//获取mgo
func GetMgoDB() *mgo.Session {
	return newInstance(MGO).(*mgo.Session)
}

func newInstance(driver Driver) interface{} {
	s := getInstance()
	if val, ok := s.services.Load(driver); ok {
		return val
	}
	var r interface{}
	switch driver {
	case REDIS:
		r, _ = s.getOrSetMap(REDIS, newRidis())
	case MYSQL:
		r, _ = s.getOrSetMap(MYSQL, newMysql())
	case MGO:
		r, _ = s.getOrSetMap(MGO, newMGO())
	default:
	}
	return r
}

func getInstance() *singleton {
	if instance == nil {
		once.Do(func() {
			instance = &singleton{services: &sync.Map{}}
		})
		instance.getOrSetMap(REDIS, newRidis())
	}
	return instance
}

func (s *singleton) getOrSetMap(name Driver, service interface{}) (interface{}, bool) {
	return s.services.LoadOrStore(name, service)
}

func newRidis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Redis.Host, conf.Redis.Port),
		Password: conf.Redis.Pass,
		DB:       utils.On(conf.Debug, 1, 0).(int),
		PoolSize: conf.Redis.Pool,
	})
	client.Ping().Result()
	return client
}

func newMysql() *gorm.DB {
	mysql, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			conf.Mysql.User,
			conf.Mysql.Pass,
			conf.Mysql.Host,
			conf.Mysql.Port,
			conf.Mysql.DbName,
			conf.Mysql.Charset))

	if err != nil {
		util.LoggerError(err)
	}
	mysql.DB().SetMaxIdleConns(2)
	mysql.DB().SetMaxOpenConns(conf.Mysql.Pool)
	return mysql
}

func newMGO() *mgo.Session {
	var session *mgo.Session
	var err error

	if conf.Debug {
		session, err = mgo.Dial(fmt.Sprintf("%s:%s", conf.MGO.Host, conf.MGO.Port))
	} else {
		session, err = mgo.Dial(fmt.Sprintf("%s:%s@%s:%s",
			conf.MGO.User,
			conf.MGO.Pass,
			conf.MGO.Host,
			conf.MGO.Port))

		if err != nil {
			util.LoggerError(err)
		}
	}
	session.SetMode(mgo.Monotonic, true)
	session.SetPoolLimit(conf.MGO.Pool)
	return session.Clone()
}
