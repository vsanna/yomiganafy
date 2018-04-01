package handlers

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/ikawaha/kagome.ipadic/tokenizer"
	"strings"
	"fmt"
	"log"
)

// json: jsonでdataを受け取ったときに底からparseする
//   -d '{"name":"Joe","email":"joe@labstack"}'
// form: formのデータとして受け取った気にそこからparseする
//   -d 'name=Joe'
// query: urlに埋め込む形で受け取った際にそこからparseする
//   http://localhost:1234/yomiganafy?name=tom

type YomiganafyParams struct {
	Name string `json:"name" form:"name" query:"name"` // {name: 'hoge'} でpostを受け付ける
}

const failure = "failure"

/*
## リクエストの解釈
- header
	- [ ]
- params
	- 構造体を定義し、それにc.Bindすることで取り扱えるようになる

## レスポンスの生成
- format: c.String / c.JSON / c.HTML
	- [ ] 例えばファイル(pdf, img)をDLさせるには？
	- [ ] HTMLのテンプレートの使い方は
- status: http.StatusOK, http.StatusInternalServerError, http.StatusNotFound, ...
- header:
	- [ ] まだ不明
- body:
	- c.JSON(status, body)という感じっぽい。とりあえずJSONの場合はmapを渡せばいいようなので便利

*/

func Yomiganafy() echo.HandlerFunc {
	return func(c echo.Context) error {
		params := new(YomiganafyParams) // &YomiganafyParams{} と同じ...
		// c.Bindで構造体にparamsを埋め込んでいる.
		// [ ] 仮に値がなければここで起こってくれるのかな？
		if err := c.Bind(params); err != nil {
			return err
		}

		// ここでparamsを元にビジネスロジックを実行
		//yomigana := &Yomigana{Name: "name", Yomi: "yomi"}
		yomigana := yomiganafy(params.Name)

		// @TODO 上手いハンドリング
		if yomigana.Yomi == failure {

		}

		fmt.Println(yomigana)

		return c.JSON(http.StatusOK, map[string]interface{}{"result": yomigana})
	}
}



/*****************
# business logic
これはどこに置くといいのだろう
******************/
type Yomigana struct {
	Name string
	Yomi string
}

func yomiganafy(name string) *Yomigana {
	yomi := parse(name)

	return &Yomigana{
		Name: name,
		Yomi: yomi,
	}
}

func parse(name string) string {
	t := tokenizer.New()
	tokens := t.Tokenize(name)

	result := make(MyArray, 0)

	for _, token := range tokens {
		features := token.Features()
		log.Println("token: ", token)
		log.Println("features: ", features)

		// BOS: Begin Of Sentence, EOS: End Of Sentence.
		// これskipしないとうごかない。ようわからん
		if token.Class == tokenizer.DUMMY {
			continue
		}

		if features[0] == "名詞" {
			// ヨミがない場合はそもそもlenすら異なる。
			if len(features) < 8 {
				result = append(result, failure)
			} else {
				fmt.Println("surface: ", token.Surface)
				fmt.Println("feature: ", token.Features())
				result = append(result, features[7]) // @NOTE 7番目が読み. (8番目は口語の読み)
			}
		}
	}

	if result.contains(failure) {
		return failure
	}

	result_as_string := make([]string, len(result))
	for i, arg := range result { result_as_string[i] = arg.(string) }
	return strings.Join(result_as_string, " ")
}


type MyArray []interface{}

func (arr MyArray) contains(el interface{}) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}

	return false
}