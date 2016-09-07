module Reading
  class Member < ApplicationRecord
    validates :email, presence: true
    validates_format_of :email, :with => Devise::email_regexp
  end
end
