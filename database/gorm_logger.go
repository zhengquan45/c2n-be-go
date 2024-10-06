package database

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type LogrusLogger struct {
	Log *logrus.Logger
}

func (l *LogrusLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	switch level {
	case logger.Silent:
		newlogger.Log.SetLevel(logrus.WarnLevel)
	case logger.Error:
		newlogger.Log.SetLevel(logrus.ErrorLevel)
	case logger.Warn:
		newlogger.Log.SetLevel(logrus.WarnLevel)
	case logger.Info:
		newlogger.Log.SetLevel(logrus.InfoLevel)
	}
	return &newlogger
}

func (l *LogrusLogger) Info(ctx context.Context, s string, i ...interface{}) {
	l.Log.WithContext(ctx).Infof(s, i...)
}

func (l *LogrusLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	l.Log.WithContext(ctx).Warnf(s, i...)
}

func (l *LogrusLogger) Error(ctx context.Context, s string, i ...interface{}) {
	l.Log.WithContext(ctx).Errorf(s, i...)
}

func (l *LogrusLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		l.Log.WithContext(ctx).WithFields(logrus.Fields{
			"duration": elapsed,
			"rows":     rows,
			"error":    err,
		}).Error(sql)
	} else {
		l.Log.WithContext(ctx).WithFields(logrus.Fields{
			"duration": elapsed,
			"rows":     rows,
		}).Info(sql)
	}
}
