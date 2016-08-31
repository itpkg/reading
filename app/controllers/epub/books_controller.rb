class Epub::BooksController < ApplicationController
  def index
    @books = Epub::Book.order(rate: :desc).page params[:page]
  end

  def show
    @page = Epub::Page.where(epub_books_id: params[:id], media_type: 'application/x-dtbncx+xml').first
    @doc =Nokogiri::XML(@page.body)
  end
end
