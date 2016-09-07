require_dependency "reading/application_controller"

module Reading
  class NotesController < ApplicationController
    def index
      @notes = Note.order(updated_at: :desc).page params[:page]
    end

    def new
      @note = Note.new
      @note.book_id = params[:book_id]
      authorize @note
      render 'form'
    end

    def create
      @note = Note.new params.require(:note).permit(:body, :book_id)
      authorize @note
      @note.user = current_user
      if @note.save
        redirect_to book_path(@note.book)
      else
        render 'form'
      end

    end

    def edit
      @note = Note.find params[:id]
      authorize @note
      render 'form'
    end

    def update
      @note = Note.find params[:id]
      authorize @note
      if @note.update params.require(:note).permit(:body)
        redirect_to book_path(@note.book)
      else
        render 'form'
      end

    end


    def destroy
      @note = Note.find params[:id]
      authorize @note
      redirect_to book_path(@note.book)
    end
  end
end
