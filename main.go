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
		Title:       "WordSuggestions",
		Description: "Give alternative word suggestions based on the input word.",
		Author:      "ielab",
		Version:     "3.Dec.2019",
		ProjectURL:  "ielab.io/searchrefiner",
	}
}

var Wordsuggestion = WordSuggestionPlugin{}
