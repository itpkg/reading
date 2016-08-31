require 'test_helper'

class Cms::ArticlesControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get cms_articles_index_url
    assert_response :success
  end

end
