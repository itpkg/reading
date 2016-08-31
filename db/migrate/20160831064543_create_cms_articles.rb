class CreateCmsArticles < ActiveRecord::Migration[5.0]
  def change
    create_table :cms_articles do |t|
      t.string :title, null: false
      t.string :locale, null: false, limit: 5, default: :en
      t.string :summary
      t.text :body, null: false

      t.integer :rate, null: false, default: 0

      t.timestamps
    end

    add_reference :cms_articles, :user, foreign_key: true
    add_index :cms_articles, :locale

  end
end
