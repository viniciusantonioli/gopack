package http

import (
	"fmt"
	"net/http"
)

func Listen(port int, serverName, version, env string, router http.Handler) error {
	fmt.Printf("%s v%s ouvindo na porta: %d\n", serverName, version, port)
	if env != "development" {
		fmt.Printf("Gravar logs ativado\n")
	}

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), logHandler(router))
	if err != nil {
		return fmt.Errorf("Problema ao iniciar o servidor: %s", err)
	}
	return nil
}

func Handle(pattern, method string, handle func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if method != r.Method {
			w.Write([]byte(fmt.Sprintf("Método não permitido: %s", r.Method)))
			return
		}

		handle(w, r)
	})
}
