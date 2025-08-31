package services

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type ConfigService interface {
	Get(key string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetFloat64(key string) float64
	GetStringSlice(key string) []string
	Set(key string, value interface{})
	SetDefault(key string, value interface{})
	IsSet(key string) bool
	ReadConfig() error
	WriteConfig() error
	SetConfigName(name string)
	SetConfigType(configType string)
	AddConfigPath(path string)
	InitializeConfig() error
}

type ViperConfigService struct {
	viper *viper.Viper
}

func NewConfigService() ConfigService {
	service := &ViperConfigService{
		viper: viper.New(),
	}
	service.InitializeConfig()
	return service
}

func (c *ViperConfigService) Get(key string) interface{} {
	return c.viper.Get(key)
}

func (c *ViperConfigService) GetString(key string) string {
	return c.viper.GetString(key)
}

func (c *ViperConfigService) GetBool(key string) bool {
	return c.viper.GetBool(key)
}

func (c *ViperConfigService) GetInt(key string) int {
	return c.viper.GetInt(key)
}

func (c *ViperConfigService) GetFloat64(key string) float64 {
	return c.viper.GetFloat64(key)
}

func (c *ViperConfigService) GetStringSlice(key string) []string {
	return c.viper.GetStringSlice(key)
}

func (c *ViperConfigService) Set(key string, value interface{}) {
	c.viper.Set(key, value)
}

func (c *ViperConfigService) SetDefault(key string, value interface{}) {
	c.viper.SetDefault(key, value)
}

func (c *ViperConfigService) IsSet(key string) bool {
	return c.viper.IsSet(key)
}

func (c *ViperConfigService) ReadConfig() error {
	return c.viper.ReadInConfig()
}

func (c *ViperConfigService) WriteConfig() error {
	return c.viper.WriteConfig()
}

func (c *ViperConfigService) SetConfigName(name string) {
	c.viper.SetConfigName(name)
}

func (c *ViperConfigService) SetConfigType(configType string) {
	c.viper.SetConfigType(configType)
}

func (c *ViperConfigService) AddConfigPath(path string) {
	c.viper.AddConfigPath(path)
}

func (c *ViperConfigService) InitializeConfig() error {
	// Set config name and type
	c.viper.SetConfigName("git-pitch")
	c.viper.SetConfigType("yaml")

	// Set environment variable prefix
	c.viper.SetEnvPrefix("GIT_PITCH")
	c.viper.AutomaticEnv()

	// Add configuration paths in order of precedence
	// 1. Current working directory's .git folder
	if cwd, err := os.Getwd(); err == nil {
		gitDir := filepath.Join(cwd, ".git")
		if _, err := os.Stat(gitDir); err == nil {
			c.viper.AddConfigPath(gitDir)
		}
	}

	// 2. User's home directory
	if homeDir, err := os.UserHomeDir(); err == nil {
		c.viper.AddConfigPath(homeDir)
	}

	// 3. System-wide config directory
	c.viper.AddConfigPath("/etc/git-pitch")

	// 4. Current directory as fallback
	c.viper.AddConfigPath(".")

	// Set some default values
	c.setDefaults()

	// Try to read config file (it's okay if it doesn't exist)
	if err := c.viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error was produced
			return err
		}
		// Config file not found; ignore error
	}

	return nil
}

func (c *ViperConfigService) setDefaults() {
	// Add your default configuration values here
	c.viper.SetDefault("debug", false)
	c.viper.SetDefault("timeout", 30)
	c.viper.SetDefault("output_format", "text")
}