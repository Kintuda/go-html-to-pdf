package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/kintuda/go-html-to-pdf/pkg/config"

	// "github.com/kintuda/golang-template-fx-fiber-zero/pkg/notification"
	"github.com/rs/zerolog/log"
)

type Router struct {
	engine              *fiber.App
	cfg                 *config.AppConfig
	notificationHandler *NotificationHandler
}

func NewRouter(
	cfg *config.AppConfig,
	notificationProvider notification.NotificationProvider,
) *Router {
	app := fiber.New(fiber.Config{
		AppName:      "notification-server",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	})

	app.Use(requestid.New())
	app.Use(idempotency.New(idempotency.Config{
		Lifetime:  60 * time.Minute,
		KeyHeader: "x-idempotency-key",
		Next: func(c *fiber.Ctx) bool {
			return fiber.IsMethodSafe(c.Method())
		},
	}))

	return &Router{
		engine: app,
		cfg:    cfg,
	}
}

func (r *Router) Start() error {
	if err := r.engine.Listen(r.cfg.HttpPort); err != nil {
		log.Error().Err(err).Msg("error while starting fiber server")
		return err
	}

	return nil
}

func (r *Router) Stop() error {
	if err := r.engine.ShutdownWithTimeout(2 * time.Minute); err != nil {
		log.Error().Err(err).Msg("error while stoping fiber server")
		return err
	}

	return nil
}

func (r *Router) RegisterRoutes() {
	v1 := r.engine.Group("/v1")

	notifications := v1.Group("notifications")
	notifications.Post("/sms", r.notificationHandler.SendSMS)
	notifications.Post("/email", r.notificationHandler.SendEmail)
}
