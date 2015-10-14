var gulp = require('gulp');
var less = require('gulp-less');
var autoprefixer = require('gulp-autoprefixer');
var minify = require('gulp-minify-css');
var concat = require('gulp-concat');
var rename = require('gulp-rename');
var uglify = require('gulp-uglify');
var plumber = require('gulp-plumber');
var uncss = require('gulp-uncss');

gulp.task('styles', function() {
	gulp.src([
		'elwinar.less',
		'bower_components/prism/themes/prism.css',
	])
	.pipe(plumber())
	.pipe(concat('app.less'))
	.pipe(less({
		paths: [ '.' ],
	}))
	.pipe(rename({extname: '.css'}))
	.pipe(uncss({html: ['*.html']}))
	.pipe(autoprefixer())
	.pipe(minify())
	.pipe(gulp.dest('public/'));
});

gulp.task('fonts', function() {
	gulp.src([
		'bower_components/bootstrap/fonts/*',
		'bower_components/fontawesome/fonts/*',
		'fonts/*',
	])
	.pipe(gulp.dest('public/fonts/'));
});

gulp.task('scripts', function () {
	gulp.src([
		'bower_components/jquery/dist/jquery.js',
		'bower_components/bootstrap/dist/js/bootstrap.js',
		'bower_components/speakingurl/speakingurl.min.js',
		'bower_components/prims/components/prism-bash.js',
		'bower_components/prism/components/prism-core.js',
		'bower_components/prism/components/prism-clike.js',
		'bower_components/prism/components/prism-javascript.js',
		'bower_components/prism/components/prism-markup.js',
		'bower_components/prism/components/prism-php.js',
		'bower_components/prism/components/prism-ruby.js',
		'*.js',
	])
	.pipe(plumber())
	.pipe(concat('app.js'))
	.pipe(uglify())
	.pipe(rename({extname: '.js'}))
	.pipe(gulp.dest('public'));
});

gulp.task('watch', function () {
	gulp.watch('*.js', ['scripts']);
	gulp.watch('*.less', ['styles']);
	gulp.watch('*.html', ['styles']);
});

gulp.task('default', ['styles', 'fonts', 'scripts']);
