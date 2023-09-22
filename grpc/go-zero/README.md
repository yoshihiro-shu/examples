# HOW TO USE

## API定義の作成

apiディレクトリを作成し、その中にhelloworld.apiという名前のファイルを作成します。そして、以下の内容をhelloworld.apiファイルに追加します。

```api
type (
  Request {
    Name string `path:"name,options=[you,me]"` // parameters are auto validated
  }

  Response {
    Message string `json:"message"`
  }
)

service greet-api {
  @handler GreetHandler
  get /greet/from/:name(Request) returns (Response)
}
```

## コードの生成

ターミナルで、以下のコマンドを実行して、API定義からコードを生成します。

```zsh
goctl api go -api api/helloworld.api -dir .
```
