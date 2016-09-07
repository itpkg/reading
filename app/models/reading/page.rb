module Reading
  class Page < ApplicationRecord
    belongs_to :book, class_name: 'Reading::Book'

    validates :media_type, :entry_name, :book_id, :body, presence: true
  end
end
