require_dependency "reading/application_controller"

module Reading
  class HomeController < ApplicationController
    def index
      @books = Book.order(rate: :desc).limit(32)
      @title = t 'reading.index'
    end

    def page
      page = Page.where(book_id: params[:book_id], entry_name: "#{params[:name]}.#{params[:format]}").first
      send_data page.body, type: page.media_type, disposition: 'inline'
    end
  end
end
