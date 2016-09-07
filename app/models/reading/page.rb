module Reading
  class Page < ApplicationRecord
    belongs_to :book, class_name: 'Reading::Book'
  end
end
