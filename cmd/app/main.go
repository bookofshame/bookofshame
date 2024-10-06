package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bookofshame/bookofshame/handler"
	"github.com/bookofshame/bookofshame/internal/auth"
	"github.com/bookofshame/bookofshame/internal/gender"
	"github.com/bookofshame/bookofshame/internal/location"
	"github.com/bookofshame/bookofshame/internal/offender"
	"github.com/bookofshame/bookofshame/internal/user"
	"github.com/bookofshame/bookofshame/pkg/captcha"
	"github.com/bookofshame/bookofshame/pkg/config"
	"github.com/bookofshame/bookofshame/pkg/constants"
	"github.com/bookofshame/bookofshame/pkg/database"
	"github.com/bookofshame/bookofshame/pkg/email"
	"github.com/bookofshame/bookofshame/pkg/jwt"
	"github.com/bookofshame/bookofshame/pkg/locale"
	"github.com/bookofshame/bookofshame/pkg/logging"
	"github.com/bookofshame/bookofshame/pkg/sms"
	"github.com/bookofshame/bookofshame/pkg/storage"
	"github.com/invopop/ctxi18n"
)

func main() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)

	ctx := context.Background()
	cfg := config.New()

	logger := logging.NewLoggerFromEnv().With("version", cfg.Version)
	ctx = logging.WithLogger(ctx, logger)

	if err := ctxi18n.LoadWithDefault(locale.Content, constants.DefaultLanguage); err != nil {
		logger.Fatalf("error loading locales: %v", err)
	}

	db, err := database.New(ctx, cfg)
	if err != nil {
		logger.Fatalln("terminating due to db connection issue. error: %w", err)
	}

	j := jwt.New(ctx, cfg)

	// 3rd party services
	storageClient := storage.NewCloudflareR2(ctx, cfg)
	emailClient := email.NewEmailClient(ctx, cfg)
	smsClient := sms.NewSmsClient(ctx, cfg)
	recaptcha := captcha.NewReCaptcha(ctx, cfg)

	// Repos
	locationRepo := location.NewRepository(ctx, db)
	userRepo := user.NewRepository(ctx, db)
	offenderRepo := offender.NewRepository(ctx, db)
	genderRepo := gender.NewRepository(ctx, db)

	// Services
	locationService := location.NewService(cfg, locationRepo)
	authService := auth.NewService(cfg, userRepo)
	userService := user.NewService(cfg, userRepo, emailClient, smsClient)
	offenderService := offender.NewService(cfg, offenderRepo, storageClient)
	genderService := gender.NewService(cfg, genderRepo)

	mux := handler.SetupRoutes(
		j,
		recaptcha,
		genderService,
		locationService,
		authService,
		userService,
		offenderService,
	)

	addr := fmt.Sprintf(cfg.Ip + ":" + cfg.Port)
	server := &http.Server{Addr: addr, Handler: mux}

	go func() {
		logger.Infof("listening on :%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalln("server - server.ServeHttp: %w", err)
		}
	}()

	// Wait for interrupt signal
	<-exitSignal

	timeoutCtx, done := context.WithTimeout(ctx, 30*time.Second)
	defer done()

	if err := server.Shutdown(timeoutCtx); err != nil {
		logger.Errorf("HTTP server shutdown error: %s", err)
	}

	logger.Infoln("Http server shutdown complete")

	if err := db.Close(); err != nil {
		logger.Errorf("Failed to close database: %s", err)
	}
	logger.Infoln("Database connection closed")
}
