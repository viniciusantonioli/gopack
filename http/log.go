package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func stdout() io.Writer {
	return os.Stdout
}

func logFile(fileName string) (io.Writer, error) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func multiWriter(logFile, stdout io.Writer) (io.Writer, error) {
	multi := io.MultiWriter(logFile, os.Stdout)
	return multi, nil
}

func Logger(logPath, env string) {
	switch env {
	case "development":
		fmt.Println("Ambiente: Desenvolvimento")
		stdout := stdout()
		multiWriter, err := multiWriter(stdout, nil)
		if err != nil {
			log.Fatal("MultiWriter:", err)
		}
		log.SetOutput(multiWriter)
	case "staging":
		fmt.Println("Ambiente: Homologação")
		logFile, logfErr := logFile(logPath)
		if logfErr != nil {
			log.Fatal("Logfile:", logfErr)
		}
		stdout := stdout()
		multiWriter, mwErr := multiWriter(logFile, stdout)
		if mwErr != nil {
			log.Fatal("MultiWriter:", mwErr)
		}
		log.SetOutput(multiWriter)
	case "production":
		fmt.Println("Ambiente: Produção")
		logFile, logfErr := logFile(logPath)
		if logfErr != nil {
			log.Fatal("Logfile:", logfErr)
		}
		stdout := stdout()
		multiWriter, mwErr := multiWriter(logFile, stdout)
		if mwErr != nil {
			log.Fatal("MultiWriter:", mwErr)
		}
		log.SetOutput(multiWriter)
	default:
		log.Fatal("Ambiente não definido.")
	}
	if env == "development" {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		return
	}
	log.SetFlags(log.Ldate | log.Ltime)
}

func LogHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		if string(body) != "" {
			log.Printf("%s %s %s\nRequisição: %s\n", r.RemoteAddr, r.Method, r.URL, body)
		} else {
			log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		}
		handler.ServeHTTP(w, r)
	})
}
