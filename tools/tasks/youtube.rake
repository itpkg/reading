require 'json'
require 'open-uri'
require 'net/http'

def http_json_get(url)
  JSON.parse Net::HTTP.get(URI(url))
end

desc 'youtube videos crawler'
task :youtube do
  key = ENV['YOUTUBE_KEY']
  channels=[]

  ENV['YOUTUBE_USERS'].split(' ').each do |user|
    puts "抓取用户[#{user}]"
    http_json_get("https://www.googleapis.com/youtube/v3/channels?part=snippet&maxResults=50&forUsername=#{user}&key=#{key}").fetch('items').map do |channel|
      ch_sn = channel.fetch('snippet')

      puts "\t抓取频道[#{ch_sn.fetch('title')}]"
      channels << {
          id: channel.fetch('id'),
          type: 'youtube',
          title: ch_sn.fetch('title'),
          description: ch_sn.fetch('description'),
          playlist: http_json_get("https://www.googleapis.com/youtube/v3/playlists?part=snippet&channelId=#{channel.fetch('id')}&maxResults=50&key=#{key}").fetch('items').map do |playlist|
            pl_sn = playlist.fetch('snippet')
            puts "\t\t抓取播放列表[#{pl_sn.fetch('title')}]"
            {
                id: playlist.fetch('id'),
                title: pl_sn.fetch('title'),
                description: pl_sn.fetch('description'),
                videos: http_json_get("https://www.googleapis.com/youtube/v3/playlistItems?part=snippet&maxResults=50&playlistId=#{playlist.fetch('id')}&key=#{key}").fetch('items').map do |video|
                  v_sn = video.fetch('snippet')
                  {
                      id: v_sn.fetch('resourceId').fetch('videoId'),
                      title: v_sn.fetch('title'),
                      description: v_sn.fetch('description')
                  }
                end
            }
          end
      }

    end
  end

  File.open('youtube.sql', 'w') do |f|
    f.puts "DELETE FROM channels WHERE type = 'youtube';"
    f.puts "DELETE FROM playlists WHERE type = 'youtube';"
    f.puts "DELETE FROM videos WHERE type = 'youtube';"
    channels.each do |ch|
      f.puts "INSERT INTO channels(cid, type, title, description) VALUES('#{ch.fetch :id}', 'youtube', '#{ch.fetch :title}', '#{ch.fetch :description}');"
      ch.fetch(:playlist).each do |pl|
        f.puts "INSERT INTO playlists(cid, pid, type, title, description) VALUES('#{ch.fetch :id}','#{pl.fetch :id}', 'youtube', '#{pl.fetch :title}', '#{pl.fetch :description}');"
        pl.fetch(:videos).each do |v|
          f.puts "INSERT INTO videos(pid, vid, type, title, description) VALUES('#{pl.fetch :id}','#{v.fetch :id}', 'youtube', '#{v.fetch :title}', '#{v.fetch :description}');"
        end
      end

    end
  end

end