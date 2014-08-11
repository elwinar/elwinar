/*global module:false*/
module.exports = function(grunt) {

	// Project configuration.
	grunt.initConfig({
		// Metadata.
		pkg: grunt.file.readJSON('package.json'),
		less: {
			options: {
				ieCompat: true,
			},
			build: {
				files: {
					'public/css/style.css': 'public/less/*',
				},
			},
		},
		watch: {
			grunt: {
				files: [ 'Gruntfile.js' ],
				options: {
					reload: true,
				},
			},
			less: {
				files: [ 'public/less/*' ],
				tasks: [ 'less' ],
				options: {
					spawn: false,
				},
			},
		},
	});

	// These plugins provide necessary tasks.

	grunt.loadNpmTasks('grunt-contrib-less');
	grunt.loadNpmTasks('grunt-contrib-watch');

	// Default task.
	grunt.registerTask('default', ['less']);

};
