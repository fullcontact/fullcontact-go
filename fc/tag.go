package fullcontact

import "strings"

//Tags

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

//Tags Request

type TagsRequestOption func(tr *TagsRequest)

type TagsRequest struct {
	RecordId string `json:"recordId,omitempty"`
	Tags     []*Tag `json:"tags,omitempty"`
}

func NewTagsRequest(option ...TagsRequestOption) (*TagsRequest, error) {
	tagsRequest := &TagsRequest{}

	for _, opt := range option {
		opt(tagsRequest)
	}
	err := validateTagsRequest(tagsRequest)
	if err != nil {
		tagsRequest = nil
	}
	return tagsRequest, err
}

func validateTagsRequest(tagsRequest *TagsRequest) error {
	if !isPopulated(tagsRequest.RecordId) {
		return NewFullContactError("RecordId must be present for creating Tags")
	}
	if len(tagsRequest.Tags) < 1 {
		return NewFullContactError("Tags must be populated in Tags Create request")
	}
	for _, tag := range tagsRequest.Tags {
		if !tag.isValid() {
			return NewFullContactError("Both Key and Value must be populated for adding a Tag")
		}
	}
	return nil
}

func WithTag(tag *Tag) TagsRequestOption {
	return func(tagsRequest *TagsRequest) {
		if tagsRequest.Tags == nil {
			tagsRequest.Tags = make([]*Tag, 0)
		}
		tagsRequest.Tags = append(tagsRequest.Tags, tag)
	}
}

func WithTags(tags []*Tag) TagsRequestOption {
	return func(tagsRequest *TagsRequest) {
		if tagsRequest.Tags == nil {
			tagsRequest.Tags = make([]*Tag, 0)
		}
		tagsRequest.Tags = append(tagsRequest.Tags, tags...)
	}
}

func WithRecordIdForTags(recordId string) TagsRequestOption {
	return func(tagsRequest *TagsRequest) {
		tagsRequest.RecordId = recordId
	}
}

//Tags Response

type TagsResponse struct {
	RecordId  string `json:"recordId"`
	PartnerId string `json:"partnerId"`
	Tags      []Tag  `json:"tags"`
}
