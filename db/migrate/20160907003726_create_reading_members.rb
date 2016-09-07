class CreateReadingMembers < ActiveRecord::Migration[5.0]
  def change
    create_table :reading_members do |t|
      t.string :email, unique:true, null:false
      t.timestamps
    end
  end
end
