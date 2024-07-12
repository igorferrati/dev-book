package controllers

import "net/http"

// CarregarTelaDeLogin carrega tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de login"))
}
