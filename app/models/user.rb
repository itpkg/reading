class User < ApplicationRecord
  include DeviseTokenAuth::Concerns::User
  include RailsSettings::Extend

  rolify

  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable and :omniauthable
  devise :invitable, :database_authenticatable, :registerable,
         :recoverable, :rememberable, :trackable, :validatable,
         :confirmable, :lockable, :timeoutable,
         :invitable

  belongs_to :role


end
