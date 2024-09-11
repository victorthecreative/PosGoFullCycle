package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/victorthecreative/PosGoFullCycle/API/configs"
	_ "github.com/victorthecreative/PosGoFullCycle/API/docs"
	"github.com/victorthecreative/PosGoFullCycle/API/internal/entity"
	"github.com/victorthecreative/PosGoFullCycle/API/internal/infra/database"
	"github.com/victorthecreative/PosGoFullCycle/API/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Victor Coelho
// @contact.url    http://www.fullcycle.com.br
// @contact.email  atendimento@fullcycle.com.br

// @license.name   Full Cycle License
// @license.url    http://www.fullcycle.com.br

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Carrega as configurações
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// Configura o banco de dados SQLite
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrar a estrutura do banco de dados
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	// Inicializar os repositórios e handlers
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	// Configurar o roteador Chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Middleware para inserir valores de JWT e expiração no contexto
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn)) // Corrigido para "jwtExpiresIn"

	// Rotas de produtos (com autenticação JWT)
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	// Rotas de usuários
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/generate_token", userHandler.GetJWT)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	// Iniciar o servidor na porta 8000
	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
