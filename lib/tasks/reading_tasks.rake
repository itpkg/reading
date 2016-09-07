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

        bk.subject = meta.subjects.empty? ? ' ': meta.subjects.first.content
        bk.date = meta.date.content
        bk.save


        book.each_content do |page|
          name = page.entry_name
          puts "find page #{name}"
          pg = Reading::Page.where(book_id: bk.id, entry_name: name).first
          unless pg
            pg = Reading::Page.new
          end
          pg.media_type = page.media_type
          pg.entry_name = name
          pg.body = page.read
          pg.book_id = bk.id

          pg.save
        end

      end

    end
  end
end

