package main

import (
	"fmt"
	"github.com/kyajpauley/stanford-corenlp"
)

func main() {
	var (
		tagger *stanford_corenlp.CoreNLPTagger
		res    []*stanford_corenlp.Result
		err    error
	)

	modelPath := "examples/jar/*"
	propertiesPath := "examples/StanfordCoreNLP-chinese.properties"

	tagger, err = stanford_corenlp.NewCoreNLPTagger(modelPath, propertiesPath)
	if err != nil {
		fmt.Println(err)
	}

	res, err = tagger.Tag("我来 到 北京 清华大学")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}
