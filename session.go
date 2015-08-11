/**
 *
 *
 * @author Zhang Chaoren <zhangchaoren@openeasy.net>
 * @link http://openeasy.net/openeasy/session
 * @copyright Copyright &copy; 2014-2015 openeasy.net
 * @license http://www.apache.org/licenses/LICENSE-2.0.txt
 * @package sessions
 * @since 1.0
 *
 */

package session

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/astaxie/beego/session"
)

type Store session.SessionStore

var (
	logger *log.Logger
	m      *session.Manager
)

func init() {
	logger = log.New(os.Stdout, "[sessions] ", log.Ldate|log.Ltime)
}

func NewManager(config Config) error {
	var (
		j   []byte
		err error
	)

	j, _ = json.Marshal(config)

	m, err = session.NewManager(config.Type, string(j))
	if err == nil {
		go m.GC()
	}

	return err
}

func GetInst(w http.ResponseWriter, req *http.Request) (Store, error) {
	var (
		s   Store
		err error
	)
	s, err = m.SessionStart(w, req)
	if err == nil {
		defer s.SessionRelease(w)
	}

	return s, err
}
