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

var eco = require('gulp-eco'); //forked https://github.com/chizhovdee/gulp-eco

var sass = require('gulp-sass');

var root = ".";
var build_path = "../server/assets";
var watch_js_path = root + "/js/**/*.coffee";
var watch_eco_path = root + "./client/js/views/**/*.eco";
var vendor_path = root + "/js/vendor";

var compiled_js = "app.js";

gulp.task('clean', function(cb) {
  cb.force = true;
  return del([
    build_path + "/js",
    build_path + "/css"
    ], cb);
});

gulp.task('copy-files', ['copy-vendor-js'], function() {

});

gulp.task('copy-vendor-js', function() {
  return gulp.src(
    [
      vendor_path + "/jquery.js",
      vendor_path + "/spine.js",
      vendor_path + "/underscore.js"
    ]
  ).pipe(gulp.dest(build_path + "/js"));
});


gulp.task('compile-eco', function () {
  return gulp.src(root + '/js/views/**/*.eco')
    .pipe(eco())
    .pipe(concat('templates.js'))
    .pipe(gulp.dest(build_path + "/js"));
});

gulp.task('compile-coffee', function() {
  return browserify(root + "/js/main.coffee")
    .transform(coffeeify)
    .bundle()
    .pipe(vinyl(compiled_js))
    .pipe(buffer())
    //.pipe(uglify())
    .pipe(gulp.dest(build_path + "/js"));
});

gulp.task('compile-js', ['compile-eco', 'compile-coffee'], function() {
  return gulp.src([
      build_path + "/js/templates.js",
      build_path + "/js/" + compiled_js
    ])
    .pipe(concat(compiled_js))
    .pipe(gulp.dest(build_path + "/js"))
    .on('end', function(){
      del(build_path + "/js/templates.js", {force: true})
    });
});

gulp.task('compile-sass', function () {
  gulp.src(root + '/css/**/*.scss')
    .pipe(sass().on('error', sass.logError))
    .pipe(gulp.dest(build_path + "/css"))
    .on('end', function(){
      del([
          build_path + "/css/**/*",
          "!" + build_path + "/css/app.css"
        ],
        {force: true})
    });
});


gulp.task('watch', function() {
  gulp.watch([watch_js_path, watch_eco_path], ['compile-js']);
});

gulp.task('default', ['clean', 'compile-js', 'compile-sass', "copy-files"]);

