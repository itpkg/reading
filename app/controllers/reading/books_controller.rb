require_dependency "reading/application_controller"

module Reading
  class BooksController < ApplicationController
    def index
      @books = Book.order(updated_at: :desc).page params[:page]
    end

    def show
      @page = Page.where(book_id: params[:id], media_type: 'application/x-dtbncx+xml').first
      @doc =Nokogiri::XML(@page.body)
    end

    def destroy
      book = Book.find params[:id]
      authorize book
      book.destroy
      redirect_to books_path
    end
  end
end
