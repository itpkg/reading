module Reading
  class BookPolicy < ApplicationPolicy

    def destroy?
      !user.nil? &&  user.is_admin?
    end
  end
end
