package belajar_golang_viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	config := viper.New()

	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		t.Log(err.Error())
	}

	assert.Nil(t, err)
}

func TestJSON2(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("../")

	err := config.ReadInConfig()
	if err != nil {
		t.Log(err.Error())
	}

	appName := config.GetString("app.name")
	appVersion := config.GetString("app.version")
	appAuthor := config.GetString("app.Author")

	databaseShowSql := config.GetBool("database.show_sql")
	databaseHost := config.GetString("database.host")
	databasePort := config.GetInt("database.port")

	//app
	assert.Equal(t, "belajar_golang_viper", appName)
	assert.Equal(t, "1.0.0", appVersion)
	assert.Equal(t, "adib hauzan sofyan", appAuthor)
	//end app

	//database
	assert.Equal(t, true, databaseShowSql)
	assert.Equal(t, "localhost", databaseHost)
	assert.Equal(t, 3306, databasePort)
	//end database

}

func TestYAML(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		t.Log(err.Error())
	}

	appName := config.GetString("app.name")
	appVersion := config.GetString("app.version")
	appAuthor := config.GetString("app.Author")

	databaseShowSql := config.GetBool("database.show_sql")
	databaseHost := config.GetString("database.host")
	databasePort := config.GetInt("database.port")

	//app
	assert.Equal(t, "belajar_golang_viper", appName)
	assert.Equal(t, "1.0.0", appVersion)
	assert.Equal(t, "adib hauzan sofyan", appAuthor)
	//end app

	//database
	assert.Equal(t, true, databaseShowSql)
	assert.Equal(t, "localhost", databaseHost)
	assert.Equal(t, 3306, databasePort)
	//end database

}
