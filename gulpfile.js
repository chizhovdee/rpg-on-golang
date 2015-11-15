// npm install --save-dev gulp gulp-coffeelint gulp-uglify browserify browserify-shim coffeeify del vinyl-buffer vinyl-source-stream gulp-concat

var gulp       = require('gulp');
var coffeelint = require('gulp-coffeelint');
var uglify     = require('gulp-uglify');

var browserify = require('browserify');
var coffeeify  = require('coffeeify');
var del        = require('del');
var buffer     = require('vinyl-buffer');
var vinyl      = require('vinyl-source-stream');

var concat = require('gulp-concat');

var root_path = "./client";
var tmp_path = root_path + "/assets";
var js_path = root_path + "/js/*.coffee";
var res_path = root_path + "/res/*.coffee";

var compiled_js = "main.js";

gulp.task('clean', function(cb) {
  del(tmp_path, cb);
});

gulp.task('compile-coffee', function() {
  return browserify(root_path + "/js/main.coffee")
    .transform(coffeeify)
    .bundle()
    .pipe(vinyl(compiled_js))
    .pipe(buffer())
    //.pipe(uglify())
    .pipe(gulp.dest(tmp_path + "/js"));
});

gulp.task('compile-js', ['compile-coffee'], function() {
  return gulp.src(tmp_path + "/js/" + compiled_js)
    .pipe(concat(compiled_js))
    .pipe(gulp.dest(tmp_path + "/js"));
});

// не используется
gulp.task('watch', function() {
  gulp.watch(js_path, ['compile-js']);
});

gulp.task('default', ['clean', 'compile-js']);

