'use strict'

gulp = require 'gulp'
$ = (require 'gulp-load-plugins')()

paths =
  coffee: 'public/coffee/*.coffee'

gulp.task 'coffee', ->
  gulp.src paths.coffee
    .pipe $.plumber()
    .pipe $.coffeelint()
    .pipe $.coffee()
    .pipe gulp.dest 'public/scripts/'

gulp.task 'watch', ->
  gulp.watch paths.coffee, ['coffee']

gulp.task 'build', ['coffee']
