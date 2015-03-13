'use strict'
gulp = require 'gulp'
jade = require 'gulp-jade'


gulp.task 'jade', ->
  gulp.src './src/templates/*.jade'
    .pipe jade()
    .pipe gulp.dest './build/html'

gulp.task 'watch', ->
  gulp.watch './src/templates/*.jade', ['jade']


gulp.task 'default', ['watch']
