module Reading
  class NotePolicy < ApplicationPolicy
    def create?
      !user.nil?
    end

    def update?
      !user.nil? && (record.user_id == user.id || user.is_admin?)
    end

    def destroy?
      !user.nil? && (record.user_id == user.id || user.is_admin?)
    end
  end
end
