class Cms::Article < ApplicationRecord
  belongs_to :user
  has_many :comments
  has_many :tags
end
