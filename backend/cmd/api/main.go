package main

import (
	"flag"
	"log"
	"os"

	server "github.com/voyagegroup/treasure-2020-b"
	"github.com/voyagegroup/treasure-2020-b/zoom"
)

func main() {
	var databaseDatasource string
	var serviceAccountKeyPath string
	var port int
	flag.StringVar(&databaseDatasource, "databaseDatasource", "root:password@tcp(localhost:3306)/treasure_app", "Should looks like root:password@tcp(hostname:port)/dbname")
	flag.StringVar(&serviceAccountKeyPath, "serviceAccountKeyPath", "", "Path to service account key")
	flag.IntVar(&port, "port", 1991, "Web server port")
	flag.Parse()

	zoomClientID := os.Getenv("ZOOM_CLIENT_ID")
	zoomClientSecret := os.Getenv("ZOOM_CLIENT_SECRET")
	zoomRedirectURI := os.Getenv("ZOOM_REDIRECT_URI")
	zoomFrontendRedirectURI := os.Getenv("FRONTEND_BASE") + "/dashboard"
	zoomAuthClient := zoom.NewZoomAuthClient(
		zoomClientID,
		zoomClientSecret,
		zoomRedirectURI,
		zoomFrontendRedirectURI,
	)

	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	s := server.NewServer()
	if err := s.Init(databaseDatasource, serviceAccountKeyPath, zoomAuthClient); err != nil {
		log.Fatal(err)
	}
	s.Run(port)
}
