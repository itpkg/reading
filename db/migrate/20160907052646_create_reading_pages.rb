class CreateReadingPages < ActiveRecord::Migration[5.0]
  def change
    create_table :reading_pages do |t|
      t.string :media_type, null: false
      t.string :entry_name, null: false
      t.binary :body, null: false
      t.integer :book_id, foreign_key: true, null:false
      t.timestamps
    end

    add_index :reading_pages, [:entry_name, :book_id], unique: true
    add_index :reading_pages, :entry_name
  end
end
