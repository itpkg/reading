class User < ApplicationRecord
  include RailsSettings::Extend

  rolify

  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable and :omniauthable
  devise :invitable, :database_authenticatable, :registerable,
         :recoverable, :rememberable, :trackable, :validatable,
         :confirmable, :lockable, :timeoutable,
         :invitable
  include DeviseTokenAuth::Concerns::User

  belongs_to :role

  def send_devise_notification(notification, *args)
    devise_mailer.send(notification, self, *args).deliver_later
  end


  before_create :skip_duplicate_devise_confirmation_email
  # Fixes problem with duplicate account confirmation emails
  def skip_duplicate_devise_confirmation_email
    skip_confirmation_notification!
  end
end
