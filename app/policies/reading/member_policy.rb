module Reading
  class MemberPolicy < ApplicationPolicy
    def index?
      !user.nil? && user.is_admin?
    end

    def destroy?
      !user.nil? && user.is_admin?
    end
  end
end
