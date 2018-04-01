package handlers

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/bluele/mecab-golang"
	"strings"
	"fmt"
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
		yomigana := yomiganafy(params.Name)

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
	m, err := mecab.New("-Owakati")
	if err != nil {
		panic(err) // TODO: echoのraise errorどうやるのか
	}
	defer m.Destroy()

	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(name)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	result := make([]string, 0)

	node := tg.ParseToNode(lt)
	for {
		feature := strings.Split(node.Feature(), ",")
		if feature[0] == "名詞" {
			fmt.Println("surface: ", node.Surface())
			fmt.Println("feature: ", node.Feature())
			result = append(result, feature[7]) // @NOTE 7番目が読み. (8番目は口語の読み)
		}
		if node.Next() != nil {
			break
		}
	}

	return strings.Join(result, " ")
}