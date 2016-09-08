require 'elasticsearch/model'

module Reading
  class Page < ApplicationRecord
    include Elasticsearch::Model
    include Elasticsearch::Model::Callbacks

    belongs_to :book, class_name: 'Reading::Book'

    validates :media_type, :entry_name, :book_id, presence: true

    def as_indexed_json(options={})
      as_json(only: [:body, :title, :book_id, :entry_name])
    end

    def is_html?
      self.media_type == 'application/xhtml+xml'
    end

    def is_home?
      self.media_type == 'application/x-dtbncx+xml'
    end
  end
end
