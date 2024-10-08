package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var MentionMessagePattern = regexp.MustCompile(`<at oc_id=["']([^"']+)["']></at>`)

// ReplaceMentionMessage 函数会根据idToName映射表中的对应关系，将给定文本中的提及消息替换为相应的名字。
// 该函数使用MentionMessagePattern正则表达式模式来查找提及消息，并使用"@{Name}"的格式将其替换为对应的名字。
// 如果在idToName映射表中找不到对应的名字，那么提及消息将保持不变。函数返回修改后的文本。
func ReplaceMentionMessage(text string, idToName map[string]string) string {
	return MentionMessagePattern.ReplaceAllStringFunc(
		text, func(match string) string {
			id := MentionMessagePattern.FindStringSubmatch(match)[1] // 提取出oc_id值
			if name, ok := idToName[id]; ok {
				// return "@" + name // 用@{Name}替换
				return name
			}
			return "" // 如果没有找到对应的name, 为了防止ai-chat无法处理, 直接不处理.
		},
	)
}

func ReplaceMentionMessageWithAt(text string, idToName map[string]string) string {
	return MentionMessagePattern.ReplaceAllStringFunc(
		text, func(match string) string {
			id := MentionMessagePattern.FindStringSubmatch(match)[1] // 提取出oc_id值
			if name, ok := idToName[id]; ok {
				return "@" + name // 用@{Name}替换
			}
			return "" // 如果没有找到对应的name, 为了防止ai-chat无法处理, 直接不处理.
		},
	)
}

// FindAllMentionMessagePositions 函数会返回给定文本中所有提及消息的位置。
// 该函数使用MentionMessagePattern正则表达式模式来查找提及消息，并返回所有匹配项的起始和结束位置的切片。
func FindAllMentionMessagePositions(text string) [][]int {
	return MentionMessagePattern.FindAllStringIndex(text, -1)
}

// byteIndexToRuneIndex 函数将字节索引转换为rune索引。
func byteIndexToRuneIndex(text string, byteIndex int) int {
	return len([]rune(text[:byteIndex]))
}

// SimplifyContent 函数会简化给定的内容, 确保不会截断提及消息, 并且每个@消息只计作5个字符, 最多返回50个字符.
func SimplifyContent(content string) string {
	if len(content) == 0 {
		return ""
	}

	remainingLength := 50
	positions := FindAllMentionMessagePositions(content)
	runes := []rune(content)
	length := len(runes)
	var simplifiedContent strings.Builder

	for i := 0; i < length && remainingLength > 0; {
		found := false
		for _, pos := range positions {
			start := byteIndexToRuneIndex(content, pos[0])
			end := byteIndexToRuneIndex(content, pos[1])
			if i == start {
				// 找到提及消息
				atMessage := runes[start:end]
				simplifiedContent.WriteString(string(atMessage))
				remainingLength -= 5
				i = end
				found = true
				break
			}
		}
		if !found {
			// 普通字符
			simplifiedContent.WriteRune(runes[i])
			remainingLength--
			i++
		}
	}

	if remainingLength == 0 && len(runes) > len([]rune(simplifiedContent.String())) {
		return fmt.Sprintf("%s...", simplifiedContent.String())
	}

	return simplifiedContent.String()
}
