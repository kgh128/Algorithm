import (
    "sort"
    "strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    var str1, str2 string

    sort.Strings(word1)
    sort.Strings(word2)
    
    strings.Join(word1, str1)
    strings.Join(word2, str2)

    return str1 == str2
}
