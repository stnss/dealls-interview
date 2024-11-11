package appctx

import (
	"fmt"
	"github.com/stnss/dealls-interview/pkg/file"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
	"time"
)

const (
	configPath = "config/"
)

var (
	cfgOnce sync.Once
	_cfg    *Config
)

// FileReader is a function type for reading files
type FileReader func(string) ([]byte, error)

// YAMLUnmarshaler is a function type for unmarshaling YAML content
type YAMLUnmarshaler func([]byte, any) error

type (
	Config struct {
		App     App        `yaml:"app" json:"app"`
		Logger  Logger     `yaml:"log" json:"log"`
		DBWrite *Database  `yaml:"db_write" json:"db_write"`
		DBRead  *Database  `yaml:"db_read" json:"db_read"`
		Redis   *RedisConf `yaml:"redis" json:"redis"`
		RSAKey  RSA        `yaml:"rsa" json:"rsa"`
		JWT     JWT        `yaml:"jwt" json:"jwt"`
	}

	App struct {
		Name         string        `yaml:"name" json:"name"`
		Port         int           `yaml:"port" json:"port"`
		Debug        bool          `yaml:"debug" json:"debug"`
		Timezone     string        `yaml:"timezone" json:"timezone"`
		Env          string        `yaml:"env" json:"env"`
		ReadTimeout  time.Duration `yaml:"read_timeout" json:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout" json:"write_timeout"`
	}

	Logger struct {
		Level string `yaml:"level" json:"level"`
	}

	// Database configuration structure
	Database struct {
		Name         string        `yaml:"name" json:"name"`
		User         string        `yaml:"user" json:"user"`
		Pass         string        `yaml:"pass" json:"pass"`
		Host         string        `yaml:"host" json:"host"`
		Port         int           `yaml:"port" json:"port"`
		MaxOpen      int           `yaml:"max_open" json:"max_open"`
		MaxIdle      int           `yaml:"max_idle" json:"max_idle"`
		DialTimeout  time.Duration `yaml:"dial_timeout" json:"dial_timeout"`
		MaxLifeTime  time.Duration `yaml:"life_time" json:"max_life_time"`
		ReadTimeout  time.Duration `yaml:"read_timeout" json:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout" json:"write_timeout"`
		Charset      string        `yaml:"charset" json:"charset"`
		Driver       string        `yaml:"driver" json:"driver"`
		Timezone     string        `yaml:"timezone" json:"timezone"`
	}

	// RedisConf general config redis
	RedisConf struct {
		Hosts string `yaml:"host"`
		DB    int    `yaml:"db"`

		// ClientName will execute the `CLIENT SETNAME ClientName` command for each conn.
		ClientName string `yaml:"client_name"`
		// specify 2 for RESP 2 or 3 for RESP 3
		Protocol         int    `yaml:"protocol"`
		Username         string `yaml:"username"`
		Password         string `yaml:"password"`
		SentinelUsername string `yaml:"sentinel_username"`
		SentinelPassword string `yaml:"sentinel_password"`

		MaxRetries      int           `yaml:"max_retries"`
		MinRetryBackoff time.Duration `yaml:"min_retry_backoff"`
		MaxRetryBackoff time.Duration `yaml:"max_retry_backoff"`

		DialTimeout           time.Duration `yaml:"dial_timeout"`
		ReadTimeout           time.Duration `yaml:"read_timeout"`
		WriteTimeout          time.Duration `yaml:"write_timeout"`
		ContextTimeoutEnabled bool          `yaml:"context_timeout_enabled"`

		// PoolFIFO uses FIFO mode for each node connection pool GET/PUT (default LIFO).
		PoolFIFO        bool          `yaml:"pool_fifo"`
		PoolSize        int           `yaml:"pool_size"`
		PoolTimeout     time.Duration `yaml:"pool_timeout"`
		MinIdleConn     int           `yaml:"min_idle_conn"`
		MaxIdleConn     int           `yaml:"max_idle_conn"`
		MaxActiveConn   int           `yaml:"max_active_conn"`
		ConnIdleTime    time.Duration `yaml:"conn_idle_time"`
		ConnMaxLifeTime time.Duration `yaml:"conn_max_life_time"`

		//IdleFrequencyCheck int `yaml:"idle_frequency_check"`

		// Only cluster clients.
		RouteByLatency bool `yaml:"route_by_latency"`
		RouteRandomly  bool `yaml:"route_randomly"`
		ReadOnly       bool `yaml:"read_only"`
		MaxRedirect    int  `yaml:"max_redirect"`

		ClusterMode        bool `yaml:"cluster_mode"`
		TLSEnable          bool `yaml:"tls_enable"`
		InsecureSkipVerify bool `yaml:"insecure_skip_verify"`

		// The sentinel master name.
		// Only failover clients.
		MasterName string `yaml:"master_name"`
	}

	RSA struct {
		PrivateKey string `yaml:"private_key"`
		PublicKey  string `yaml:"public_key"`
	}

	JWT struct {
		AccessSecret  string        `yaml:"access_secret"`
		RefreshSecret string        `yaml:"refresh_secret"`
		ExpiredTime   time.Duration `yaml:"expired_time"`
	}
)

func NewConfig() *Config {
	cfgPath := []string{configPath}
	cfgOnce.Do(func() {
		c, err := readConfig("app.yaml", readFileFunc, yamlReadFunc, cfgPath...)
		if err != nil {
			log.Fatal("failed to load config")
		}
		_cfg = c
	})
	return _cfg
}

func readConfig(configFile string, fileReader file.ReadFileFunc, yamlReader file.YAMLUnmarshalFunc, configPaths ...string) (*Config, error) {
	var (
		cfg  *Config
		errs []error
	)

	for _, path := range configPaths {
		cfgPath := fmt.Sprint(path, configFile)
		if err := file.ReadFromYAML(cfgPath, &cfg, fileReader, yamlReader); err != nil {
			errs = append(errs, fmt.Errorf("file %s error %s", cfgPath, err.Error()))
			continue
		}
		break
	}

	if cfg == nil {
		return nil, fmt.Errorf("file config parse error %v", errs)
	}

	return cfg, nil
}

func readFileFunc(s string) ([]byte, error) {
	return os.ReadFile(s)
}

func yamlReadFunc(bytes []byte, a any) error {
	return yaml.Unmarshal(bytes, a)
}
