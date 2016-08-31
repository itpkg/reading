class CreateEpubPages < ActiveRecord::Migration[5.0]
  def change
    create_table :epub_pages do |t|
      t.string :name, null:false
      t.string :title, null:false
      t.text :body, null:false
      t.timestamps
    end

    add_reference :epub_pages, :epub_books, foreign_key: true
    add_index :epub_pages, [:name, :epub_books_id], unique: true
  end
end
