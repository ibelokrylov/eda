package helpers

import "regexp"

type Replacer struct {
	patterns []string
}

func NewReplacer(patterns []string) *Replacer {
	return &Replacer{
		patterns: patterns,
	}
}

func (r *Replacer) Replace(text string) (string, error) {
	for _, pattern := range r.patterns {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return "", err
		}
		text = re.ReplaceAllStringFunc(text, func(s string) string {
			matched := re.FindStringSubmatch(s)
			if len(matched) == 3 {
				return matched[1] + "XXX"
			}
			return s
		})
	}
	return text, nil
}
