require 'test_helper'

class Epub::PagesControllerTest < ActionDispatch::IntegrationTest
  test "should get show" do
    get epub_pages_show_url
    assert_response :success
  end

end
