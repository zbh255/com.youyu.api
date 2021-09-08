package controller

type TagsApi interface {
	CheckTag(text string)
	AddTag(text string)
}
