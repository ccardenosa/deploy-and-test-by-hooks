package backend

import (
	"encoding/base64"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Config struct {
	ListenUri string
}

type TestResult struct {
	Pass      bool   `json:"pass"`
	B64Stdout string `json:"b64Stdout"`
}

type TestType struct {
	Name    string         `json:"name" binding:"required"`
	Type    string         `json:"type" binding:"required"`
	Path    string         `json:"path" binding:"required"`
	Context map[string]any `json:"context"`
	Result  TestResult
	Tests   []TestType `json:"tests"`
}

var tests = make(map[string]TestType)

func StartBackend(config Config) {

	r := gin.Default()
	// r.GET("/run/tests", listDevelHandler)
	// r.POST("/run/tests/:name", postRunTestHandler)
	r.POST("/run/tests", postRunTestsHandler)
	r.GET("/tests/results", listTestResultsLangHandler)
	// r.GET("/run/tests/:test", getDevelHandler)
	r.Run(config.ListenUri)
}

func postRunTestsHandler(c *gin.Context) {
	var t TestType
	c.ShouldBindJSON(&t)
	_, exists := tests[t.Name]
	if !exists {
		tests[t.Name] = t
		err := runTests(tests)
		if err != nil {
			c.String(401, "Test Failed: ", err)
		}
	}
	c.String(200, "Success")
}

func listTestResultsLangHandler(c *gin.Context) {
	c.JSON(200, gin.H{"results": tests})
}

func runTests(t map[string]TestType) error {
	var tr TestResult
	for k, v := range t {
		if v.Type == "test" {
			log.Printf("Running test %v in context[%v]...", v.Name, v.Context)
			tr.Pass, tr.B64Stdout = runFromShell(v.Path)
			v.Result = tr
		} else {
			for idx, vv := range v.Tests {
				log.Printf("Running test %v in context[%v] for TS[%v]...", vv.Name, vv.Context, v.Name)
				tr.Pass, tr.B64Stdout = runFromShell(filepath.Join(v.Path, vv.Path))
				vv.Result = tr
				t[k].Tests[idx] = vv
				if !tr.Pass {
					v.Result = tr
				}
			}
		}
		t[k] = v
	}
	return nil
}

func runFromShell(scriptFile string) (bool, string) {
	log.Println("Running script... ", scriptFile)
	cmd := exec.Command(scriptFile)

	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	log.Println("StdOut: ", output)
	return true, base64.StdEncoding.EncodeToString([]byte(output))
}
