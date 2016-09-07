module Reading
  class Book < ApplicationRecord
    # some book missing?
    # validates :creator, :language, :subject, :publisher, presence: true
    validates :identifier, presence: true, uniqueness: true

    has_many :notes, class_name: 'Reading::Note'
    has_many :pages, class_name: 'Reading::Page'
  end
end
