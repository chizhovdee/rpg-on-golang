class InitialMigrations < ActiveRecord::Migration
  def change
    create_table :characters do |t|
      t.integer :ep
      t.integer :energy
      t.integer :hp
      t.integer :health
      t.integer :experience, default: 0
      t.integer :points,     default: 0

      t.timestamps null: false
    end
  end
end
