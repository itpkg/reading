require 'epub/parser'

namespace :reading do
  namespace :books do

    desc 'Scan e-books'
    task :scan, [:dir] => :environment do |_, args|
      root = File.join(Rails.application.root, args.dir, '**', '*.epub')
      puts "scan books from #{root}"
      Dir.glob(root).each do |file|
        puts "find file #{file}"
        book = EPUB::Parser.parse(file)

        meta = book.metadata



        bid = meta.identifiers.first.content
        bk = Reading::Book.where(identifier: bid).first
        unless bk
          bk = Reading::Book.new
        end

        bk.title = meta.title
        bk.creator = meta.creators.first.content

        bk.identifier = bid
        bk.language = meta.language.content
        bk.publisher = meta.publishers.first.content

        bk.subject = meta.subjects.empty? ? '' : meta.subjects.first.content
        bk.date = meta.date.content
        bk.save
        unless bk.valid?
          raise bk.errors.inspect
        end


        book.each_content do |page|
          name = page.entry_name
          puts "find page #{name}"
          pg = Reading::Page.where(book_id: bk.id, entry_name: name).first
          unless pg
            pg = Reading::Page.new
          end
          pg.media_type = page.media_type
          pg.entry_name = name

          if pg.is_html?
            doc =Nokogiri::XML(page.read)
            pg.title = doc.css('head title').first.content
            pg.body = doc.css('body').first.inner_html
          else
            pg.payload = page.read
          end

          pg.book_id = bk.id

          pg.save
          unless pg.valid?
            raise pg.errors.inspect
          end
        end

      end

    end
  end
end

