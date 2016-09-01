class CreateLogs < ActiveRecord::Migration[5.0]
  def change
    create_table :logs do |t|
      t.string :message, null: false
      t.column :flag, :integer, default: 0, null:false
      t.datetime :created_at, null:false
    end
    add_reference :logs, :user, foreign_key: true

  end
end
