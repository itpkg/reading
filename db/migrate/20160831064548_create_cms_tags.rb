class CreateCmsTags < ActiveRecord::Migration[5.0]
  def change
    create_table :cms_tags do |t|
      t.string :name, null: false

      t.integer :rate, null: false, default: 0

      t.timestamps
    end

    add_index :cms_tags, :name, unique: true
    create_join_table :cms_articles, :cms_tags

  end
end
