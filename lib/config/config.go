package config

import "github.com/spf13/viper"

var AddConfigPath = viper.AddConfigPath // func AddConfigPath(in string)
var AddRemoteProvider = viper.AddRemoteProvider // func AddRemoteProvider(provider, endpoint, path string) error
// func AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
// func AllKeys() []string
// func AllSettings() map[string]interface{}
// func AutomaticEnv()
// func BindEnv(input ...string) error
// func BindFlagValue(key string, flag FlagValue) error
// func BindFlagValues(flags FlagValueSet) error
// func BindPFlag(key string, flag *pflag.Flag) error
// func BindPFlags(flags *pflag.FlagSet) error
// func ClearConfigSettings()
// func ClearPFlag(key string)
// func ConfigFileUsed() string
// func Debug()
// func Desc(key string) (string, UseLevel, int)
// func Filesys() afero.Fs
// func Get(key string) interface{}
// func GetBool(key string) bool
// func GetDuration(key string) time.Duration
// func GetFloat64(key string) float64
var GetInt = viper.GetInt // func GetInt(key string) int
// func GetInt64(key string) int64
// func GetSizeInBytes(key string) uint
var GetString = viper.GetString // func GetString(key string) string
// func GetStringMap(key string) map[string]interface{}
// func GetStringMapString(key string) map[string]string
// func GetStringMapStringSlice(key string) map[string][]string
// func GetStringSlice(key string) []string
// func GetTime(key string) time.Time
// func InConfig(key string) bool
// func IsSet(key string) bool
// func MergeConfig(in io.Reader) error
// func MergeInConfig() error
// func OnConfigChange(run func(in fsnotify.Event))
// func ReadConfig(in io.Reader) error
// func ReadInConfig() error
// func ReadRemoteConfig() error
// func RegisterAlias(alias string, key string)
// func Reset()
// func Set(key string, value interface{})
// func SetConfigFile(in string)
// func SetConfigName(in string)
// func SetConfigType(in string)
var SetDefault = viper.SetDefault // func SetDefault(key string, value interface{})
// func SetDesc(key string, desc string, useLevel UseLevel, useScope int)
// func SetEnvKeyReplacer(r *strings.Replacer)
// func SetEnvPrefix(in string)
// func SetFilesys(fs afero.Fs)
// func SetFs(fs afero.Fs)
// func SetPFlags(flags *pflag.FlagSet) (err error)
// func SetTypeByDefaultValue(enable bool)
// func Unmarshal(rawVal interface{}) error
// func UnmarshalKey(key string, rawVal interface{}) error
// func WatchConfig()
// func WatchRemoteConfig() error
// type ConfigFileNotFoundError
// func (fnfe ConfigFileNotFoundError) Error() string
// type ConfigParseError
// func (pe ConfigParseError) Error() string
// type FlagValue
// type FlagValueSet
// type RemoteConfigError
// func (rce RemoteConfigError) Error() string
// type RemoteProvider
// type UnsupportedConfigError
// func (str UnsupportedConfigError) Error() string
// type UnsupportedRemoteProviderError
// func (str UnsupportedRemoteProviderError) Error() string
// type UseLevel
// func UseLevelString2UseLevel(s string) UseLevel
// func (ul UseLevel) String() string
// type Viper
// func GetSingleton() *Viper
var GetConfig = viper.GetViper // func GetViper() *Viper
// func New() *Viper
// func Sub(key string) *Viper
// func (v *Viper) AddConfigPath(in string)
// func (v *Viper) AddRemoteProvider(provider, endpoint, path string) error
// func (v *Viper) AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
// func (v *Viper) AllKeys() []string
// func (v *Viper) AllSettings() map[string]interface{}
// func (v *Viper) AutomaticEnv()
// func (v *Viper) BindEnv(input ...string) error
// func (v *Viper) BindFlagValue(key string, flag FlagValue) error
// func (v *Viper) BindFlagValues(flags FlagValueSet) (err error)
// func (v *Viper) BindPFlag(key string, flag *pflag.Flag) error
// func (v *Viper) BindPFlags(flags *pflag.FlagSet) error
// func (v *Viper) ClearConfigSettings()
// func (v *Viper) ClearPFlag(key string)
// func (v *Viper) ConfigFileUsed() string
// func (v *Viper) Debug()
// func (v *Viper) Desc(key string) (string, UseLevel, int)
// func (v *Viper) Get(key string) interface{}
// func (v *Viper) GetBool(key string) bool
// func (v *Viper) GetDuration(key string) time.Duration
// func (v *Viper) GetFloat64(key string) float64
// func (v *Viper) GetInt(key string) int
// func (v *Viper) GetInt64(key string) int64
// func (v *Viper) GetSizeInBytes(key string) uint
// func (v *Viper) GetString(key string) string
// func (v *Viper) GetStringMap(key string) map[string]interface{}
// func (v *Viper) GetStringMapString(key string) map[string]string
// func (v *Viper) GetStringMapStringSlice(key string) map[string][]string
// func (v *Viper) GetStringSlice(key string) []string
// func (v *Viper) GetTime(key string) time.Time
// func (v *Viper) InConfig(key string) bool
// func (v *Viper) IsSet(key string) bool
// func (v *Viper) MergeConfig(in io.Reader) error
// func (v *Viper) MergeInConfig() error
// func (v *Viper) OnConfigChange(run func(in fsnotify.Event))
// func (v *Viper) ReadConfig(in io.Reader) error
// func (v *Viper) ReadInConfig() error
// func (v *Viper) ReadRemoteConfig() error
// func (v *Viper) RegisterAlias(alias string, key string)
// func (v *Viper) Set(key string, value interface{})
// func (v *Viper) SetConfigFile(in string)
// func (v *Viper) SetConfigName(in string)
// func (v *Viper) SetConfigType(in string)
// func (v *Viper) SetDefault(key string, value interface{})
// func (v *Viper) SetDesc(key string, desc string, useLevel UseLevel, useScope int)
// func (v *Viper) SetEnvKeyReplacer(r *strings.Replacer)
// func (v *Viper) SetEnvPrefix(in string)
// func (v *Viper) SetFs(fs afero.Fs)
// func (v *Viper) SetPFlags(flags *pflag.FlagSet) (err error)
// func (v *Viper) SetTypeByDefaultValue(enable bool)
// func (v *Viper) String() string
// func (v *Viper) Sub(key string) *Viper
// func (v *Viper) Unmarshal(rawVal interface{}) error
// func (v *Viper) UnmarshalExact(rawVal interface{}) error
// func (v *Viper) UnmarshalKey(key string, rawVal interface{}) error
// func (v *Viper) WatchConfig()
// func (v *Viper) WatchRemoteConfig() error