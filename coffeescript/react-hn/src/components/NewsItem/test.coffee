'use strict'

$ = require 'jquery'
NewsItem = require '.'
React = require 'react'

$.ajax
    url: '/json/items.json'
  .then (items) ->
    console.log 'items', items
    React.render React.createElement(NewsItem, {item: items[0], rank: 1}), $('#content')[0]
