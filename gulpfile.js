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
		'src/styles/*.less',
		'node_modules/highlightjs/styles/atom-one-light.css',
	])
	.pipe(plumber())
	.pipe(concat('app.less'))
	.pipe(less({
		paths: [ 'src/styles' ],
	}))
	.pipe(rename({extname: '.css'}))
	.pipe(uncss({html: ['src/views/*.html']}))
	.pipe(autoprefixer())
	.pipe(minify())
	.pipe(gulp.dest('build/public/'));
});

gulp.task('fonts', function() {
	gulp.src([
		'node_modules/bootstrap/fonts/*',
		'node_modules/fontawesome/fonts/*',
		'src/fonts/*',
	])
	.pipe(gulp.dest('build/public/fonts/'));
});

gulp.task('scripts', function () {
	gulp.src([
		'node_modules/jquery/dist/jquery.js',
		'node_modules/bootstrap/dist/js/bootstrap.js',
		'node_modules/speakingurl/speakingurl.min.js',
		'node_modules/highlightjs/highlight.pack.js',
		'src/scripts/*.js',
	])
	.pipe(plumber())
	.pipe(concat('app.js'))
	.pipe(uglify())
	.pipe(gulp.dest('build/public'));
});

gulp.task('watch', function () {
	gulp.watch('src/scripts/*.js', ['scripts']);
	gulp.watch('src/styles/*.less', ['styles']);
	gulp.watch('src/views/*.html', ['styles']);
});

gulp.task('default', ['styles', 'fonts', 'scripts']);
