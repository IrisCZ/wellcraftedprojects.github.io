var gulp = require('gulp');
var concat = require('gulp-concat')
var watch = require('gulp-watch');

gulp.task('default', function() {

});

var dependencies = [
  'node_modules/jquery/dist/jquery.min.js',
  'node_modules/underscore/underscore-min.js',
  'node_modules/backbone/backbone-min.js'
];

var sources = [
  'js/router.js',
  'js/views/*.js'
]

gulp.task('build-dependencies', function() {
  return gulp.src(dependencies)
    .pipe(concat('libs.js'))
    .pipe(gulp.dest('./js/lib/'));
});

gulp.task('build-js', function() {
  watch('js/views/**/*.js', function(){
    return gulp.src(sources)
        .pipe(concat('app.js'))
        .pipe(gulp.dest('./js/'));
  });
});