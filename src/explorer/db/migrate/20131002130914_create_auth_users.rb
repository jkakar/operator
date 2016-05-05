class CreateAuthUsers < ActiveRecord::Migration
  def change
    create_table :auth_users do |t|
      t.string :email, null: false
      t.string :name, null: false
      t.string :uid, null: false
      t.timestamps null: false
    end
    add_index :auth_users, :email, unique: true
    add_index :auth_users, :uid, unique: true
  end
end
