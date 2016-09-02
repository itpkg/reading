class Api::LeaveWordsController < ApplicationController

  def index

    if current_user.is_admin?
      render json: LeaveWord.order(created_at: :desc).limit(200)
    else
      head :forbidden
    end
  end

  def create
    lw = LeaveWord.create params.permit(:content)
    if lw.valid?
      render json: lw
    else
      render json:{errors:lw.errors.messages}
    end
  end

  def destroy
    if current_user.is_admin?
      lw = LeaveWord.find params[:id]
      lw.destroy
      render json: lw
    else
      head :forbidden
    end
  end
end
