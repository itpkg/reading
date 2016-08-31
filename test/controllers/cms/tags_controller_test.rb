require 'test_helper'

class Cms::TagsControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get cms_tags_index_url
    assert_response :success
  end

end
