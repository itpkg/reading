class CreateCmsTags < ActiveRecord::Migration[5.0]
  def change
    create_table :cms_tags do |t|
      t.string :name, null: false
      t.string :lang, null: false, limit: 5, default: :en
      t.timestamps
    end
    add_index :cms_tags, :name
    add_index :cms_tags, :lang
    add_index :cms_tags, [:name, :lang], unique: true

    create_join_table :cms_articles, :cms_tags
  end
end
