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

type Config struct {
	Type string `json:"Type"`

	*ManagerConfig
}

// eg.
// {"Type": "memory", "cookieName": "gosessionsid", "gclifetime": 3600, "ProviderConfig": ""}
// {"Type": "file", "cookieName": "gosessionsid", "gclifetime": 3600, "ProviderConfig": "./tmp"}
// {"Type": "redis", "cookieName": "gosessionsid", "gclifetime": 3600, "ProviderConfig": "127.0.0.1:6379,100,sessions"}
// {"Type": "mysql", "cookieName": "gosessionsid", "gclifetime": 3600, "ProviderConfig": "username:password@protocol(address)/dbname?param=value"}
