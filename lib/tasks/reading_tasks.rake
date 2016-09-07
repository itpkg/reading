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
        bk.subject = meta.subjects.first.content
        bk.date = meta.date.content

        book.each_content do |page|
          puts "find page #{page.entry_name}"
          if page.media_type == 'application/x-dtbncx+xml'
            bk.home = page.read
            break
          end
        end
        bk.save

      end

    end
  end
end

