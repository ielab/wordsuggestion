package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ielab/searchrefiner"
)

type WordSuggestionPlugin struct {
}

func (WordSuggestionPlugin) Serve(s searchrefiner.Server, c *gin.Context) {
	rawQuery := c.PostForm("query")
	lang := c.PostForm("lang")
	c.Render(http.StatusOK, searchrefiner.RenderPlugin(searchrefiner.TemplatePlugin("plugin/wordsuggestion/index.html"), struct {
		searchrefiner.Query
		View string
	}{searchrefiner.Query{QueryString: rawQuery, Language: lang}, c.Query("view")}))
	return
}

func (WordSuggestionPlugin) PermissionType() searchrefiner.PluginPermission {
	return searchrefiner.PluginUser
}

func (WordSuggestionPlugin) Details() searchrefiner.PluginDetails {
	return searchrefiner.PluginDetails{
		Title:       "KeywordSuggest",
		Description: "Provide query suggestions for a search.",
		Author:      "ielab",
		Version:     "23.Jan.2020",
		ProjectURL:  "ielab.io/searchrefiner",
	}
}

var Wordsuggestion = WordSuggestionPlugin{}
