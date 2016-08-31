class Epub::PagesController < ApplicationController
  def show
    bid = params[:bid]

    page = Epub::Page.where(epub_books_id: bid, entry_name: "#{params[:name]}.#{params[:format]}").first

    send_data page.body, type: page.media_type, disposition: 'inline'

  end
end
