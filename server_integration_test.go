import (
	"io"
	"net/http"
	"os"
	"bytes"
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"encoding/json"
)

func TestGetAllUser(t *testing.T) {
	var c User
	body := bytes.NewBufferString(`{
		"name":"GP",
		"age":19
	}`)
	err := request(http.MethodPost, uri("users"), body).Decode(&c)
	if err != nil {
		t.Fatal("can't create users", err)
	}

	var us []User
	res := request(http.MethodGet, uri("users"), nil)
	err = res.Decode(&us)

	assert.Nil(t, err)
	assert.EqaulValues(t, http.StatusOk, res.StatusCode)
	assert.Greater(t, len(us), 0)
}

func uri(paths ...string) string {
	host := "http://localhost:2550"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

type Response struct {
	http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.body).Decode(v)
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}