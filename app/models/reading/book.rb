module Reading
  class Book < ApplicationRecord
    validates :creator, :language, :subject, :publisher, presence: true
    validates :identifier, presence: true, uniqueness: true
  end
end
