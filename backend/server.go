package server

import (
	"fmt"

	"log"
	"net/http"
	"os"

	"firebase.google.com/go/auth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/rs/cors"

	"github.com/voyagegroup/treasure-2020-b/controller"
	"github.com/voyagegroup/treasure-2020-b/db"
	"github.com/voyagegroup/treasure-2020-b/firebase"
	"github.com/voyagegroup/treasure-2020-b/middleware"
	"github.com/voyagegroup/treasure-2020-b/zoom"
)

type Server struct {
	db             *sqlx.DB
	router         *mux.Router
	authClient     *auth.Client
	zoomAuthClient *zoom.ZoomAuthClient
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string, serviceAccountKeyPath string, zoomAuthClient *zoom.ZoomAuthClient) error {
	authClient, err := firebase.InitAuthClient(serviceAccountKeyPath)
	if err != nil {
		return fmt.Errorf("failed init auth client. %s", err)
	}
	s.authClient = authClient

	s.zoomAuthClient = zoomAuthClient
	log.Println(s.zoomAuthClient)

	cs := db.NewDB(datasource)
	dbcon, err := cs.Open()
	if err != nil {
		return fmt.Errorf("failed db init. %s", err)
	}
	s.db = dbcon
	s.router = s.Route()
	return nil
}

func (s *Server) Run(port int) {
	log.Printf("Listening on port %d", port)
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuth(s.authClient, s.db)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	authChain := commonChain.Append(
		authMiddleware.Handler,
	)

	r := mux.NewRouter()
	roomController := controller.NewRoom(s.db)
	ownerController := controller.NewOwner(s.db)
	zoomController := controller.NewZoom(s.db, s.zoomAuthClient)
	lessonController := controller.NewLesson(s.db, s.zoomAuthClient)
	reservationController := controller.NewReservation(s.db)
	webController := controller.NewWeb(s.db)
	r.Methods(http.MethodGet).Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	r.Methods(http.MethodGet).Path("/room").Handler(authChain.Then(AppHandler{roomController.Index}))
	r.Methods(http.MethodPost).Path("/room").Handler(authChain.Then(AppHandler{roomController.Create}))

	r.Methods(http.MethodPost).Path("/owner").Handler(authChain.Then(AppHandler{ownerController.Create}))
	r.Methods(http.MethodGet).Path("/owner/me").Handler(authChain.Then(AppHandler{ownerController.GetMe}))
	r.Methods(http.MethodGet).Path("/lesson").Handler(authChain.Then(AppHandler{lessonController.GetAll}))
	r.Methods(http.MethodGet).Path("/lesson/{lesson_id:[0-9]+}").Handler(commonChain.Then(AppHandler{lessonController.GetByID}))
	r.Methods(http.MethodPost).Path("/lesson").Handler(authChain.Then(commonChain.Then(AppHandler{lessonController.Create})))
	r.Methods(http.MethodGet).Path("/owner/{owner_id:[0-9]+}/lesson").Handler(commonChain.Then(AppHandler{lessonController.GetAllByOwnerID}))
	r.Methods(http.MethodPost).Path("/lesson/{lesson_id:[0-9]+}/reservation").Handler(commonChain.Then(AppHandler{reservationController.Create}))
	r.Methods(http.MethodGet).Path("/owner/{owner_id:[0-9]+}/web").Handler(commonChain.Then(AppHandler{webController.GetAllByOwnerID}))
	r.Methods(http.MethodPost).Path("/web").Handler(authChain.Then(commonChain.Then(AppHandler{webController.Create})))
	r.Methods(http.MethodGet).Path("/web/{web_id:[0-9]+}").Handler(commonChain.Then(AppHandler{webController.Get}))

	r.Methods(http.MethodGet).Path("/owner/{firebase_uid}/zoom_auth").Handler(AppHandler{zoomController.Create})

	r.PathPrefix("").Handler(commonChain.Then(http.StripPrefix("/img", http.FileServer(http.Dir("./img")))))
	return r
}
