package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas rotas da api
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		//router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost) similar
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	return r
}
