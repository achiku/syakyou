'use strict'

gulp = require 'gulp'
$ = (require 'gulp-load-plugins')()
browserSync = require 'browser-sync'
runSequence = require 'run-sequence'

paths =
  coffee: 'public/coffee/*.coffee'


gulp.task 'browser-sync', ->
  browserSync
    logLevel: 'debug'
    proxy: '0.0.0.0:3000'


gulp.task 'coffee', ->
  gulp.src paths.coffee
    .pipe $.cjsx({bare: true}).on('error', $.util.log)
    .pipe gulp.dest 'public/scripts/'


gulp.task 'server', ->
  runSequence 'browser-sync', ['coffee', 'watch']

gulp.task 'watch', ->
  gulp.watch paths.coffee, ['coffee', browserSync.reload]

gulp.task 'build', ['coffee']
