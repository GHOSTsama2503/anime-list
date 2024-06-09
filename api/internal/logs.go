package internal

import (
	"anime-list/internal/env"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/natefinch/lumberjack.v2"
)

func logFile(filename string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    env.LogsMaxSize,
		MaxBackups: env.LogsMaxBackup,
		MaxAge:     env.LogsMaxAge,
		Compress:   env.LogsCompress,
		LocalTime:  env.LogsLocalTime,
	}
}

func SetupLogger() {
	log.SetTimeFormat(time.DateTime)

	if env.IsProduction {
		log.SetOutput(logFile(env.LogsAppPath))
	}
}

func LoggerMidleware(h http.Handler) http.Handler {

	style := log.DefaultStyles()
	for l := range style.Levels {
		delete(style.Levels, l)
	}

	logger := log.New(os.Stdout)
	logger.SetReportTimestamp(true)
	logger.SetTimeFormat(time.DateTime)
	logger.SetStyles(style)
	logger.SetPrefix("SERV")

	if env.IsProduction {
		logger.SetOutput(logFile(env.LogsAccessPath))
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		h.ServeHTTP(ww, r)

		took := time.Since(startTime)

		if env.IsDevelopment || env.PrettyLogs {
			logger.Info(fmt.Sprintf(
				"%s %s %s %s %s - %s %s %s %s - (%s)",
				ColorMagenta.Render(r.Method),
				ColorCyan.Render(r.URL.String()),
				ColorYellow.Render(r.Proto),
				ColorMagenta.Render("from"),
				ColorCyan.Render(r.RemoteAddr),
				ColorYellow.Render(fmt.Sprint(ww.Status())),
				ColorBlue.Render(fmt.Sprintf("%dB", ww.BytesWritten())),
				ColorMagenta.Render("in"),
				ColorGreen.Render(took.String()),
				r.UserAgent(),
			))
		} else {
			logger.Info("",
				"method", r.Method,
				"url", r.URL,
				"proto", r.Proto,
				"addr", r.RemoteAddr,
				"status", ww.Status(),
				"lenght", fmt.Sprint(ww.BytesWritten()),
				"took", took.String(),
				"ua", r.UserAgent(),
			)
		}
	}

	return http.HandlerFunc(fn)
}
