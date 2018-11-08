package stanford_corenlp

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

type CoreNLPTagger struct {
	modelPath   string
	properties  string
	java        string
	javaOptions []string
	separator   string
	encoding    string
}

type Result struct {
	result string
	Word   string
}

func NewCoreNLPTagger(m, p string) (*CoreNLPTagger, error) {
	separator := ":"
	if runtime.GOOS == "windows" {
		separator = ";"
	}

	ner := &CoreNLPTagger{
		java:        "java",
		encoding:    "utf8",
		javaOptions: []string{"-mx3g"},
		separator:   separator,
	}

	if err := ner.setModelPath(m); err != nil {
		return nil, err
	}
	if err := ner.setProperties(p); err != nil {
		return nil, err
	}

	return ner, nil
}

func (tagger *CoreNLPTagger) setModelPath(m string) error {
	tagger.modelPath = m
	return nil
}

func (tagger *CoreNLPTagger) setProperties(p string) error {
	if _, err := os.Stat(p); err != nil {
		return errors.New("properties not found (invalid path)")
	}
	tagger.properties = p
	return nil
}

func (tagger *CoreNLPTagger) setJavaPath(j string) {
	tagger.java = j
}

func (tagger *CoreNLPTagger) setJavaOptions(options []string) {
	tagger.javaOptions = options
}

func (tagger *CoreNLPTagger) setEncoding(e string) {
	tagger.encoding = e
}

func (tagger *CoreNLPTagger) parse(out string) []*Result {
	//words := strings.Split(out, " ")
	res := make([]*Result, 0)
	fmt.Println(out)
	return res
}

func (tagger *CoreNLPTagger) Tag(input string) ([]*Result, error) {
	var (
		tmp  *os.File
		err  error
		args []string
	)

	if tmp, err = ioutil.TempFile("", "nlptemp"); err != nil {
		return nil, err
	}
	defer os.Remove(tmp.Name())
	if _, err = tmp.WriteString(input); err != nil {
		return nil, err
	}

	args = append(tagger.javaOptions, []string{
		"-cp",
		tagger.modelPath + tagger.separator,
		"edu.stanford.nlp.pipeline.StanfordCoreNLP",
		"-props", tagger.properties,
		"-file", tmp.Name(),
		"-encoding", tagger.encoding,
		//"-tokenize.whitespace", "-ssplit.eolonly",
		"-outputDirectory", "out/",
	}...)

	cmd := exec.Command(tagger.java, args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s: %s", err, stderr.String())
	}
	return tagger.parse("meep"), nil
}
