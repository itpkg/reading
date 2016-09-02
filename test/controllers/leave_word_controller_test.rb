require 'test_helper'

class LeaveWordControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get leave_word_index_url
    assert_response :success
  end

end
