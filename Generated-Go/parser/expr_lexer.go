// Code generated from parser/Expr.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ExprLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func exprlexerLexerInit() {
	staticData := &ExprLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "':'", "'='", "'*'", "'+'", "", "", "'INT'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "IDENT", "NUM", "INT_TYPE", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "IDENT", "NUM", "INT_TYPE", "COMMENT",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 69, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 5, 4, 30, 8, 4, 10, 4, 12,
		4, 33, 9, 4, 1, 5, 1, 5, 3, 5, 37, 8, 5, 1, 5, 1, 5, 5, 5, 41, 8, 5, 10,
		5, 12, 5, 44, 9, 5, 3, 5, 46, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 5, 7, 56, 8, 7, 10, 7, 12, 7, 59, 9, 7, 1, 7, 1, 7, 1, 8, 4,
		8, 64, 8, 8, 11, 8, 12, 8, 65, 1, 8, 1, 8, 0, 0, 9, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 1, 0, 6, 1, 0, 97, 122, 3, 0, 48,
		57, 65, 90, 97, 122, 1, 0, 49, 57, 1, 0, 48, 57, 2, 0, 10, 10, 13, 13,
		2, 0, 9, 10, 32, 32, 74, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0, 3, 21,
		1, 0, 0, 0, 5, 23, 1, 0, 0, 0, 7, 25, 1, 0, 0, 0, 9, 27, 1, 0, 0, 0, 11,
		45, 1, 0, 0, 0, 13, 47, 1, 0, 0, 0, 15, 51, 1, 0, 0, 0, 17, 63, 1, 0, 0,
		0, 19, 20, 5, 58, 0, 0, 20, 2, 1, 0, 0, 0, 21, 22, 5, 61, 0, 0, 22, 4,
		1, 0, 0, 0, 23, 24, 5, 42, 0, 0, 24, 6, 1, 0, 0, 0, 25, 26, 5, 43, 0, 0,
		26, 8, 1, 0, 0, 0, 27, 31, 7, 0, 0, 0, 28, 30, 7, 1, 0, 0, 29, 28, 1, 0,
		0, 0, 30, 33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 10,
		1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 46, 5, 48, 0, 0, 35, 37, 5, 45, 0,
		0, 36, 35, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 42,
		7, 2, 0, 0, 39, 41, 7, 3, 0, 0, 40, 39, 1, 0, 0, 0, 41, 44, 1, 0, 0, 0,
		42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1,
		0, 0, 0, 45, 34, 1, 0, 0, 0, 45, 36, 1, 0, 0, 0, 46, 12, 1, 0, 0, 0, 47,
		48, 5, 73, 0, 0, 48, 49, 5, 78, 0, 0, 49, 50, 5, 84, 0, 0, 50, 14, 1, 0,
		0, 0, 51, 52, 5, 45, 0, 0, 52, 53, 5, 45, 0, 0, 53, 57, 1, 0, 0, 0, 54,
		56, 8, 4, 0, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0,
		0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61,
		6, 7, 0, 0, 61, 16, 1, 0, 0, 0, 62, 64, 7, 5, 0, 0, 63, 62, 1, 0, 0, 0,
		64, 65, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1,
		0, 0, 0, 67, 68, 6, 8, 0, 0, 68, 18, 1, 0, 0, 0, 7, 0, 31, 36, 42, 45,
		57, 65, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ExprLexerInit initializes any static state used to implement ExprLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewExprLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExprLexerInit() {
	staticData := &ExprLexerLexerStaticData
	staticData.once.Do(exprlexerLexerInit)
}

// NewExprLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewExprLexer(input antlr.CharStream) *ExprLexer {
	ExprLexerInit()
	l := new(ExprLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ExprLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Expr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExprLexer tokens.
const (
	ExprLexerT__0     = 1
	ExprLexerT__1     = 2
	ExprLexerT__2     = 3
	ExprLexerT__3     = 4
	ExprLexerIDENT    = 5
	ExprLexerNUM      = 6
	ExprLexerINT_TYPE = 7
	ExprLexerCOMMENT  = 8
	ExprLexerWS       = 9
)