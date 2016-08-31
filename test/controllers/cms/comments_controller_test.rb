require 'test_helper'

class Cms::CommentsControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get cms_comments_index_url
    assert_response :success
  end

end
