require 'epub/parser'

namespace :epub do
  desc 'import books'
  task :scan, [:dir] => :environment do |_, args|
    root = File.join(Rails.application.root, args.dir, '**', '*.epub')
    puts "scan books from #{root}"
    Dir.glob(root).each do |file|
      puts "find file #{file}"
      book = EPUB::Parser.parse(file)

      puts book.ocf.container.rootfile.inspect



      meta = book.metadata

      bid = meta.identifiers.first.content
      bk = Epub::Book.where(identifier: bid).first
      unless bk
        bk = Epub::Book.new
      end

      bk.title = meta.title
      bk.creator = meta.creators.first.content
      bk.identifier = bid
      bk.language = meta.language.content
      bk.publisher = meta.publishers.first.content
      bk.subject = meta.subjects.first.content
      bk.date = meta.date.content
      bk.home = book.rootfiles.first.full_path
      bk.save


      book.each_content do |page|
        name = page.entry_name
        puts "find page #{name}"
        pg = Epub::Page.where(epub_books_id: bk.id, entry_name: name).first
        unless pg
          pg = Epub::Page.new
        end
        pg.media_type = page.media_type
        pg.entry_name = name
        pg.body = page.read
        pg.epub_books_id = bk.id

        pg.save
      end




    end
  end
end
