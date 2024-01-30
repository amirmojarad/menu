package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfig(t *testing.T) {
	appCfg, err := GetConfig()

	assert.NotEmpty(t, appCfg.Database.Name)
	assert.NotEmpty(t, appCfg.Database.Port)
	assert.NotEmpty(t, appCfg.Database.User)
	assert.NotEmpty(t, appCfg.Database.Host)
	assert.NotEmpty(t, appCfg.Database.Password)
	assert.NotEmpty(t, appCfg.Database.MigrationPath)
	assert.NotEmpty(t, appCfg.App.Debug)
	assert.NotEmpty(t, appCfg.App.Port)
	assert.NotEmpty(t, appCfg.App.Host)
	assert.NotEmpty(t, appCfg.App.Name)
	assert.Nil(t, err)

}
