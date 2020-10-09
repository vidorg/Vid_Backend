package logger

import (
	"github.com/Aoi-hosizora/ahlib-more/xlogrus"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/sirupsen/logrus"
	"github.com/vidorg/vid_backend/src/config"
	"github.com/vidorg/vid_backend/src/provide/sn"
	"time"
)

func Setup() (*logrus.Logger, error) {
	c := xdi.GetByNameForce(sn.SConfig).(*config.Config)

	logger := logrus.New()
	logLevel := logrus.WarnLevel
	if c.Meta.RunMode == "debug" {
		logLevel = logrus.DebugLevel
	}

	logger.SetLevel(logLevel)
	logger.SetReportCaller(false)
	logger.SetFormatter(&xlogrus.CustomFormatter{
		ForceColor: true,
	})

	if c.Meta.LogRotate {
		rotateHook := xlogrus.NewRotateLogHook(&xlogrus.RotateLogConfig{
			MaxAge:       15 * 24 * time.Hour,
			RotationTime: 24 * time.Hour,
			Filepath:     c.Meta.LogPath,
			Filename:     c.Meta.LogName,
			Level:        logLevel,
			Formatter:    &logrus.JSONFormatter{TimestampFormat: time.RFC3339},
		})
		logger.AddHook(rotateHook)
	}
	if c.Meta.LogMq {
		mqHook, err := NewMQLogHook(&MQLogHookConfig{
			Formatter: &logrus.JSONFormatter{TimestampFormat: time.RFC3339},
		})
		if err != nil {
			return nil, err
		}
		logger.AddHook(mqHook)
	}

	return logger, nil
}
