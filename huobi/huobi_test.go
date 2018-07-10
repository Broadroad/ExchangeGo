package huobi
import(
	"testing"
	"time"
	"net/http"
)

func TestNewHuobi(t *testing.T) {
	hb := NewHuobi(&http.Client{}, "", "", "")
	hb.createWsConn()
	time.Sleep(2 * time.Second)
	defer hb.ws.CloseWs()
}