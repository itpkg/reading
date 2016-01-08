require 'redcarpet'
require 'htmlcompressor'
require 'fileutils'

namespace :convert do

  desc 'convert epub files'
  task :epub do
    Dir.glob('src/**/*.epub').each do |src|
      puts "Fine file #{src}"
      dst = "dst/#{src[4..-6]}"
      if Dir.exists?(dst)
        puts "#{dst} exists, ignore."
      else
        puts `mkdir -p #{dst} && unzip -q #{src} -d #{dst}`
      end
    end
  end

  desc 'convert markdown files'
  task :markdown do
    Dir.glob('src/**/*.md').each do |src|
      puts "Fine file #{src}"
      dst = "dst/#{src[4..-4]}.html"

      md = Redcarpet::Markdown.new(Redcarpet::Render::HTML, autolink: true, tables: true)
      cp = HtmlCompressor::Compressor.new
      if File.exists?(dst)
        puts "#{dst} exists, ignore."
      else
        root = File.dirname dst
        unless File.directory?(root)
          FileUtils.mkdir_p root
        end
        File.open(src, 'r') do |mf|
          File.open(dst, 'w') do |hf|
            hf.write cp.compress(md.render(mf.read))
          end
        end
      end

    end
  end

end
