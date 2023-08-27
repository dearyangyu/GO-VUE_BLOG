package testdata

import (
	"Go-blog-server/models/res"
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

type ErrMap map[res.ErrorCode]string

const filename = "models/res/err_code.json"

func main() {
	byteData, err := os.ReadFile(filename)

	if err != nil {
		logrus.Error(err)
		return
	}
	var errMap := ErrMap{}
	err = json.Unmarshal(byteData, &errMap)

	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(errMap)
	fmt.Println(errMap[re.SettingsError])
}
