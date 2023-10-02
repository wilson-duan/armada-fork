package common

import (
	"crypto/tls"
	"fmt"
	"github.com/armadaproject/armada/internal/common/certs"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/weaveworks/promrus"

	"github.com/armadaproject/armada/internal/common/armadacontext"
	commonconfig "github.com/armadaproject/armada/internal/common/config"
	"github.com/armadaproject/armada/internal/common/logging"
)

const baseConfigFileName = "config"

// RFC3339Millis
const logTimestampFormat = "2006-01-02T15:04:05.999Z07:00"

func BindCommandlineArguments() {
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Error()
		os.Exit(-1)
	}
}

// TODO Move code relating to config out of common into a new package internal/serverconfig
func LoadConfig(config interface{}, defaultPath string, overrideConfigs []string) *viper.Viper {
	v := viper.NewWithOptions(viper.KeyDelimiter("::"))
	v.SetConfigName(baseConfigFileName)
	v.AddConfigPath(defaultPath)
	if err := v.ReadInConfig(); err != nil {
		log.Errorf("Error reading base config path=%s name=%s: %v", defaultPath, baseConfigFileName, err)
		os.Exit(-1)
	}
	log.Infof("Read base config from %s", v.ConfigFileUsed())

	for _, overrideConfig := range overrideConfigs {
		v.SetConfigFile(overrideConfig)
		err := v.MergeInConfig()
		if err != nil {
			log.Errorf("Error reading config from %s: %v", overrideConfig, err)
			os.Exit(-1)
		}
		log.Infof("Read config from %s", v.ConfigFileUsed())
	}

	v.SetEnvKeyReplacer(strings.NewReplacer("::", "_"))
	v.SetEnvPrefix("ARMADA")
	v.AutomaticEnv()

	err := v.Unmarshal(config, commonconfig.CustomHooks...)
	if err != nil {
		log.Error(err)
		os.Exit(-1)
	}

	return v
}

func UnmarshalKey(v *viper.Viper, key string, item interface{}) error {
	return v.UnmarshalKey(key, item, commonconfig.CustomHooks...)
}

// TODO Move logging-related code out of common into a new package internal/logging
func ConfigureCommandLineLogging() {
	commandLineFormatter := new(logging.CommandLineFormatter)
	log.SetFormatter(commandLineFormatter)
	log.SetOutput(os.Stdout)
}

func ConfigureLogging() {
	log.SetLevel(readEnvironmentLogLevel())
	log.SetFormatter(readEnvironmentLogFormat())
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
}

func readEnvironmentLogLevel() log.Level {
	level, ok := os.LookupEnv("LOG_LEVEL")
	if ok {
		logLevel, err := log.ParseLevel(level)
		if err == nil {
			return logLevel
		}
	}
	return log.InfoLevel
}

func readEnvironmentLogFormat() log.Formatter {
	formatStr, ok := os.LookupEnv("LOG_FORMAT")
	if !ok {
		formatStr = "colourful"
	}

	textFormatter := &log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: logTimestampFormat,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			return "", fileName
		},
	}

	switch strings.ToLower(formatStr) {
	case "json":
		return &log.JSONFormatter{TimestampFormat: logTimestampFormat}
	case "colourful":
		return textFormatter
	case "text":
		textFormatter.ForceColors = false
		textFormatter.DisableColors = true
		return textFormatter
	default:
		println(os.Stderr, fmt.Sprintf("Unknown log format %s, defaulting to colourful format", formatStr))
		return textFormatter
	}
}

func ServeMetrics(port uint16) (shutdown func()) {
	return ServeMetricsFor(port, prometheus.DefaultGatherer)
}

func ServeMetricsFor(port uint16, gatherer prometheus.Gatherer) (shutdown func()) {
	hook := promrus.MustNewPrometheusHook()
	log.AddHook(hook)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(gatherer, promhttp.HandlerOpts{}))
	return ServeHttp(port, mux)
}

// ServeHttp starts an HTTP server listening on the given port.
// TODO: Make block until a context passed in is cancelled.
func ServeHttp(port uint16, mux http.Handler) (shutdown func()) {
	return serveHttp(port, mux, false, "", "")
}

func ServeHttps(port uint16, mux http.Handler, certFile, keyFile string) (shutdown func()) {
	return serveHttp(port, mux, true, certFile, keyFile)
}

func serveHttp(port uint16, mux http.Handler, useTls bool, certFile, keyFile string) (shutdown func()) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	scheme := "http"
	if useTls {
		scheme = "https"
	}

	go func() {
		log.Printf("Starting %s server listening on %d", scheme, port)
		var err error = nil
		if useTls {
			certWatcher := certs.NewCachedCertificateService(certFile, keyFile, time.Minute)
			go func() {
				certWatcher.Run(armadacontext.Background())
			}()
			srv.TLSConfig = &tls.Config{
				GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
					return certWatcher.GetCertificate(), nil
				},
			}
			err = srv.ListenAndServeTLS("", "")
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err) // TODO Don't panic, return an error
		}

	}()
	// TODO There's no need for this function to panic, since the main goroutine will exit.
	// Instead, just log an error.
	return func() {
		ctx, cancel := armadacontext.WithTimeout(armadacontext.Background(), 5*time.Second)
		defer cancel()
		log.Printf("Stopping %s server listening on %d", scheme, port)
		e := srv.Shutdown(ctx)
		if e != nil {
			panic(e)
		}
	}
}
