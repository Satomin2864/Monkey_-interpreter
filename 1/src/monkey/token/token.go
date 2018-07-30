package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  // 未知のtokenや文字列
  ILLEGAL = "ILLEGAL"
  // 終端文字
  EOF = "EOF"

  // 識別子 + リテラル
  INDENT = "INDENT" // add, foobar, x, y, etc...
  INT = "INT" // 12345...

  // 演算子
  ASSIGN = "="
  PLUS = "+"

  // デリミタ
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // キーワード
  FUNCTION = "FUNCTION"
  LET = "LET"
)
