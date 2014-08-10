<?php

use Illuminate\Database\Migrations\Migration;

class UsersTable extends Migration {

	/**
	 * Run the migrations.
	 *
	 * @return void
	 */
	public function up()
	{
        Schema::create('users', function($table)
        {
            $table->increments('id');
            $table->string('username', 50)->unique();
            $table->string('password', 60);
            $table->boolean('administrator');
            $table->timestamps();
        });
        
        $user = new User();
        $user->username = 'elwinar';
        $user->password = Hash::make('el01ro');
        $user->administrator = true;
        $user->save();
	}

	/**
	 * Reverse the migrations.
	 *
	 * @return void
	 */
	public function down()
	{
        Schema::drop('users');
	}

}