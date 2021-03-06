'use strict'

gulp = require 'gulp'
$ = (require 'gulp-load-plugins')()
browserify = require 'browserify'
watchify = require 'watchify'
transform = require 'vinyl-transform'

srcPath = 'src'
buildPath = 'build'

paths =
  mainCoffee: [
    "#{srcPath}/coffee/index.coffee"
    "#{srcPath}/coffee/list.coffee"
    "#{srcPath}/coffee/detail.coffee"
  ]
  mainJade: [
    "#{srcPath}/index.jade"
    "#{srcPath}/list.jade"
    "#{srcPath}/detail.jade"
  ]

gulp.task 'coffee', ->
  browserified = transform (filename) ->
    browserify filename
      .bundle()
  gulp
    .src paths.mainCoffee
    .pipe browserified
    .pipe $.rename extname: '.js'
    .pipe gulp.dest "#{buildPath}/js"


gulp.task 'browserify', ->
  bundler = (filename, isWatch) ->
    options =
      debug: true
      cache: {}
      packageCache: {}
      fullPaths: true
      extensions: ['.coffee', '.js', '.jade']

    if isWatch
      watchify browserify(filename, options)
        .on 'update', bundle
    else
      browserify(filename, options)
  
  bundle = (filename) ->
    return transform (filename) ->
      console.log 'compiling', filename
      bundler(filename, true).bundle()

  gulp
    .src paths.mainCoffee
    .pipe bundle()
    .pipe $.rename extname: '.js'
    .pipe gulp.dest "#{buildPath}/js"



gulp.task 'jade', ->
  gulp
    .src paths.mainJade
    .pipe $.jade pretty: true
    .pipe gulp.dest "#{buildPath}"



gulp.task 'watch', ->
  gulp.watch paths.mainCoffee, ['coffee']
  gulp.watch paths.mainJade, ['jade']

gulp.task 'build', ['coffee', 'jade']
