package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	stdlog "log"

	"github.com/jonayrodriguez/translation-service/internal/translation/config"
	"github.com/jonayrodriguez/translation-service/internal/translation/controller"
	"github.com/jonayrodriguez/translation-service/internal/translation/service"

	"github.com/jonayrodriguez/translation-service/internal/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// configuration defaults
const (
	DefaultPort        = 50051
	DefaultServiceName = "Translation Service"
	DefaultDBHost      = "127.0.0.1"
	DefaultDBPort      = 50059
	DefaultDBSchema    = "translation"
	DefaultDBUserName  = "test"
	DefaultDBPassword  = ""
)

var (
	conf config.Config
	// version is set at compile time
	version = "debug"
	rootCmd = &cobra.Command{
		Use: "configuration",
		Run: func(cmd *cobra.Command, args []string) {
			runServer()
		},
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		stdlog.Fatalf("Failed to retrieve the args: %v", err)
	}
}

// runServer to start up the service
func runServer() {
	logger, _ := log.Config(conf.Logging)
	logger.Info(fmt.Sprintf("Starting %s. Version: %s", conf.Service.Name, version))

	address := fmt.Sprintf("0.0.0.0:%d", conf.Server.Port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to listen: %v", err))
	}
	s := grpc.NewServer()
	err = registerController(s)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to register the controller: %v", err))
	}
	reflection.Register(s)

	// Graceful Shutdown
	gracefulShutDown(s)

	logger.Info(fmt.Sprintf("Listening on %s", address))
	if err := s.Serve(lis); err != nil {
		logger.Fatal(fmt.Sprintf("Failed to serve %v", err))
		_ = log.Logger.Sync()
	}
}

// registerController will register the controller in the server
func registerController(server grpc.ServiceRegistrar) error {
	// TODO - Replace nil with the repository
	svc, err := service.NewTranslateService(nil)

	if err != nil {
		return err
	}
	controller.
		NewTranslationService(svc).
		Register(server)
	return nil
}

// Basic Channel to handle SIGINT and SIGTERM for a graceful shutdown
func gracefulShutDown(s *grpc.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer signal.Stop(ch)
		sig := <-ch
		log.Logger.Info(fmt.Sprintf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done"))
		// Flush logger buffer before stopping
		_ = log.Logger.Sync()

		s.GracefulStop()
	}()
}

// init method to configure cobra to use viper.
//nolint
func init() {
	cobra.OnInitialize(initViperConfig)

	// Get service config from flag, or use default
	rootCmd.Flags().StringVarP(&conf.Service.Name, "service", "s", DefaultServiceName, "service name")
	rootCmd.Flags().IntVarP(&conf.Server.Port, "port", "p", DefaultPort, "gRPC server port")

	// Get email server config from flag, or use default
	rootCmd.Flags().StringVar(&conf.DB.Host, "db_host", DefaultDBHost, "database host")
	rootCmd.Flags().IntVar(&conf.DB.Port, "db_port", DefaultDBPort, "databaser port")
	rootCmd.Flags().StringVar(&conf.DB.Schema, "db_schema", DefaultDBSchema, "database schema")
	rootCmd.Flags().StringVar(&conf.DB.Username, "db_username", DefaultDBUserName, "username for database")
	rootCmd.Flags().StringVar(&conf.DB.Password, "db_password", DefaultDBPassword, "database password")

	// Get logging configuration
	rootCmd.Flags().StringVarP(&conf.Logging.Level, "loglevel", "l", "info", "log level")
	rootCmd.Flags().BoolVarP(&conf.Logging.Development, "logdev", "d", false, "development logging")

}

// initConfig reads in config file and ENV variables if set.
func initViperConfig() {
	v := viper.New()
	v.AutomaticEnv()      // read in environment variables that match
	bindFlags(rootCmd, v) // Bind the current command's flags to viper
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STRING_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			_ = v.BindEnv(f.Name, envVarSuffix)
		}
		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
