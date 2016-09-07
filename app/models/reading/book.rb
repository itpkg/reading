module Reading
  class Book < ApplicationRecord
    validates :creator, :language, :subject, :publisher, presence: true
    validates :identifier, presence: true, uniqueness: true

    has_many :notes, class_name: 'Reading::Note'
  end
end
