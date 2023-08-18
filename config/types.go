package config

type server struct {
	Secret  []byte
	Version string
	Name    string
}

type snowflake struct {
	WorkerID      int64 `yaml:"worker-id"`
	DatancenterID int64 `yaml:"datancenter-id"`
}

type service struct {
	Name     string
	AddrList []string
	LB       bool `yaml:"load-balance"`
}

type mySQL struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type etcd struct {
	Addr string
}

type config struct {
	Server    server
	Snowflake snowflake
	MySQL     mySQL
	Etcd      etcd
	RabbitMQ  rabbitMQ
	Redis     redis
	OSS       oss
}

type rabbitMQ struct {
	Addr     string
	Username string
	Password string
}

type redis struct {
	Addr     string
	Password string
}

type oss struct {
	Endpoint        string
	AccessKeyID     string `yaml:"accessKey-id"`
	AccessKeySecret string `yaml:"accessKey-secret"`
	BucketName      string
	MainDirectory   string `yaml:"main-directory"`
	DomainName      string `yaml:"domain-name"`
}
