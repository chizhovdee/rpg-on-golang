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

var build_path = "../server/assets";
var watch_js_path = "./js/**/*.coffee";
var res_path = "./res/*.coffee";
var vendor_path = "./js/vendor";

var compiled_js = "game.js";

gulp.task('clean', function(cb) {
  del(build_path, cb);
});

gulp.task('copy-files', ['copy-vendor-js'], function() {
    return gulp.src(['project.json'])
    .pipe(gulp.dest(build_path));
});

gulp.task('copy-vendor-js', function() {
  return gulp.src(
    [
      vendor_path + "/zepto.js",
      vendor_path + "/cocos2d-js-v3.8.js"
    ]
  ).pipe(gulp.dest(build_path + "/js"));
});

gulp.task('compile-coffee', function() {
  return browserify("./js/main.coffee")
    .transform(coffeeify)
    .bundle()
    .pipe(vinyl(compiled_js))
    .pipe(buffer())
    //.pipe(uglify())
    .pipe(gulp.dest(build_path + "/js"));
});

gulp.task('compile-js', ['compile-coffee'], function() {
  return gulp.src(build_path + "/js/" + compiled_js)
    .pipe(concat(compiled_js))
    .pipe(gulp.dest(build_path + "/js"));
});

// не используется
gulp.task('watch', function() {
  gulp.watch(watch_js_path, ['compile-js']);
});

gulp.task('default', ['clean', 'compile-js', "copy-files"]);

