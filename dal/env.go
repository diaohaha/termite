package dal

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

type Environment struct {
	//only allow int64,float64,string,bool,struct
	App           string
	IsDebug       bool   `env:"DEBUG"`
	MicroRegistry string `env:"MICRO_REGISTRY_HOST"`
	Timeout       time.Duration

	Sentry_Dsn string `env:"SENTRY_DSN"`

	Kafka_Termite_event struct {
		Host  string `env:"KAFKA_HOST"`
		Topic string `env:"KAFKA_TOPIC"`
	}

	DB_default struct {
		Host     string `env:"DB_HOST"`
		Port     int64  `env:"DB_PORT"`
		Username string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Name     string `env:"DB_NAME"`
	}

	DB_read struct {
		Host     string `env:"DB_READ_HOST"`
		Port     int64  `env:"DB_READ_PORT"`
		Username string `env:"DB_READ_USER"`
		Password string `env:"DB_READ_PASSWORD"`
		Name     string `env:"DB_READ_NAME"`
	}

	Redis struct {
		Addr     string `env:"REDIS_ADDR"`
		Password string `env:"REDIS_PASSWORD"`
		Db       int    `env:"REDIS_DB"`
	}
}

func load_struct(struct_object reflect.Value) {
	filed_num := struct_object.Type().NumField()
	for i := 0; i < filed_num; i++ {
		if struct_object.Field(i).Type().Kind() == reflect.Struct {
			load_struct(struct_object.Field(i))
			continue
		}
		tag := struct_object.Type().Field(i).Tag.Get("env")
		if tag == "" {
			continue
		}
		env_value := os.Getenv(tag)
		if env_value == "" {
			continue
		}
		switch struct_object.Field(i).Type().Kind() {
		case reflect.Bool:
			env_v, err := strconv.ParseBool(env_value)
			if err != nil {
				continue
			}
			struct_object.Field(i).SetBool(env_v)
		case reflect.Int64:
			env_v, err := strconv.ParseInt(env_value, 10, 64)
			if err != nil {
				continue
			}
			struct_object.Field(i).SetInt(env_v)

		case reflect.Int:
			env_v, err := strconv.ParseInt(env_value, 10, 64)
			if err != nil {
				continue
			}
			struct_object.Field(i).SetInt(env_v)
		case reflect.String:
			struct_object.Field(i).SetString(env_value)
		case reflect.Float64:
			env_v, err := strconv.ParseFloat(env_value, 64)
			if err != nil {
				continue
			}
			struct_object.Field(i).SetFloat(env_v)
		}
	}
}

var Env = Environment{}

func init() {
	println("exporting env")
	load_struct(reflect.ValueOf(&Env).Elem())
	println("exported env")
	println(Env.DB_default.Username)
	println(Env.Kafka_Termite_event.Topic)
	log.Println(Env.Sentry_Dsn)
}
