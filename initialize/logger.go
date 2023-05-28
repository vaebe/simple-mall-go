package initialize

import (
	"encoding/json"
	"go.uber.org/zap"
	"simple-mall/global"
)

func InitLogger() {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "./tmp/logs"],
	  "errorOutputPaths": ["stderr", "./tmp/error"],
	  "initialFields": {"version": "v1"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	global.Logger = zap.Must(cfg.Build())
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(global.Logger)

	zap.ReplaceGlobals(global.Logger)
}
