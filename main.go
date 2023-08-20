package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const projectName = "Go Project Template" // REPLACE WITH YOUR PROJECT NAME HERE

func usage() {
	fmt.Fprintf(os.Stderr, `
<Project description>

	<binary name> [flags]

<Project details/usage>

%s accepts the following flags:

`[1:], projectName)
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, `

For more information, see https://github.com/<user>/<repo>.
`[1:])
}

func main() {
	os.Exit(mainRetCode())
}

func mainRetCode() int {
	var (
		debugLogs    bool
		logPath      string
		printVersion bool
	)

	flag.Usage = usage
	flag.BoolVar(&debugLogs, "debug", false, "enable debug logging")
	flag.StringVar(&logPath, "l", "stdout", "path to log to")
	flag.BoolVar(&printVersion, "version", false, "print version and build information and exit")
	flag.Parse()

	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Fprintln(os.Stderr, "build information not found")
		return 1
	}

	if printVersion {
		printVersionInfo(info)
		return 0
	}

	// build logger
	logCfg := zap.NewProductionConfig()
	logCfg.OutputPaths = []string{logPath}
	if debugLogs {
		logCfg.Level.SetLevel(zap.DebugLevel)
	}
	logCfg.EncoderConfig.TimeKey = "time"
	logCfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	logCfg.DisableCaller = true

	logger, err := logCfg.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating logger: %v", err)
		return 1
	}

	// log current version/commit
	versionFields := []zap.Field{
		zap.String("version", version),
	}
	for _, buildSetting := range info.Settings {
		if buildSetting.Key == "vcs.revision" {
			versionFields = append(versionFields, zap.String("commit", buildSetting.Value))
			break
		}
	}
	logger.Info("starting "+projectName, versionFields...)

	// may also want to add golang.org/x/sys/unix.SIGTERM on unix based OSes
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := mainErr(ctx, logger); err != nil {
		var exitCode *errJustExit
		if errors.As(err, &exitCode) {
			return int(*exitCode)
		}
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}

type errJustExit int

func (e errJustExit) Error() string { return fmt.Sprintf("exit: %d", e) }

// this disables a linter warning because nil is unconditionally returned
// here, remove this when adding your own code that can return errors
//
//nolint:unparam
func mainErr(ctx context.Context, logger *zap.Logger) error {
	// START MAIN LOGIC HERE

	<-ctx.Done()
	logger.Info("shutting down")

	// STOP MAIN LOGIC HERE

	return nil
}
