require 'elasticsearch/model'

module Reading
  class Note < ApplicationRecord
    include Elasticsearch::Model
    include Elasticsearch::Model::Callbacks

    belongs_to :user
    belongs_to :book, class_name: 'Reading::Book'

    validates :user_id, :book_id, :body, presence: true

    def as_indexed_json(options={})
      as_json(only: [:body])
    end
  end
end
