class CreateNotices < ActiveRecord::Migration[5.0]
  def change
    create_table :notices do |t|
      t.text :body, null: false
      t.string :lang, null: false, limit: 5, default: :en
      t.timestamps
    end
    add_index :notices, :lang
  end
end
