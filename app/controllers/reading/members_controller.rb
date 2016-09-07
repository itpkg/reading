require_dependency "reading/application_controller"

module Reading
  class MembersController < ApplicationController

    def create
      m = Member.new params.permit(:email)
      if m.valid?
        # todo send mail
        case params[:flag]
          when 'subscribe'
          when 'unsubscribe'
          else
        end
        flash[:notice] = ' '
      else
        flash[:alert] = m.errors
      end
      redirect_to new_member_path
    end
  end

  def show
    # todo parse jwt
  end
end
