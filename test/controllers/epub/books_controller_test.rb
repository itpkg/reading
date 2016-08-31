require 'test_helper'

class Epub::BooksControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get epub_books_index_url
    assert_response :success
  end

end
