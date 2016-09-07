require_dependency "reading/application_controller"

module Reading
  class HomeController < ApplicationController
    def index
      @books = Book.order(rate: :desc).limit(32)
      @title = t 'reading.index'
    end
  end
end
