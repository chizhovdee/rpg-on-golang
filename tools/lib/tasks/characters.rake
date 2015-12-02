namespace :characters do
  desc "Create test character"
  task :create_test => :environment do
    Character.new.tap do |c|
      c.energy = c.ep = 10
      c.health = c.hp = 100
      c.save!
    end
  end
end