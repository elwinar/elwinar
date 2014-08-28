/*global module:false*/
module.exports = function(grunt) {

	// Project configuration.
	grunt.initConfig({
		less: {
			style: {
				files: {
					'public/less.css': 'app/styles/style.less'
				},
			},
		},
		concat: {
			script: {
				src: [
					'bower_components/jquery/dist/jquery.js',
					'bower_components/highlightjs/highlight.pack.js',
					'app/scripts/*.js'
				],
				dest: 'public/script.js',
			},
			style: {
				src: [
					'bower_components/highlightjs/styles/github.css',
					'public/less.css',
					'app/styles/*.css',
				],
				dest: 'public/style.css',
			},
		},
		uglify: {
			options: {
				mangle: true
			},
			script: {
				files: {
					'public/script.min.js': ['public/script.js']
				}
			},
		},
		cssmin: {
			style: {
				src: 'public/style.css',
				dest: 'public/style.min.css',
			}
		},
		watch: {
			styles: {
				files: 'app/styles/*.less',
				tasks: ['style'],
			},
			scripts: {
				files: 'app/scripts/*.js',
				tasks: ['script'],
			},
		},
	});

	grunt.loadNpmTasks('grunt-contrib-less');
	grunt.loadNpmTasks('grunt-contrib-concat');
	grunt.loadNpmTasks('grunt-contrib-uglify');
	grunt.loadNpmTasks('grunt-contrib-cssmin');
	grunt.loadNpmTasks('grunt-contrib-watch');
	
	grunt.registerTask('style', ['less:style','concat:style','cssmin:style']);
	grunt.registerTask('script', ['concat:script','uglify:script']);
	grunt.registerTask('default', ['style','script']);

};
