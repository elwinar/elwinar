/*global module:false*/
module.exports = function(grunt) {

	// Project configuration.
	grunt.initConfig({
		concat: {
			options: {
				separator: ';',
			},
			build: {
				src: [
				'./bower_components/jquery/dist/jquery.js',
				'./app/scripts/script.js'
				],
				dest: './public/script.js',
			},
		},
		less: {
			build: {
				files: {
					"public/style.css": "app/styles/style.less"
				},
			},
		},
		uglify: {
		},
		watch: {
			styles: {
				files: 'app/styles/*.less',
				tasks: ['less'],
			},
			scripts: {
				files: 'app/scripts/*.js',
				tasks: ['concat', 'uglify'],
			},
		},
	});

	grunt.loadNpmTasks('grunt-contrib-concat');
	grunt.loadNpmTasks('grunt-contrib-uglify');
	grunt.loadNpmTasks('grunt-contrib-less');
	grunt.loadNpmTasks('grunt-contrib-watch');

	grunt.registerTask('default', ['less','']);

};
