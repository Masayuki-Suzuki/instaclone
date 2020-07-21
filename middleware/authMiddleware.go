package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	
	"github.com/Masayuki-Suzuki/instaclone/config"
	"github.com/Masayuki-Suzuki/instaclone/types"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Initialise Auth Func.
		auth, err := config.GetFirebaseAuthClient()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("error initializing auth function: %v\n", err)
		} else {
			// Get JWT from Client
			authHeader := r.Header.Get("Authorization")
			idToken := strings.Replace(authHeader, "Bearer ", "", 1)
			
			// Validate JWT
			_, err := auth.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				// If JWT is invalid, work error handling instead of Handler.
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Printf("error verifying ID Token: %v\n", err)
				
				errorMessage := types.ErrorMessage{ErrorMessage: "error verifying ID Token"}
				
				log.Println(json.NewEncoder(w).Encode(errorMessage))
				return
			}
			next.ServeHTTP(w, r)
		}
	}
}
