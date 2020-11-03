package fullcontact

import "strings"

type TagOptions func(*Tag)

type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewTag(options ...TagOptions) *Tag {
	tag := &Tag{}

	for _, opts := range options {
		opts(tag)
	}

	return tag
}

func WithTagKey(key string) TagOptions {
	return func(tag *Tag) {
		tag.Key = key
	}
}

func WithTagValue(value string) TagOptions {
	return func(tag *Tag) {
		tag.Value = value
	}
}

func (tag *Tag) isValid() bool {
	return isPopulated(tag.Key) && isPopulated(tag.Value) && !strings.Contains(tag.Key, "'")
}
