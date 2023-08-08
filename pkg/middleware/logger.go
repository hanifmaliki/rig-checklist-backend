package middleware

import (
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// New creates a new middleware handler
func Logger() fiber.Handler {
	// Set PID once
	pid := os.Getpid()

	// Set variables
	var (
		once       sync.Once
		errHandler fiber.ErrorHandler
	)

	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		var start, stop time.Time
		// Set error handler once
		once.Do(func() {
			// override error handler
			errHandler = c.App().Config().ErrorHandler
		})

		// Set latency start time
		start = time.Now()

		// Handle request, store err for logging
		chainErr := c.Next()

		// Manually call error handler
		if chainErr != nil {
			if err := errHandler(c, chainErr); err != nil {
				_ = c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		// Set latency stop time
		stop = time.Now()

		// Get request id from previous middleware
		reqID := c.Locals("requestid").(string)

		if chainErr != nil {
			log.Error().
				Err(chainErr).
				Str("req_id", reqID).
				Int("pid", pid).
				Int("status_code", c.Response().StatusCode()).
				Str("method", c.Method()).
				Str("host", c.Hostname()).
				Str("path", c.Path()).
				Dur("latency", stop.Sub(start).Round(time.Millisecond)).
				Str("ua", c.Get(fiber.HeaderUserAgent)).
				Msg("error")
		} else {
			log.Info().
				Str("req_id", reqID).
				Int("pid", pid).
				Int("status_code", c.Response().StatusCode()).
				Str("method", c.Method()).
				Str("host", c.Hostname()).
				Str("path", c.Path()).
				Dur("latency", stop.Sub(start).Round(time.Millisecond)).
				Str("ua", c.Get(fiber.HeaderUserAgent)).
				Msg("ok")
		}

		// End chain
		return nil
	}
}
