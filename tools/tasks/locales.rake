require 'yaml'

def loop_hash(rst, r, h)
  h.each do |k, v|
    k = "#{r}.#{k}" unless r == ''
    if v.is_a?(Hash)
      loop_hash rst, k, v
    else
      rst[k] = v
    end
  end
end

desc 'build locales.sql'
task :locales do
  File.open('locales.sql', 'w') do |lf|
    lf.puts 'DELETE FROM locales;'
    Dir.glob('locales/**/*.yml').each do |fn|
      lang = fn[-9..-5]
      puts "Find language #{lang}."
      items = YAML.load File.read(fn)

      rst = {}
      loop_hash rst,'', items
      puts "#{rst.length} items."
      rst.each do |k, v|
        lf.puts "INSERT INTO locales(lang, code, message) VALUES('#{lang}', '#{k}', '#{escape_string v}');"
      end
    end
  end
end