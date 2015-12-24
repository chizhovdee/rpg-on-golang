/**
 *  gulp task list:
 *  compile-assets
 *  production:compile-assets
 *  compile-js
 *  compile-css
 *  watch
 *  clean
 */

//Error: watch ENOSPC fix with
//echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p

var gulp       = require('gulp');

var _ = require("underscore");

var coffeelint = require('gulp-coffeelint'); // проверка синтаксиса
var uglify     = require('gulp-uglify'); // сжатие кофе
var rev = require('gulp-rev'); // Static asset revisioning by appending content hash to filenames
var browserify = require('browserify'); // common js require
var del        = require('del'); // for deleting files
var buffer     = require('vinyl-buffer'); // temporary vinyl format
var vinyl      = require('vinyl-source-stream');
var include = require("gulp-include");
var concat = require('gulp-concat'); // files concatination
var minify = require("gulp-minify-css");

var coffeeify  = require('coffeeify'); // coffee compilation
var eco = require('gulp-eco'); // eco compilation    forked https://github.com/chizhovdee/gulp-eco
var sass = require('gulp-sass'); // sass compilation

var notify = require('gulp-notify');
var shell = require('gulp-shell');
var gulpif = require('gulp-if');

var map     = require('map-stream');
var yaml    = require('js-yaml');
var gutil = require("gulp-util");

var root = ".";

var server_assets_path = "../server/public/assets";
var build_path = server_assets_path;

var coffee_files_path = root + "/js/**/*.coffee";
var eco_files_path = root + "/js/views/**/*.eco";
var vendor_path_js = root + "/js/vendor";
var sass_files_path = root + "/css/**/*.scss";

var coffe_main_path = root + "/js/main.coffee";
var compiled_js = "application.js";
var compiled_eco_js = "templates.js";
var compiled_css = "application.css";

var assets = [
  build_path + "/" + compiled_js,
  build_path + "/" + compiled_css
];

// common for production and development
function compileEco(){
  return gulp.src(eco_files_path)
    .pipe(eco({basePath: 'js/views'}))
    .pipe(concat(compiled_eco_js))
    .pipe(gulp.dest(build_path));
}

function compileCss(){
  return gulp.src(sass_files_path)
    .pipe(include())
    .pipe(gulpif('application.scss', sass()))
    .on('error', notify.onError(function (error) {
      return 'An error occurred while compiling sass.\nLook in the console for details.\n' + error;
    }))
    .pipe(gulp.dest(build_path))
    .on('end', function(){
      del([
        build_path + "/*.css", "!" + build_path + "/" + compiled_css,
        build_path + "/*scss"], {force: true});
    });
}

gulp.task('coffee-lint', function () {
  return gulp.src(coffee_files_path)
    .pipe(coffeelint())
    .pipe(notify({
      title: "Coffeelint error",
      message: function(file){
        if(!file.coffeelint.success){
          return "Found " + file.coffeelint.errorCount + " Errors and " +
            file.coffeelint.warningCount + " Warnings \n" + "File path: " + file.relative;
        }
      }
    })).pipe(coffeelint.reporter());
});

gulp.task("compile-css", function() {
  return compileCss()
});

gulp.task('compile-coffee', ['coffee-lint'], function() {
  return browserify(coffe_main_path)
    .transform(coffeeify)
    .bundle()
    .pipe(vinyl(compiled_js))
    .pipe(buffer())
    .pipe(gulp.dest(build_path));
});

gulp.task('compile-eco-and-coffee', ['compile-coffee'], function () {
  return compileEco();
});

// development
gulp.task('watch-locales', ["compile-assets"], function(){
  gulp.watch('../locales/**/*.yml', ["locales"])
});

gulp.task('watch-css', ['watch-locales'], function() {
  gulp.watch(sass_files_path, ['compile-css']);
});

gulp.task('watch', ['watch-css'], function() {
  gulp.watch([coffee_files_path, eco_files_path], ['compile-js']);
});

gulp.task('compile-js', ['compile-eco-and-coffee'], function() {
  return gulp.src([
    vendor_path_js + "/preloadjs.min.js",
    vendor_path_js + "/jquery.js",
    vendor_path_js + "/underscore.js",
    vendor_path_js + "/spine.js",
    vendor_path_js + "/visibility.min.js",
    vendor_path_js + "/i18n.js",
    build_path + "/" + compiled_eco_js,
    build_path + "/" + compiled_js
  ])
    .pipe(concat(compiled_js))
    .pipe(gulp.dest(build_path)).on('end', function(){
      del(build_path + "/" + compiled_eco_js, {force: true})
    });
});

gulp.task("compile-css-and-compile-js", ["compile-js"], function(){
  return compileCss()
});

gulp.task("compile-assets", ["compile-css-and-compile-js"]);

// production
gulp.task('compile-and-compress-js', ['compile-eco-and-coffee'], function() {
  return gulp.src([
    build_path + "/" + compiled_eco_js,
    build_path + "/" + compiled_js
  ])
    .pipe(concat(compiled_js))
    .pipe(uglify())
    .pipe(gulp.dest(build_path)).on('end', function(){
      del(build_path + "/" + compiled_eco_js, {force: true})
    });
});

gulp.task('production:compile-js', ['compile-and-compress-js'], function() {
  return gulp.src([
    vendor_path_js + "/jquery.min.js",
    vendor_path_js + "/underscore.min.js",
    vendor_path_js + "/spine.min.js",
    vendor_path_js + "/visibility.min.js",
    build_path + "/" + compiled_js
  ])
    .pipe(concat(compiled_js))
    .pipe(gulp.dest(build_path));
});

gulp.task("production:compile-css-and-compile-js", ['production:compile-js'], function() {
  return gulp.src(sass_files_path)
    .pipe(include())
    .pipe(sass())
    .pipe(gulp.dest(build_path))
    .pipe(minify())
    .pipe(gulp.dest(build_path))
    .on('end', function(){
      del([build_path + "/*.css", "!" + build_path + "/" + compiled_css], {force: true});
    });
});

gulp.task("production:compile-assets", ["production:compile-css-and-compile-js"], function(){
  return gulp.src(assets)
    .pipe(buffer())
    .pipe(rev())
    .pipe(gulp.dest(build_path))
    .pipe(rev.manifest({path: 'manifest.json'}))
    .pipe(gulp.dest(build_path));
});


// services tasks
gulp.task('compress-spine', function(){
  return gulp.src(vendor_path_js + "/spine.js")
    .pipe(uglify())
    .pipe(concat("spine.min.js"))
    .pipe(buffer())
    .pipe(gulp.dest(vendor_path_js))
});


gulp.task('clean', function(cb) {
  cb.force = true;

  return del([build_path], cb);
});

gulp.task("server", function() {
  gulp.src('').pipe(shell('cd ../server && fresh'));
});

gulp.task('locales', function(){
  var data = {
    ru: {},
    en: {}
  };

  gulp.src('../locales/**/*.yml')
    .pipe(map(function(file,cb){
      if (file.isNull()) return cb(null, file); // pass along
      if (file.isStream()) return cb(new Error("Streaming not supported"));

      var json;
      var lng = "";

      try {
        json = yaml.load(String(file.contents.toString('utf8')));

        if(json.ru != null){
          lng = "ru";
        } else if (json.en != null) {
          lng = "en";
        } else {
          return cb(new Error("Language not supported"));
        }

        _.extend(data[lng], json[lng]);

      } catch(e) {
        console.log(e);
        console.log(json);
      }

      file = new gutil.File({
        cwd: "",
        base: "",
        path: lng + ".json",
        contents: new Buffer(JSON.stringify(data[lng]))
      });

      cb(null, file);
    }))
    .pipe(gulp.dest('../server/public/locales'));
});