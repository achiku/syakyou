'use strict'

$ = require 'jquery'
React = require 'react'
jade = require 'react-jade'
url = require 'url'
template = jade.compileFile "#{__dirname}/template.jade"

NewsItem = React.createClass
  getCommentText: ->
    commentText = 'discuss'
    if @props.item.kids and @props.item.kids.length
      commentText = @props.item.kids.length + ' comments'
    return commentText

  getDomain: ->
    return url.parse(@props.item.url).hostname

  render: ->
    template @


module.exports = NewsItem
