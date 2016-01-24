import gulp from 'gulp';
import babel from 'gulp-babel';

gulp.task('build', () => {
  return gulp.src('src/app.js')
    .pipe(babel({
        presets: ['es2015']
    }))
    .pipe(gulp.dest('dist'));
});

gulp.task('watch', ['build'], () => {
  gulp.watch('./src/*.js', ['build']);
});

gulp.task('default', ['watch']);
