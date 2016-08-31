class CreateEpubPages < ActiveRecord::Migration[5.0]
  def change
    create_table :epub_pages do |t|
      t.string :media_type, null:false
      t.string :entry_name, null:false
      t.binary :body, null:false
      t.timestamps
    end

    add_reference :epub_pages, :epub_books, foreign_key: true
    add_index :epub_pages, [:entry_name, :epub_books_id], unique: true
    add_index :epub_pages, :entry_name
  end
end
