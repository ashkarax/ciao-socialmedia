package config

import (
	"github.com/spf13/viper"
)

type ApiKey struct {
	Key string `mapstructure:"API_KEY"`
}

type PortManager struct {
	RunnerPort string `mapstructure:"PORTNO"`
}

type DataBase struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBName     string `mapstructure:"DBNAME"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	DBHost     string `mapstructure:"DBHOST"`
	DBPort     string `mapstructure:"DBPORT"`
}

type Token struct {
	AdminSecurityKey      string `mapstructure:"ADMIN_TOKENKEY"`
	RestaurantSecurityKey string `mapstructure:"RESTAURANT_TOKENKEY"`
	UserSecurityKey       string `mapstructure:"USER_TOKENKEY"`
	TempVerificationKey   string `mapstructure:"TEMPERVERY_TOKENKEY"`
}

type Smtp struct {
	SmtpSender   string `mapstructure:"SMTP_SENDER"`
	SmtpPassword string `mapstructure:"SMTP_APPKEY"`
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpPort     string `mapstructure:"SMTP_PORT"`
}

type AWS struct {
	Region     string `mapstructure:"AWS_REGION"`
	AccessKey  string `mapstructure:"AWS_ACCESS_KEY_ID"`
	SecrectKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	Endpoint   string `mapstructure:"AWS_ENDPOINT"`
}

type Auth2o struct {
	ClientId                string `mapstructure:"CLIENT_ID"`
	ProjectId               string `mapstructure:"PROJECT_ID"`
	AuthUri                 string `mapstructure:"AUTH_URI"`
	TokenUri                string `mapstructure:"TOKEN_URI"`
	AuthProviderX509CentUrl string `mapstructure:"AUTH_PROVIDER_X509_CENT_URL"`
}

type Config struct {
	ApiKey   ApiKey
	PortMngr PortManager
	DB       DataBase
	Token    Token
	Smtp     Smtp
	AwsS3    AWS
	Auth     Auth2o
}

func LoadConfig() (*Config, error) {
	var portmngr PortManager
	var db DataBase
	var token Token
	var smtp Smtp
	var apikey ApiKey
	var awsS3 AWS
	var Auth Auth2o

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&portmngr)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&db)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&token)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&smtp)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&apikey)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&awsS3)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&Auth)
	if err != nil {
		return nil, err
	}

	config := Config{ApiKey: apikey, PortMngr: portmngr, DB: db, Token: token, Smtp: smtp, AwsS3: awsS3, Auth: Auth}
	return &config, nil

}
