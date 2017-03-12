package engine

import (
	"fmt"
	"sync"
	"time"

	"github.com/cheggaaa/pb"

	"github.com/lucaszanella/go-freeling/nlp"
	. "github.com/lucaszanella/go-freeling/terminal"
)

type Engine struct {
	semaphore *sync.Mutex
	NLP       *nlp.NLPEngine
	Lang      string
	Path      string
	Ready     bool
}

func NewEngine() *Engine {
	return &Engine{
		semaphore: new(sync.Mutex),
		Ready:     false,
		Lang:      "en",
		Path:      "./",
	}
}

func (e *Engine) SetLanguage(lang string) {
	e.Lang = lang
}

func (e *Engine) SetPath(path string) {
	e.Path = path
}

func (e *Engine) InitNLP() {
	e.semaphore.Lock()
	defer e.semaphore.Unlock()
	if e.Ready {
		return
	}
	Infoln("Init Natural Language Processing Engine")
	initialized := false
	count := 80
	bar := pb.StartNew(count)
	bar.ShowPercent = true
	bar.ShowCounters = false

	inc := func() {
		for i := 0; i < 10; i++ {
			bar.Increment()
		}
	}

	start := time.Now().UnixNano()
	nlpOptions := nlp.NewNLPOptions(e.Path+"data", e.Lang, inc)
	nlpOptions.Severity = nlp.ERROR
	nlpOptions.TokenizerFile = "tokenizer.dat"
	nlpOptions.SplitterFile = "splitter.dat"
	nlpOptions.TaggerFile = "tagger.dat"
	nlpOptions.ShallowParserFile = "chunker/grammar-chunk.dat"
	nlpOptions.SenseFile = "senses.dat"
	nlpOptions.UKBFile = "" //"ukb.dat"
	nlpOptions.DisambiguatorFile = "common/knowledge.dat"

	macoOptions := nlp.NewMacoOptions(e.Lang)
	macoOptions.SetDataFiles("", e.Path+"data/common/punct.dat", 
							 e.Path+"data/"+e.Lang+"/dicc.src", 
							 "", "", e.Path+"data/"+e.Lang+"/locucions.dat", 
							 e.Path+"data/"+e.Lang+"/np.dat", "", 
							 e.Path+"data/"+e.Lang+"/probabilitats.dat")

	nlpOptions.MorfoOptions = macoOptions

	nlpEngine := nlp.NewNLPEngine(nlpOptions)

	stop := time.Now().UnixNano()
	delta := (stop - start) / (1000 * 1000)
	initialized = true
	bar.FinishPrint(fmt.Sprintf("Data loaded in %dms", delta))
	e.NLP = nlpEngine
	e.Ready = initialized
}
