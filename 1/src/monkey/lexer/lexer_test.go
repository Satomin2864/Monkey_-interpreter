package lexer

import (
  "testing"

  "../token"
)
func TestNextToken(t *testing.T) {
  input :=`=+(){},;`
  let ten = 10;
  let add = fn(x, y) {
    x + y;
  };
  tests  := []struct {
    expectedType token.TokenType
    expectedLiteral string
  }{
    {token.ASSIGN, "="},
    {token.PLUS, "+"},
    {token.LPAREN, "("},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.RBRACE, "}"},
    {token.COMMA, ","},
    {token.SEMICOLON, ";"},
    {token.EOF, ""},
  }
  l := New(input)

  for i, tt := range tests {
    tok := l.NextToken()

    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
        i, tt.expectedLiteral, tok.Literal)
    }

    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
        i, tt.expectedLiteral, tok.Literal)
    }
  }
}

func (l *Lexer) readChar(){
  if l.readPosition >= len(l.input) {
    l.ch = 0
  } else {
    l.ch = l.input[l.readPosition]
  }
  l.position = l.readPosition
  l.readPosition += 1
}
