package pathx

import "regexp"

// IsTemplateVariable 函数会返回 true，如果文本是一个模板变量
// 文本必须以点号开头，并且是一个有效的模板。
func IsTemplateVariable(text string) bool {
	match, _ := regexp.MatchString(`(?m)^{{(\.\w+)+}}$`, text)
	return match
}

// TemplateVariable 函数返回模板的变量名。
func TemplateVariable(text string) string {
	if IsTemplateVariable(text) {
		return text[3 : len(text)-2]
	}
	return ""
}
