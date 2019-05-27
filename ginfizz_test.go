/*
 * @Author: berryberry
 * @since: 2019-05-27 17:32:16
 * @lastTime: 2019-05-27 18:12:16
 * @LastAuthor: Do not edit
 */
package ginfizz

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupHTML(t *testing.T) *httptest.Server {
	InitFizz()
	router := Engine()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./testdata/*")
	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello.tmpl", map[string]string{"name": "world"})
	})
	return httptest.NewServer(router)
}

func TestRoute(t *testing.T) {
	server := setupHTML(t)
	res, err := http.Get(fmt.Sprintf("%s/test", server.URL))
	if err != nil {
		fmt.Println(err)
	}

	resp, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, "<h1>Hello world</h1>", string(resp))
}
