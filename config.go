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
	Type string `yaml:"Type"`

	*ManagerConfig
}
