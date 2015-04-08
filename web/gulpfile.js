var gulp = require('gulp');
var concat = require('gulp-concat')
var watch = require('gulp-watch');

gulp.task('dependencies', function() {

  var dependencies = [
    'node_modules/jquery/dist/jquery.min.js',
    'node_modules/underscore/underscore-min.js',
    'node_modules/backbone/backbone-min.js',
    'js/vendor/**/*.js'
  ];

  return gulp.src(dependencies)
    .pipe(concat('libs.js'))
    .pipe(gulp.dest('./js/lib/'));
});

var sources = [
  'js/model/*.js',
  'js/views/*.js',
  'js/router.js'
]

gulp.task('refresh-js', function(){
  gulp.src(sources)
      .pipe(concat('app.js'))
      .pipe(gulp.dest('./js/'))
});

gulp.task('start', function() {
  gulp.start('dependencies')
  gulp.start('refresh-js')
  gulp.watch(sources,['refresh-js'])
});