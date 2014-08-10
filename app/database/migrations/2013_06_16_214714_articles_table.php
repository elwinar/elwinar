<?php

use Illuminate\Database\Migrations\Migration;

class ArticlesTable extends Migration {

	/**
	 * Run the migrations.
	 *
	 * @return void
	 */
	public function up()
	{
        Schema::create('articles', function($table)
        {
            $table->increments('id');
            $table->string('slug', 50)->unique();
            $table->string('title', 50);
            $table->string('tagline', 300);
            $table->text('text');
            $table->text('tags');
            $table->timestamps();
        });
	}

	/**
	 * Reverse the migrations.
	 *
	 * @return void
	 */
	public function down()
	{
        Schema::drop('articles');
	}

}