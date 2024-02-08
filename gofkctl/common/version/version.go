package version

import (
	"encoding/json"
	"strings"
)

// BuildVersion is the version of gofkctl.
const BuildVersion = "1.0.0"

var tag = map[string]int{"pre-alpha": 0, "alpha": 1, "pre-release": 2, "beta": 3, "released": 4, "": 5}

// GetVersion returns BuildVersion
func GetVersion() string {
	return BuildVersion
}

// IsVersionGreaterThan 比较两个版本号
// 它接受两个字符串参数：version（当前 goctl 版本）和 target（目标版本），并返回一个布尔值，表明当前版本是否大于目标版本。
func IsVersionGreaterThan(version, target string) bool {
	versionNumber, versionTag := convertVersion(version)
	targetVersionNumber, targetTag := convertVersion(target)
	if versionNumber > targetVersionNumber {
		return true
	} else if versionNumber < targetVersionNumber {
		return false
	} else {
		// 首先比较主版本号，如果相等，则比较标签的优先级。
		return tag[versionTag] > tag[targetTag]
	}
}

// 解析一个版本字符串，并将其分解为数值部分和标签部分
// 数值部分被转换为 float64 类型，标签保持为字符串。如果版本号中包含多个点号，函数会忽略第一个点号之后的点号
func convertVersion(version string) (versionNumber float64, tag string) {
	splits := strings.Split(version, "-")
	tag = strings.Join(splits[1:], "")
	var flag bool
	numberStr := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}

		if r == '.' {
			if flag {
				return '_'
			}
			flag = true
			return r
		}
		return '_'
	}, splits[0])
	numberStr = strings.Replace(numberStr, "_", "", -1)
	versionNumber, _ = json.Number(numberStr).Float64()
	return
}
