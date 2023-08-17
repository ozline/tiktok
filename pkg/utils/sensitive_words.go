package utils

import (
	//"github.com/chain-zhang/pinyin"
	"regexp"
	"strings"
)

// HansCovertPinyin 中文汉字转拼音
//func HansCovertPinyin(contents []string) []string {
//	pinyinContents := make([]string, 0)
//	for _, content := range contents {
//		chineseReg := regexp.MustCompile("[\u4e00-\u9fa5]")
//		if !chineseReg.Match([]byte(content)) {
//			continue
//		}
//
//		// 只有中文才转
//		pin := pinyin.New(content)
//		pinStr, err := pin.Convert()
//		fmt.Println(content, "->", pinStr)
//		if err == nil {
//			pinyinContents = append(pinyinContents, pinStr)
//		}
//	}
//	return pinyinContents
//}

// TrieNode 敏感词前缀树节点
type TrieNode struct {
	childMap map[rune]*TrieNode // 本节点下的所有子节点
	//Data     string             // 在最后一个节点保存完整的一个内容
	End bool // 标识是否最后一个节点
}

// SensitiveTrie 敏感词前缀树
type SensitiveTrie struct {
	root *TrieNode
}

// NewSensitiveTrie 构造敏感词前缀树实例
func NewSensitiveTrie() *SensitiveTrie {
	return &SensitiveTrie{
		root: &TrieNode{End: false},
	}
}

// FilterSpecialChar 过滤特殊字符
func (st *SensitiveTrie) FilterSpecialChar(text string) string {
	text = strings.ToLower(text)
	text = strings.Replace(text, " ", "", -1) // 去除空格

	// 过滤除中英文及数字以外的其他字符
	otherCharReg := regexp.MustCompile("[^\u4e00-\u9fa5a-zA-Z0-9]")
	text = otherCharReg.ReplaceAllString(text, "")
	return text
}

// AddWord 添加敏感词
func (st *SensitiveTrie) AddWord(sensitiveWord string) {
	// 添加前先过滤一遍
	sensitiveWord = st.FilterSpecialChar(sensitiveWord)

	// 将敏感词转换成utf-8编码后的rune类型(int32)
	tireNode := st.root
	sensitiveChars := []rune(sensitiveWord)
	for _, charInt := range sensitiveChars {
		// 添加敏感词到前缀树中
		tireNode = tireNode.AddChild(charInt)
	}
	tireNode.End = true
	//tireNode.Data = sensitiveWord
}

// AddWords 批量添加敏感词
func (st *SensitiveTrie) AddWords(sensitiveWords []string) {
	for _, sensitiveWord := range sensitiveWords {
		st.AddWord(sensitiveWord)
	}
}

// Match 匹配敏感词
func (st *SensitiveTrie) Match(text string) bool {
	if st.root == nil {
		return false
	}

	// 过滤特殊字符
	filteredText := st.FilterSpecialChar(text)
	textChars := []rune(filteredText)
	for i, textLen := 0, len(textChars); i < textLen; i++ {
		trieNode := st.root.FindChild(textChars[i])
		if trieNode == nil {
			continue
		}

		// 匹配到了敏感词的前缀，从后一个位置继续
		j := i + 1
		for ; j < textLen && trieNode != nil; j++ {
			if trieNode.End {
				// 完整匹配到了敏感词
				return true
			}
			trieNode = trieNode.FindChild(textChars[j])
		}

		// 文本尾部命中敏感词情况
		if j == textLen && trieNode != nil && trieNode.End {
			return true
		}
	}
	return false
}

// AddChild 前缀树添加字节点
func (tn *TrieNode) AddChild(c rune) *TrieNode {

	if tn.childMap == nil {
		tn.childMap = make(map[rune]*TrieNode)
	}

	if trieNode, ok := tn.childMap[c]; ok {
		// 存在不添加了
		return trieNode
	} else {
		// 不存在
		tn.childMap[c] = &TrieNode{
			childMap: nil,
			End:      false,
		}
		return tn.childMap[c]
	}
}

// FindChild 前缀树查找字节点
func (tn *TrieNode) FindChild(c rune) *TrieNode {
	if tn.childMap == nil {
		return nil
	}

	if trieNode, ok := tn.childMap[c]; ok {
		return trieNode
	}
	return nil
}
