package app

import (
	"github.com/VrMolodyakov/crypto-comparing/kraken/internal/config"
	"github.com/VrMolodyakov/crypto-comparing/kraken/pkg/logging"
)

type app struct {
	cfg    *config.Config
	Logger logging.Logger
}

func New() *app {
	return &app{}
}

func (a *app) InitLogger() {
	loggerCfg := logging.NewLogerConfig(
		a.cfg.Logger.DisableCaller,
		a.cfg.Logger.Development,
		a.cfg.Logger.DisableStacktrace,
		a.cfg.Logger.Encoding,
		a.cfg.Logger.Level,
	)
	a.Logger = logging.NewLogger(loggerCfg)
	a.Logger.InitLogger()
}

func (a *app) ReadConfig() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}
	a.cfg = cfg
	return nil
}
