package errors

import log "github.com/sirupsen/logrus"

func checkError(msg string, er error) {
	if er != nil {
		log.Error(msg, er.Error())
	}
}

func ErrorMiddleware(msg string, er error) {
	checkError("[MIDDLEWARE] "+msg, er)
}

func ErrorControler(msg string, er error) {
	checkError("[CONTROLLER] "+msg, er)
}

func ErrorConfig(msg string, er error) {
	checkError("[CONFIG] "+msg, er)
}

func ErrorRepository(msg string, er error) {
	checkError("[REPOSITORY] "+msg, er)
}
