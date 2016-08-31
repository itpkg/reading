class HomeController < ApplicationController
  def index

  end
  def about
    @key = :about
    render 'page'
  end
  def help
    @key = :help
    render 'page'
  end

  def faq
    @key = :faq
    render 'page'
  end
end
