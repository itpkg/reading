class CreateCmsComments < ActiveRecord::Migration[5.0]
  def change
    create_table :cms_comments do |t|
      t.text :body, null: false
      t.timestamps
    end

    add_reference :cms_comments, :user, foreign_key: true
    add_reference :cms_comments, :cms_article, foreign_key: true
  end
end
