## gocheck 详解

>Golang下如何写单元测试？  
官方提供的testing Package略显简陋，不过好在我们有Gocheck。

[gocheck官网](http://labix.org/gocheck)

**示例-1**

```
package hello_test

import (
    "testing"
    "io/ioutil"
    "io"

    . "gopkg.in/check.v1"
)

const txt = "adfagaggafaf"

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
    dir string   // 测试用的临时目录
    f   string   // 测试用的临时文件
}

var _ = Suite(&MySuite{})

// Setupsuite 准备测试用的临时文件
func (s *MySuite) SetUpSuite(c *C) {
    dir := c.MkDir()    // Suite结束后会自动销毁c.MkDir()创建的目录

    tmpfile, err := ioutil.TempFile(dir, "")
    if err != nil {
        c.Errorf("Fail to create test file: %v\n", tmpfile.Name(), err)
    }

    err = WriteFile(tmpfile.Name(), txt)
    if err != nil {
        c.Errorf("Fail to prepare test file.%v\n", tmpfile.Name(), err)
    }
    s.dir = dir
    s.f = tmpfile.Name()   
}

func (s *MySuite) TestFoo(c *C) {
    // ... 实际测试代码
    c.Assert(bkpName, Matches, s.f+".ops_agent_bkp.+")
}
```


>“独立性”的原则告诉我们，对于需要调用外部api的功能，最好mock数据。利用gocheck的SetUpSuite()和TearDownSuite()方法，可以新建一个http test server，结束时关闭它。

**示例-2**

```$xslt
package hello_test

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"

    . "gopkg.in/check.v1"
)

const (
    resp1 = `{
   "data" : {
      "cluster" : "*****",
      "hostname" : "xxxxx"
   },
   "err_code" : 0,
   "err_msg" : ""
}
`
    resp2 = `{
   "data" : [
      {
         "hostname" : "e18h13551.XXX",
         "ip" : "100.22.33.44",
         "state" : "GOOD"
      },
      {
         "hostname" : "dddd",
         "ip" : "101.14.12.55",
         "state" : "GOOD"
      }
   ],
   "err_code" : 0,
   "err_msg" : ""
}
`
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
    ts *httptest.Server
}

func (s *MySuite) SetUpSuite(c *C) {
    h := http.NewServeMux()
    h.HandleFunc("/machine", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, resp1)
    })
    h.HandleFunc("/batch", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, resp2)
    })
    s.ts = httptest.NewServer(h)
}

func (s *MySuite) TearDownSuite(c *C) {
    s.ts.Close()
}

var _ = Suite(&MySuite{})

func (s *MySuite) TestFoo(c *C) {
    // 实际测试代码....
    clusterName, err := getClusterName(s.ts.URL, "/machine")
    c.Assert(err, IsNil)
    c.Assert(clusterName, Equals, "MiniLVSCluster-5e87-2384205713506559")
}
```