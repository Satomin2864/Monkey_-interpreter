package token

type TokenType string

// トークンの種類とリテラルを持っている
// トークン構造体を定義
type Token struct {
  Type TokenType
  Literal string
}

// Monkey言語で用いる演算子、リテラルetcを定数宣言
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
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"

  LT = "<"
  GT = ">"

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
  TRUE = "TRUE"
  FALSE = "FALSE"
  IF = "IF"
  ELSE = "ELSE"
  RETURN = "RETURN"

  // 比較用の演算子
  EQ = "=="
  NOT_EQ = "!="
)

// キーをstringとしてTokenTypeを紐付けたmap型
var keyword = map[string]TokenType{
  "fn":     FUNCTION,
  "let":    LET,
  "true":   TRUE,
  "false":  FALSE,
  "if":     IF,
  "else":   ELSE,
  "return": RETURN,
}

// 引数としてもらったindentをキーとするmapがkeywordに存在するなら、
// 対応したtokenを返す
// なければ、インデントtokenを返す
func LookupIdent(indent string) TokenType {
  if tok, ok := keyword[indent]; ok {
    return tok
  }
  return INDENT
}
