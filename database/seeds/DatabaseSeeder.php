<?php

use Illuminate\Database\Seeder;
use Illuminate\Database\Eloquent\Model;

class DatabaseSeeder extends Seeder {

	/**
	 * Run the database seeds.
	 *
	 * @return void
	 */
	public function run()
	{
		Model::unguard();

		DB::table('users')->delete();
		
		App\User::create([
			'email' => env('APP_EMAIL'),
			'password' => Hash::make('dummy'),
		]);
	}

}
