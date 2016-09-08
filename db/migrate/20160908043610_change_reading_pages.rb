class ChangeReadingPages < ActiveRecord::Migration[5.0]
  def change
    change_table :reading_pages do |t|
      t.rename :body, :payload
      t.string :title
      t.text :body
    end
    change_column :reading_pages, :payload, :binary, null: true
  end
end
