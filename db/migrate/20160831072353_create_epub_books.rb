class CreateEpubBooks < ActiveRecord::Migration[5.0]
  def change
    create_table :epub_books do |t|
      t.string :title, null:false
      t.string :identifier, null:false
      t.string :creator, null:false
      t.string :subject, null:false
      t.string :language, null:false, limit: 5
      t.string :publisher, null:false
      t.string :date, null:false
      t.string :home, null:false
      t.integer :rate, null:false, default:0
      t.timestamps
    end
    add_index :epub_books, :creator
    add_index :epub_books, :identifier, unique: true
    add_index :epub_books, :language
    add_index :epub_books, :subject
    add_index :epub_books, :publisher
  end
end
