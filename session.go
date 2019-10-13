package session

import (
	"github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	"net/http"
)

var (
	sessions *session.Manager
)

func NewSession(config Config) error {
	var (
		err error
	)

	sessions, err = session.NewManager(config.Type, config.ManagerConfig)
	if err != nil {
		return err
	}

	go sessions.GC()

	return nil
}

func GetSession(w http.ResponseWriter, r *http.Request) (Store, error) {
	s, err := sessions.SessionStart(w, r)
	if err != nil {
		return nil, err
	}

	defer s.SessionRelease(w)

	return s, nil
}
