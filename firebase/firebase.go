package firebase

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	app *firebase.App
	err error
)

func Init() {
	app, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
}

//middlewear
func APIAuthWrap(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		if bearer == "" {
			c.String(http.StatusBadRequest, "Authorizationヘッダが設定されていません")
			c.Abort()
			return
		}
		idToken := strings.TrimPrefix(bearer, "Bearer ")
		if idToken == bearer {
			c.String(http.StatusBadRequest, "Authorization: Bearer ヘッダが設定されていません")
			c.Abort()
			return
		}

		ctx := context.Background()
		client, err := app.Auth(ctx)
		if err != nil {
			log.Print(err)
		}
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			log.Print(err)
		}
		log.Print(token)
		// if token.UID != "" {
		// 	c.String(http.StatusBadRequest, "管理者ではありません")
		// 	c.Abort()
		// 	return
		// }
		// db.Transaction(func(tx *gorm.DB) error {
		// 	if err := tx.Debug().Error; err != nil {
		// 		return err
		// 	}
		// 	return nil
		// })
		// if err != nil {
		// 	log.Print(err)
		// }
	}
}
