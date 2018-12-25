package session

import "net/http"

var (
	sessions *Manager
)

func NewSession(config Config) error {
	var (
		err error
	)

	sessions, err = NewManager(config.Type, config.ManagerConfig)
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
