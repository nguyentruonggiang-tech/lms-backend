package swagger

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/swaggest/openapi-go/openapi3"
)

func Start(ginEngine *gin.Engine) {
	ginEngine.GET("/docs", func(ctx *gin.Context) {
		html := `<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta name="color-scheme" content="light dark" />
  <title>LMS Backend API</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.32.2/swagger-ui.css" />
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5.32.2/swagger-ui-bundle.js"></script>
  <script src="https://unpkg.com/swagger-ui-dist@5.32.2/swagger-ui-standalone-preset.js"></script>
  <script>
    window.onload = function () {
      window.ui = SwaggerUIBundle({
        url: "/docs.json",
        dom_id: "#swagger-ui",
        deepLinking: true,
        presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
        layout: "StandaloneLayout"
      });
    };
  </script>
</body>
</html>`
		ctx.Data(200, "text/html; charset=utf-8", []byte(html))
	})

	ginEngine.GET("/docs.json", func(ctx *gin.Context) {
		reflector := &openapi3.Reflector{}
		reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}
		reflector.Spec.Info.
			WithTitle("LMS Backend API").
			WithVersion("1.0.0").
			WithDescription("Mini LMS Platform — Backend API Documentation")

		serverDesc := "Local development"
		reflector.Spec.Servers = []openapi3.Server{
			{URL: "http://localhost:8080", Description: &serverDesc},
		}

		modules := []func(*openapi3.Reflector) error{
			auth,
			adminCategory,
			adminCourse,
			adminSection,
			adminLesson,
			adminQuiz,
			adminQuestion,
		}

		for _, mod := range modules {
			if err := mod(reflector); err != nil {
				ctx.JSON(500, gin.H{"message": "swagger build error"})
				return
			}
		}

		docJSON, err := json.MarshalIndent(reflector.Spec, "", "  ")
		if err != nil {
			ctx.JSON(500, gin.H{"message": "swagger marshal error"})
			return
		}

		ctx.Data(200, "application/json; charset=utf-8", docJSON)
	})
}
