package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kardianos/service"
)

const (
	appVersionStr = "v0.1"
	nameOfService = "sqm-certcheck"
)

var (
	routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			defaultHandler,
		},
	}
	router *mux.Router
)

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	svcConfig := &service.Config{
		Name:        nameOfService,
		DisplayName: nameOfService,
		Description: nameOfService,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	//urlUatWeb := "https://brimuat.ajg.com/"
	/*fmsUrls := []string{"https://brimuat.ajg.com/", "https://brim.ajg.com/", "https://fmsuat.ajg.com/", "https://fms.ajg.com"}
	results := []struct {
		Url    string
		Expire time.Time
		Error  error
	}{}

	for _, url := range fmsUrls {
		t, err := checkUrl(url)
		if err != nil {
			fmt.Printf("url: %s,expire: %s,error: %s\n", url, t.String(), err.Error())
		} else {
			fmt.Printf("url: %s,expire: %s\n", url, t.String())
		}
		results = append(results, struct {
			Url    string
			Expire time.Time
			Error  error
		}{
			Url:    url,
			Expire: t,
			Error:  err,
		})
	}

	fmt.Printf("%v\n", results)

	return*/

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body>We are up and running "+nameOfService+" version "+appVersionStr+" ;)</body></html>")
}
