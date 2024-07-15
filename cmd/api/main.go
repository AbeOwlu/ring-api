package main

import (
	"fmt"
	"net/http"

	"github.com/AbeOwlu/ring-api/internal/handlers"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func main() {
	loggInit := zap.NewExample()
	defer loggInit.Sync()
	logger := loggInit.Sugar()

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	logger.Info("\n\n\ninit")
	logger.Info("\n\t Starting GO API Server...")
	fmt.Println(`
	 ________         _________
	/\   ____\       /\  _____ \
	\ \  \_____      \ \ \	  \ \
	 \ \   \_\ \      \ \ \____\ \
	  \ \_______\      \ \________\
	   \/_______/       \/________/`)

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		logger.Info(err)
	}

}
