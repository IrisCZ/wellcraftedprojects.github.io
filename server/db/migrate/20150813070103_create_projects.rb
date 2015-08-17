class CreateProjects < ActiveRecord::Migration
  def change
    create_table :projects do |t|
      t.string :name
      t.string :author
      t.string :url
      t.text :description
      t.integer :positives
      t.integer :negatives
      t.string :image

      t.timestamps
    end
  end
end
