package request

import (
    "testing"
    . "github.com/onsi/gomega"
    . "github.com/franela/goblin"
    "net/http/httptest"
    "net/http"
    "fmt"
)

func TestRequest(t *testing.T) {
    g := Goblin(t)

    RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

    g.Describe("Request", func() {

        g.Describe("General request methods", func() {
            var ts *httptest.Server

            g.Before(func() {
                ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                    if r.Method == "GET" && r.URL.Path == "/foo" {
                        w.WriteHeader(200)
                        fmt.Fprint(w, "bar")
                    }
                }))
            })

            g.After(func() {
                ts.Close()
            })

            g.It("Should do a GET", func() {
                res, ok, _ := Get{ Uri: ts.URL + "/foo" }.Do()

                Expect(ok).Should(BeTrue())
                Expect(res.StatusCode).Should(Equal(200))
                Expect(res.Body).Should(Equal("bar"))
            })

            g.It("Should do a POST")
            g.It("Should do a PUT")
            g.It("Should do a DELETE")
            g.It("Should do a OPTIONS")
            g.It("Should do a PATCH")
            g.It("Should do a TRACE")
            g.It("Should do a custom method")
        })

        g.Describe("Timeouts", func() {
            g.It("Should timeout after a specified amount of ms")
            g.It("Should connect timeout after a specified amount of ms")
        })

        g.Describe("Misc", func() {
            g.It("Should offer to set request headers")
        })
    })
}