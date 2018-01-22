## gocheck 详解

>Golang下如何写单元测试？  
官方提供的testing Package略显简陋，不过好在我们有Gocheck。

[gocheck官网](http://labix.org/gocheck)

**示例**

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