class InitialMigrations < ActiveRecord::Migration
  def change
    create_table :characters do |t|
      t.integer  :level,        default: 1
      t.integer  :energy,       default: 0
      t.integer  :ep,           default: 0
      t.datetime :ep_updated_at
      t.integer  :health,       default: 0
      t.integer  :hp,           default: 0
      t.datetime :hp_updated_at
      t.integer  :experience,   default: 0
      t.integer  :points,       default: 0
      t.integer  :basic_money,  default: 0
      t.integer  :vip_money,    default: 0

      t.timestamps null: false
    end
  end
end
