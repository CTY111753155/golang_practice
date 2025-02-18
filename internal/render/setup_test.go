package render

import (
	"encoding/gob"
	"myapp/internal/config"
	"myapp/internal/models"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	//what am I going to put in the session
	gob.Register(models.Reservation{})
	//change this to true when in production
	testApp.Inproduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (mw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (mw *myWriter) WriteHeader(i int) {

}

func (mw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
