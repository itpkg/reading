require 'redcarpet'
require 'htmlcompressor'
require 'fileutils'
require 'epub/parser'

namespace :convert do

  desc 'convert latex'
  task :latex do
    File.open('latex.sql', 'w') do |sf|
      sf.puts "DELETE FROM pages WHERE type = 'latex';"
      Dir.glob('src/**/*.tex').each do |src|
        puts "Find file #{src}"
        dst = "#{src[0..-5]}.xhtml"
        puts "Write #{dst}"
        puts `pandoc -s #{src} -o #{dst}`
        sf.puts "INSERT INTO pages(title, type) VALUES('#{src[4..-5]}', 'latex');"
      end
    end
  end

  desc 'convert pictures'
  task :pictures do
    %w(png jpg jpeg gif ico).each do |ext|
      Dir.glob("src/**/*.#{ext}").each do |src|
        puts "Find file #{src}"
        dst = "dst/#{src[4..-1]}"
        check_path dst
        puts "Write file #{dst}"
        FileUtils.cp src, dst
      end
    end
  end

  desc 'convert xhtml files'
  task :xhtml do
    Dir.glob('src/**/*.xhtml').each do |src|
      puts "Find file #{src}"

      dst = "dst/#{src[4..-7]}.html"
      check_path dst

      cp = HtmlCompressor::Compressor.new
      File.open(src, 'r') do |xf|
        File.open(dst, 'w') do |df|
          puts "Write file #{dst}"
          df.write cp.compress(Nokogiri::HTML(xf).at('body').inner_html)
        end
      end


    end
  end

  def check_path(f)
    r = File.dirname f
    unless File.directory?(r)
      FileUtils.mkdir_p r
    end
  end

  # epubinfo aaa.epub
  desc 'convert epub files'
  task :epub do
    File.open('epub.sql', 'w') do |sf|
      sf.puts "DELETE FROM BOOKS WHERE type = 'epub';"
      Dir.glob('src/**/*.epub').each do |src|
        puts "Find book #{src}"

        dst = "src/#{src[4..-6]}"
        if Dir.exists?(dst)
          puts "#{dst} exists, ignore."
        else
          puts `mkdir -p #{dst} && unzip -q #{src} -d #{dst}`
        end

        book = EPUB::Parser.parse src
        sf.puts "INSERT INTO books(url, title, author, type) VALUES('#{src[4..-6]}', '#{book.metadata.title}', '#{book.metadata.creators.first.content}', 'epub');"
        # book.each_page_on_spine do |page|
        #   #puts "Content #{page.content_document.nokogiri}"
        # end
      end

    end
  end


  desc 'convert markdown files'
  task :markdown do
    File.open('markdown.sql', 'w') do |sf|
      sf.puts "DELETE FROM pages WHERE type = 'markdown';"
      Dir.glob('src/**/*.md').each do |src|
        puts "Find file #{src}"
        dst = "dst/#{src[4..-4]}.html"

        md = Redcarpet::Markdown.new(Redcarpet::Render::HTML, autolink: true, tables: true)
        cp = HtmlCompressor::Compressor.new
        if File.exists?(dst)
          puts "#{dst} exists, ignore."
        else
          check_path dst
          File.open(src, 'r') do |mf|
            File.open(dst, 'w') do |hf|
              hf.write cp.compress(md.render(mf.read))
            end
          end
        end
        sf.puts "INSERT INTO pages(title, type) VALUES('#{src[4..-4]}', 'markdown');"
      end

    end
  end

end
