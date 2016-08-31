class Epub::BooksController < ApplicationController
  def index
    @books = Epub::Book.all
  end
end
