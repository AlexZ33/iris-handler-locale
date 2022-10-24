package iris_handler_locale

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/i18n"
	"github.com/pelletier/go-toml"
)

// 国际化
// Locale  usage :  app.UseGlobal(handler.Locale())
func Locale(config *toml.Tree) iris.Handler {
	conf := GetTree(config, "i18n")
	languages := GetStringArray(config, "languages", []string{"en-US", "zh-CN"})
	locales := map[string]string{}
	for _, lang := range languages {
		locales[lang] = "./config/locale/" + lang + ".ini"
	}
	return i18n.New(i18n.Config{
		Default:      GetString(conf, "default-locale", "en-US"),
		URLParameter: GetString(conf, "query-field", "locale"),
		Languages:    locales,
	})
}
