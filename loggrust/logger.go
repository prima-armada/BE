package loggrust

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log = logrus.New()

func ConfigureLogger() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    50,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
		LocalTime:  true,
	})

	log.SetFormatter(&logrus.TextFormatter{})
	log.SetLevel(logrus.DebugLevel)
}

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		res := c.Response()

		log.WithFields(logrus.Fields{
			"method":  req.Method,
			"path":    req.URL.Path,
			"ip":      c.RealIP(),
			"headers": req.Header,
		}).Info("Request received")

		err := next(c)

		log.WithFields(logrus.Fields{
			"status":  res.Status,
			"latency": res.Header().Get("X-Response-Time"),
		}).Info("Request handled")

		// Periksa apakah terjadi kesalahan
		if err != nil {
			log.WithError(err).Fatal("Fatal error occurred")
		}

		return err
	}
}
