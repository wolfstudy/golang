## gocheck 详解

>Golang下如何写单元测试？  
官方提供的testing Package略显简陋，不过好在我们有Gocheck。

[gocheck官网](http://labix.org/gocheck)

### 测试原则

**3A原则：**

* Arrange: 测试用例执行之前需要准备测试数据，包括需要输入的数据及存量数据
* Act: 通过不同的参数来调用接口，并拿到返回
* Assert: 必须做断言，否则用例就没有任何意义了

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

## 如何获取gocheck

Install gocheck's check package with the following command:

```$xslt
go get gopkg.in/check.v1

```

To ensure you're using the latest version, run the following instead:

```$xslt
go get -u gopkg.in/check.v1

```

##小心得

如果你想把 gocheck这个项目内嵌到自己的项目中，OK，你可以把 gocheck fork到自己的github下，然后，修改其相应的路径。即可

## Assert方法

```$xslt
func (c *C) Assert(obtained interface{}, checker Checker, args ...interface{})
```

这个方法来判断,内置的 Check 接口实现有:

* DeepEquals
* Equals
* ErrorMatches
* FitsTypeOf
* HasLen
* Implements
* IsNil
* Matches
* NotNil
* PanicMatches

### 使用例子

```$xslt
c.Assert(value, DeepEquals, 42)
c.Assert(array, DeepEquals, []string{"hi", "there"})
c.Assert(value, Equals, 42)
c.Assert(err, ErrorMatches, "perm.*denied")
c.Assert(value, FitsTypeOf, int64(0))
c.Assert(value, FitsTypeOf, os.Error(nil))
c.Assert(list, HasLen, 5)
var e os.Error
c.Assert(err, Implements, &e)
c.Assert(err, IsNil)
c.Assert(err, Matches, "perm.*denied")
c.Assert(iface, NotNil)
c.Assert(func() { f(1, 2) }, PanicMatches, `open.*: no such file or directory`).
```

###另一个接口

```$xslt
func Not(checker Checker) Checker
```

### 用法

```$xslt
c.Assert(a, Not(Equals), b)
```
