require 'elasticsearch/model'

module Reading
  class Page < ApplicationRecord
    include Elasticsearch::Model
    include Elasticsearch::Model::Callbacks

    belongs_to :book, class_name: 'Reading::Book'

    validates :media_type, :entry_name, :book_id, :body, presence: true
  end
end
