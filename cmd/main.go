package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dgrijalva/jwt-go"
	"github.com/fasthttp/websocket"
	"github.com/joho/godotenv"
	"github.com/trinhdaiphuc/social-network/config"
	"github.com/trinhdaiphuc/social-network/graph"
	"github.com/trinhdaiphuc/social-network/graph/generated"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultPort  = "8080"
	databaseName = "social-network"
)

var upgrader = websocket.FastHTTPUpgrader{}

func DatabaseConnection(ctx context.Context) (*mongo.Database, error) {
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().MongoURI))
	if err != nil {
		return nil, err
	}
	db := client.Database(databaseName)
	return db, nil
}

func InitGraphQL(db *mongo.Database) (*handler.Server, http.HandlerFunc) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: graph.NewResolver(db),
		},
	))

	playground := playground.Handler("GraphQL playground", "/query")
	return srv, playground
}

func AllowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			preflightHandler(w, r)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				w.Write([]byte(fmt.Sprintf("Preflight request for %s", r.URL.Path)))
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Content-Length", "Accept", "Authorization",
		"X-Auth-Token", "Origin", "Refresh-Token"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	exposeHeaders := []string{"Content-Type", "Content-Length", "Accept", "Authorization", "X-Auth-Token", "Origin"}
	w.Header().Set("Access-Control-Expose-Headers", strings.Join(exposeHeaders, ","))
}

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) == 2 {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.GetConfig().JwtKey), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "user", claims)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	godotenv.Load()
	config.Load()
	port := config.GetConfig().Port
	if port == "" {
		port = defaultPort
	}

	mongoCtx := context.Background()
	db, err := DatabaseConnection(mongoCtx)
	if err != nil {
		panic(err)
	}

	gqlHandler, playground := InitGraphQL(db)

	http.Handle("/", playground)
	http.Handle("/query", AllowCORS(jwtMiddleware(gqlHandler)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// Listen from a different goroutine
	go func() {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal)      // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt) // When an interrupt is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	fmt.Println("Close DB connection")
	db.Client().Disconnect(mongoCtx)

	fmt.Println("Running cleanup tasks...")
}
