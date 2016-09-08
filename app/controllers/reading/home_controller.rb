require_dependency "reading/application_controller"

module Reading
  class HomeController < ApplicationController
    def index
      @books = Book.order(rate: :desc).limit(32)
      @title = t 'reading.index'
    end

    def page

      @page = Page.where(book_id: params[:book_id], entry_name: "#{params[:name]}.#{params[:format]}").first
      if @page.is_html?
        render 'page'
      else
        send_data @page.body, type: @page.media_type, disposition: 'inline'
      end
    end

    def dict
      db = "#{Rails.root}/tmp/dict"
      # @info = `sdcv --data-dir=#{db} -l`
      case request.method_symbol
        when :post
          kw = params[:keywords]
          if kw =~ /^\p{Word}+$/u
            @result = `sdcv --data-dir=#{db} #{kw}`
          else
            flash[:alert] = t 'reading.errors.bad_input'
          end
        else
      end

    end

    def search
      @items = Object.const_get(params.require(:resource)).search(params.require(:keywords))
    end
  end
end
